package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID  primitive.ObjectID `bson:"userId" json:"userId"`
	VenueID primitive.ObjectID `bson:"venueId" json:"venueId"`
	ClassID primitive.ObjectID `bson:"classId" json:"classId"`
	Status  string             `bson:"status" json:"status"`
}
