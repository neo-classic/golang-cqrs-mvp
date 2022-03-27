package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/neo-classic/golang-cqrs-mvp/pkg/adapters/kafka"
)

const (
	kafkaConn = "localhost:29092"
	topic     = "event.new"
)

func main() {
	producer, err := kafka.NewProducer([]string{kafkaConn})
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		publish(msg, producer)
	}
}

func publish(message string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
