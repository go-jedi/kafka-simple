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
	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9095"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
		return
	}
	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(producer)

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

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Println(err)
		return
	}
}
