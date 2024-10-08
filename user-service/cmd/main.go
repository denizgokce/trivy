package main

import (
	"log"
	"net/http"

	"user-service/internal/db"
	"user-service/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	db.SeedUsers(database)

	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/validate", handlers.ValidateUserHandler).Methods("POST")

	log.Println("Starting user service on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
