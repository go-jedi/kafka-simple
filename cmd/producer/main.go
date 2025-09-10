package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type MyMessage struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9095"}, nil)
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
		return
	}
	defer producer.Close()

	message := MyMessage{
		ID:    1,
		Name:  "Go-jedi",
		Value: "Hello man",
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: "topic-1",
		Value: sarama.ByteEncoder(bytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("message sent to partition %d at offset %d\n", partition, offset)
}
