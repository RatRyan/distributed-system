package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"kafka:9092"}
	topics := []string{"user", "offer"}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Panicf("")
	}
	defer consumer.Close()

	messages := make(chan *sarama.ConsumerMessage)
	for _, topic := range topics {
		partitions, err := consumer.Partitions(topic)
		if err != nil {
			log.Panicf("Error getting partitions for topic %s: %v", topic, err)
		}

		for _, partition := range partitions {
			pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				log.Panicf("Error consuming partition %d of topic %s: %v", partition, topic, err)
			}

			go func() {
				for msg := range pc.Messages() {
					messages <- msg
				}
			}()
		}
	}

	// Listen on messages channel
	go func() {
		for msg := range messages {
			switch msg.Topic {
			case "user":
				var user userPayload
				json.Unmarshal(msg.Value, &user)
				handleUser(string(msg.Key), user)
			case "offer":
				var offer offerPaylood
				json.Unmarshal(msg.Value, &offer)
				handleOffer(string(msg.Key), offer)
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
}
