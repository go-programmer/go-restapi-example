package main

import (
	"context"
	"log"
	Todo "mywebapp/v9/todo"
	"mywebapp/v9/utilities"
	Util "mywebapp/v9/utilities"
	"net/http"
)

func main() {
	var err error
	ctx := context.Background()

	router := Util.NewRouter(Todo.TodoRoutes)
	
	firestore.Client client, error err := utilities.NewFireStoreClient(ctx)
	
	// send this client to userRepo as the data source
	
	if err != nil {
		log.Fatalf("failed to initialize user repository: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
