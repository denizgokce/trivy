package main

import (
	"log"
	"net/http"

	"auth-service/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/validate", handlers.ValidateTokenHandler).Methods("POST")

	log.Println("Starting auth service on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
