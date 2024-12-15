package main

import (
	"fmt"
	"log"
	"product-management/queue"

	"github.com/streadway/amqp"
)

func publishMessage(message string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"image_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
		return err
	}

	fmt.Println("Published message to queue:", message)
	return nil
}

func main() {
	message := "https://drive.google.com/uc?export=download&id=11uSge45PgB0hbu2IrWZClQgWKp-Svkba"
	err := publishMessage(message)
	if err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	log.Println("Starting the consumer...")
	queue.StartConsumer()
}
