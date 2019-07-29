package mywebapp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mywebapp/internal/mywebapp/users"
	"net/http"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var userRepo users.Repository

func init() {
	var err error
	ctx := context.Background()
	userRepo, err = users.NewRepository(ctx)
	if err != nil {
		log.Fatalf("failed to initialize user repository: %v", err)
	}
}

type Controller struct {
	*http.ServeMux
}

func NewController() Controller {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", UsersHandler)
	mux.HandleFunc("/users/", UserHandler)
	return Controller{
		ServeMux: mux,
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		PostUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

var UserIdPathParamRe = regexp.MustCompile("^/users/([^/]*)$")

func UserHandler(w http.ResponseWriter, r *http.Request) {
	matches := UserIdPathParamRe.FindStringSubmatch(r.URL.Path)
	if matches != nil {
		userId := matches[1]
		switch r.Method {
		case http.MethodGet:
			GetUser(userId)(w, r)
		case http.MethodDelete:
			DeleteUser(userId)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Bad request : %v", err), http.StatusBadRequest)
		return
	}
	if err := userRepo.CreateUser(r.Context(), user); err != nil {
		code := status.Code(err)
		if code == codes.AlreadyExists {
			http.Error(w, fmt.Sprintf("User with ID %s already exists", user.Id), http.StatusConflict)
		} else {
			log.Printf("failed to created user %v: %v", user, err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUser(userId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := userRepo.GetUser(r.Context(), userId)
		if err != nil {
			code := status.Code(err)
			if code == codes.NotFound {
				http.Error(w, fmt.Sprintf("No user with ID %s", userId), http.StatusNotFound)
			} else {
				log.Printf("failed to get user with ID %s: %v", userId, err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Printf("failed to get user with ID %s: %v", userId, err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func DeleteUser(userId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := userRepo.DeleteUser(r.Context(), userId)
		if err != nil {
			code := status.Code(err)
			if code == codes.NotFound {
				http.Error(w, fmt.Sprintf("No user with ID %s", userId), http.StatusNotFound)
			} else {
				log.Printf("failed to get user with ID %s: %v", userId, err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
