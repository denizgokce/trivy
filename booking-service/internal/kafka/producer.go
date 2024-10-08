package kafka

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer

func InitProducer() error {
	brokers := []string{os.Getenv("KAFKA_BROKER")}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	var err error
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}

	return nil
}

func ProduceMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to produce message: %v", err)
		return err
	}

	return nil
}
