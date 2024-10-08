package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"venue-service/internal/db"
	"venue-service/internal/models"

	"github.com/gorilla/mux"
)

// CreateVenueHandler handles the creation of a new venue
func CreateVenueHandler(w http.ResponseWriter, r *http.Request) {
	var venue models.Venue
	if err := json.NewDecoder(r.Body).Decode(&venue); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	_, err = database.Exec("INSERT INTO venues (name, location, description) VALUES ($1, $2, $3)", venue.Name, venue.Location, venue.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(venue)
}

// GetVenueHandler handles fetching a venue by ID
func GetVenueHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid venue ID", http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	var venue models.Venue
	err = database.QueryRow("SELECT id, name, location, description FROM venues WHERE id = $1", id).Scan(&venue.ID, &venue.Name, &venue.Location, &venue.Description)
	if err != nil {
		http.Error(w, "Venue not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venue)
}

// UpdateVenueHandler handles updating a venue by ID
func UpdateVenueHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid venue ID", http.StatusBadRequest)
		return
	}

	var venue models.Venue
	if err := json.NewDecoder(r.Body).Decode(&venue); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	_, err = database.Exec("UPDATE venues SET name = $1, location = $2, description = $3 WHERE id = $4", venue.Name, venue.Location, venue.Description, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venue)
}

// DeleteVenueHandler handles deleting a venue by ID
func DeleteVenueHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid venue ID", http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	_, err = database.Exec("DELETE FROM venues WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
