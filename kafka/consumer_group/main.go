package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/neo-classic/golang-cqrs-mvp/kafka/consumer_group/handlers"
)

var (
	kafkaBrokers    = []string{"localhost:29092"}
	kafkaTopics     = []string{"event.new"}
	consumerGroupID = "consumer1"
)

func main() {
	// Init config, specify appropriate version
	config := sarama.NewConfig()

	//sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
	//config.Version = sarama.V2_1_0_0

	// Start with a client
	client, err := sarama.NewClient(kafkaBrokers, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(consumerGroupID, client)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = group.Close() }()
	log.Println("Consumer up and running")

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		handler := handlers.ConsumerGroupHandler{}

		err := group.Consume(ctx, kafkaTopics, handler)
		if err != nil {
			panic(err)
		}
	}
}
