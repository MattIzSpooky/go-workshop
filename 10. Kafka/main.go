package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	topic := "my-topic"
	brokerAddress := "localhost:9092"
	partition := 0

	// Setting up a kafka connection
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	defer conn.Close()

	writeMessages(conn, brokerAddress, topic)
	readMessages(brokerAddress, topic)
}

func writeMessages(conn *kafka.Conn, url string, topic string) {
	// Writing messages
	_, err := conn.WriteMessages(
		kafka.Message{Topic: topic, Value: []byte("one!")},
		kafka.Message{Topic: topic, Value: []byte("two!")},
		kafka.Message{Topic: topic, Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}

func readMessages(url string, topic string) {
	// Reading messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{url}, // can add more
		Topic:     topic,
		GroupID:   "my-group-id", // consumer groups
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
