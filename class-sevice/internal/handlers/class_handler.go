package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"class-service/internal/db"
	"class-service/internal/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var classCollection = db.GetCollection(db.Client, "classes")

func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	class.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := classCollection.InsertOne(ctx, class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func GetClassHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var class models.Class
	err = classCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&class)
	if err != nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	// Use the AvailableSlots field directly
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

func UpdateClassAvailabilityHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var class models.Class
	err = classCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&class)
	if err != nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}

	// Check availability by querying the booking-service
	availableSlots, err := getAvailableSlots(class.ID)
	if err != nil {
		http.Error(w, "Failed to check availability", http.StatusInternalServerError)
		return
	}

	// Update available slots
	class.AvailableSlots = availableSlots
	_, err = classCollection.UpdateOne(ctx, bson.M{"_id": class.ID}, bson.M{"$set": bson.M{"availableSlots": availableSlots}})
	if err != nil {
		http.Error(w, "Failed to update class availability", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
