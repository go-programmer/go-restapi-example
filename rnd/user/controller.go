package user

import (
	"net/http"
)

type userController struct {
	
}

var err error

func PostUser(w http.ResponseWriter, r *http.Request) {
	// user := user.User{}
	// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	// 	http.Error(w, fmt.Sprintf("Bad request : %v", err), http.StatusBadRequest)
	// 	return
	// }
	// if err := userRepo.CreateUser(r.Context(), user); err != nil {
	// 	code := status.Code(err)
	// 	if code == codes.AlreadyExists {
	// 		http.Error(w, fmt.Sprintf("User with ID %s already exists", user.Id), http.StatusConflict)
	// 	} else {
	// 		log.Printf("failed to created user %v: %v", user, err)
	// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	}
	// 	return
	// }
	// w.WriteHeader(http.StatusCreated)
}

// func GetUser(userId string) http.HandlerFunc {
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	var userId string

// 	if todoId, err = strconv.Atoi(vars["userId"]); err != nil {
// 		panic(err)
// 	}

// 	user, err := userRepo.GetUser(r.Context(), userId)
// 	if err != nil {
// 		code := status.Code(err)
// 		if code == codes.NotFound {
// 			http.Error(w, fmt.Sprintf("No user with ID %s", userId), http.StatusNotFound)
// 		} else {
// 			log.Printf("failed to get user with ID %s: %v", userId, err)
// 			http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	if err := json.NewEncoder(w).Encode(user); err != nil {
// 		log.Printf("failed to get user with ID %s: %v", userId, err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}

// }

// func DeleteUser(userId string) http.HandlerFunc {
// func DeleteUser(w http.ResponseWriter, r *http.Request) {

// 	err := userRepo.DeleteUser(r.Context(), userId)

// 	if err != nil {
// 		code := status.Code(err)
// 		if code == codes.NotFound {
// 			http.Error(w, fmt.Sprintf("No user with ID %s", userId), http.StatusNotFound)
// 		} else {
// 			log.Printf("failed to get user with ID %s: %v", userId, err)
// 			http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
