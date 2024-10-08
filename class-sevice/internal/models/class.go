package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Class struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string             `bson:"name" json:"name"`
	Description    string             `bson:"description" json:"description"`
	VenueID        primitive.ObjectID `bson:"venueId" json:"venueId"`
	StartTimestamp time.Time          `bson:"startTimestamp" json:"startTimestamp"`
	EndTimestamp   time.Time          `bson:"endTimestamp" json:"endTimestamp"`
	NumberOfSlots  int                `bson:"numberOfSlots" json:"numberOfSlots"`
	Status         string             `bson:"status" json:"status"`
	AvailableSlots int                `bson:"availableSlots,omitempty" json:"availableSlots,omitempty"`
}
