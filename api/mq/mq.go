package mq

import (
	"api/model"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

func ProducePasswordChange(user *model.User) {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}

	userJson, _ := json.Marshal(user)

	msg := &sarama.ProducerMessage{
		Topic: "user",
		Key:   sarama.StringEncoder("update_password"),
		Value: sarama.StringEncoder(userJson),
	}

	producer.SendMessage(msg)
}

func ProduceOfferCreated(offer *model.Offer) {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}

	offerJson, _ := json.Marshal(offer)

	msg := &sarama.ProducerMessage{
		Topic: "offer",
		Key:   sarama.StringEncoder("offer_created"),
		Value: sarama.StringEncoder(offerJson),
	}

	producer.SendMessage(msg)
}

func ProduceOfferAccepted(offer *model.Offer) {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}

	offerJson, _ := json.Marshal(offer)

	msg := &sarama.ProducerMessage{
		Topic: "offer",
		Key:   sarama.StringEncoder("offer_accepted"),
		Value: sarama.StringEncoder(offerJson),
	}

	producer.SendMessage(msg)
}

func ProduceOfferRejected(offer *model.Offer) {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}

	offerJson, _ := json.Marshal(offer)

	msg := &sarama.ProducerMessage{
		Topic: "offer",
		Key:   sarama.StringEncoder("offer_rejected"),
		Value: sarama.StringEncoder(offerJson),
	}

	producer.SendMessage(msg)
}
