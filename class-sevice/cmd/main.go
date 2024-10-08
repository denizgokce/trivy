package main

import (
	"context"
	"log"
	"net/http"

	"class-service/internal/db"
	"class-service/internal/handlers"
	"class-service/internal/kafka"

	"github.com/gorilla/mux"
)

func main() {
	client, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	err = kafka.InitConsumer()
	if err != nil {
		log.Fatal("Failed to initialize Kafka consumer:", err)
	}

	go kafka.ConsumeMessages("booking-events")

	r := mux.NewRouter()
	r.HandleFunc("/classes", handlers.CreateClassHandler).Methods("POST")
	r.HandleFunc("/classes/{id}", handlers.GetClassHandler).Methods("GET")
	r.HandleFunc("/classes/{id}/update-availability", handlers.UpdateClassAvailabilityHandler).Methods("POST")

	log.Println("Starting class service on :8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
