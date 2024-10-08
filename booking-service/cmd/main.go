package main

import (
	"context"
	"log"
	"net/http"

	"booking-service/internal/db"
	"booking-service/internal/handlers"
	"booking-service/internal/kafka"

	"github.com/gorilla/mux"
)

func main() {
	client, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	err = kafka.InitProducer()
	if err != nil {
		log.Fatal("Failed to initialize Kafka producer:", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/bookings", handlers.CreateBookingHandler).Methods("POST")
	r.HandleFunc("/bookings/{id}", handlers.GetBookingHandler).Methods("GET")
	r.HandleFunc("/bookings/{id}", handlers.UpdateBookingHandler).Methods("PUT")
	r.HandleFunc("/bookings/{id}", handlers.DeleteBookingHandler).Methods("DELETE")
	r.HandleFunc("/bookings/count", handlers.GetBookingCountHandler).Methods("GET")

	log.Println("Starting booking service on :8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
