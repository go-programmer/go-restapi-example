package gcpapp

import (
	"context"
	"log"
	Todo "mywebapp/v9/todo"
	Util "mywebapp/v9/utilities"
	"net/http"
)

func init() {
	var err error
	ctx := context.Background()

	router := Util.NewRouter(ctx, Todo.TodoRoutes)

	if err != nil {
		log.Fatalf("failed to initialize user repository: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))

}

// func main() {}
