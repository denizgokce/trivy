package db

import (
	"context"
	"log"
	"time"

	"class-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedClasses(client *mongo.Client) {
	collection := GetCollection(client, "classes")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatalf("Failed to count classes: %v", err)
	}

	if count == 0 {
		log.Println("Seeding classes...")

		// Manually create sample classes
		classes := []models.Class{
			{
				ID:             primitive.NewObjectID(),
				Name:           "Yoga Class",
				Description:    "A relaxing yoga class.",
				VenueID:        primitive.NewObjectID(), // Replace with actual VenueID from your venue collection
				StartTimestamp: time.Now().Add(24 * time.Hour),
				EndTimestamp:   time.Now().Add(26 * time.Hour),
				NumberOfSlots:  10,
				Status:         "scheduled",
			},
			{
				ID:             primitive.NewObjectID(),
				Name:           "Spinning Class",
				Description:    "A high-intensity spinning class.",
				VenueID:        primitive.NewObjectID(), // Replace with actual VenueID from your venue collection
				StartTimestamp: time.Now().Add(48 * time.Hour),
				EndTimestamp:   time.Now().Add(50 * time.Hour),
				NumberOfSlots:  15,
				Status:         "scheduled",
			},
			{
				ID:             primitive.NewObjectID(),
				Name:           "Pilates Class",
				Description:    "A core-strengthening pilates class.",
				VenueID:        primitive.NewObjectID(), // Replace with actual VenueID from your venue collection
				StartTimestamp: time.Now().Add(72 * time.Hour),
				EndTimestamp:   time.Now().Add(74 * time.Hour),
				NumberOfSlots:  12,
				Status:         "scheduled",
			},
		}

		for _, class := range classes {
			_, err := collection.InsertOne(ctx, class)
			if err != nil {
				log.Fatalf("Failed to seed class: %v", err)
			}
		}

		log.Println("Classes seeded successfully.")
	} else {
		log.Println("Classes already exist, skipping seeding.")
	}
}
