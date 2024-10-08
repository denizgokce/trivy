package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"class-service/internal/db"
	"class-service/internal/models"

	"github.com/Shopify/sarama"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var consumer sarama.Consumer

func InitConsumer() error {
	brokers := []string{os.Getenv("KAFKA_BROKER")}
	var err error
	consumer, err = sarama.NewConsumer(brokers, nil)
	if err != nil {
		return err
	}

	return nil
}

func ConsumeMessages(topic string) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consumer for partition: %v", err)
	}

	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		var booking models.Booking
		err := json.Unmarshal(message.Value, &booking)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		updateClassAvailability(booking.ClassID)
	}
}

func updateClassAvailability(classID primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	classCollection := db.GetCollection(db.Client, "classes")

	var class models.Class
	err := classCollection.FindOne(ctx, bson.M{"_id": classID}).Decode(&class)
	if err != nil {
		log.Printf("Failed to find class: %v", err)
		return
	}

	// Decrement available slots by 1
	availableSlots := class.AvailableSlots - 1

	// Update available slots
	_, err = classCollection.UpdateOne(ctx, bson.M{"_id": class.ID}, bson.M{"$set": bson.M{"availableSlots": availableSlots}})
	if err != nil {
		log.Printf("Failed to update class availability: %v", err)
	}
}
