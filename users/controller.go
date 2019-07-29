package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var iUsersRepo IUserRepository

func init() {
	ctx := context.Background()
	usersRepo, err := NewUsersRepository(ctx)

	if err != nil {
		log.Fatalf("failed to initialize user repository: %v", err)
	}

	iUsersRepo = usersRepo
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Bad request : %v", err), http.StatusBadRequest)
		return
	}
	if err := iUsersRepo.CreateUser(r.Context(), user); err != nil {
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	user, err := iUsersRepo.GetUser(r.Context(), userId)

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	err := iUsersRepo.DeleteUser(r.Context(), userId)

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
