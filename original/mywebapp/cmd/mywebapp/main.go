package main

import (
	"log"
	"mywebapp/internal/mywebapp"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", mywebapp.NewController()))
}
