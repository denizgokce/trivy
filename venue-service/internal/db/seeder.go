package db

import (
	"database/sql"
	"log"
	"venue-service/internal/models"
)

func SeedVenues(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM venues").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to count venues: %v", err)
	}

	if count == 0 {
		log.Println("Seeding venues...")

		venues := []models.Venue{
			{Name: "Sports Arena", Location: "Downtown", Description: "A large sports arena."},
			{Name: "Community Gym", Location: "Suburbs", Description: "A community gym with various facilities."},
		}

		for _, venue := range venues {
			_, err := db.Exec("INSERT INTO venues (name, location, description) VALUES ($1, $2, $3)", venue.Name, venue.Location, venue.Description)
			if err != nil {
				log.Fatalf("Failed to seed venue: %v", err)
			}
		}

		log.Println("Venues seeded successfully.")
	} else {
		log.Println("Venues already exist, skipping seeding.")
	}
}
