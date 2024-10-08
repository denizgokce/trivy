package main

import (
	"log"
	"net/http"

	"venue-service/internal/db"
	"venue-service/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	db.SeedVenues(database)

	r := mux.NewRouter()
	r.HandleFunc("/venues", handlers.CreateVenueHandler).Methods("POST")
	r.HandleFunc("/venues/{id}", handlers.GetVenueHandler).Methods("GET")
	r.HandleFunc("/venues/{id}", handlers.UpdateVenueHandler).Methods("PUT")
	r.HandleFunc("/venues/{id}", handlers.DeleteVenueHandler).Methods("DELETE")

	log.Println("Starting venue service on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
