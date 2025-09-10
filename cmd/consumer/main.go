package main

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type MyMessage struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9095"}, nil)
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("topic-1", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("failed to consume partition: %v", err)
	}
	defer partConsumer.Close()

	for {
		select {
		case msg, ok := <-partConsumer.Messages():
			if !ok {
				log.Println("channel closed, exiting")
				return
			}

			var receivedMessage MyMessage
			if err := json.Unmarshal(msg.Value, &receivedMessage); err != nil {
				log.Printf("error unmarshaling JSON: %v\n", err)
				continue
			}

			log.Printf("received message: %+v\n", receivedMessage)
		}
	}
}
