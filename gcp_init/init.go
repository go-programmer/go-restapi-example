package mywebapp

import (
	"fmt"
	users "go-practice/simprints/reengineered/users"
	utilities "go-practice/simprints/reengineered/utilities"
	"log"
	"net/http"
)

func init() {

	fmt.Println(" mywebapp init called ")

	var err error

	router := utilities.NewRouter(users.UsersRoutes)

	if err != nil {
		log.Fatalf("failed to initialize user repository: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))

}

// the idea was to define a global context and
//  access it via GetContext in other packages
/// as needed.
// var ctx context.Context
// func GetContext() context.Context {
// 	return ctx
// }
