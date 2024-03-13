package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

type MyMessage struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9095"}, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer func(consumer sarama.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(consumer)

	partConsumer, err := consumer.ConsumePartition("topic-1", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
	}
	defer func(partConsumer sarama.PartitionConsumer) {
		err := partConsumer.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(partConsumer)

	for {
		select {
		case msg, ok := <-partConsumer.Messages():
			if !ok {
				log.Println("Channel closed, exiting")
				return
			}

			var receivedMessage MyMessage
			err := json.Unmarshal(msg.Value, &receivedMessage)

			if err != nil {
				log.Printf("Error unmarshaling JSON: %v\n", err)
				continue
			}

			log.Printf("Received message: %+v\n", receivedMessage)
		}
	}
}
