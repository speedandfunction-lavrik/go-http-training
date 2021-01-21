package main

import (
	"fmt"
	"http-go/handlers"
	"http-go/storage"
	"log"
	"net/http"
)

const port = "4040"

func main() {
	storage := storage.Users{}
	srv := http.NewServeMux()

	srv.HandleFunc("/users/create", handlers.CreateUser(storage))
	srv.HandleFunc("/users/list", handlers.GetUsers(storage))
	srv.HandleFunc("/users/get", handlers.GetUser(storage))
	srv.HandleFunc("/users/update", handlers.UpdateUser(storage))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), srv)

	if err != nil {
		log.Panic(err)
	}
}
