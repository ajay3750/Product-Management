package queue

import (
	"log"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channel *amqp.Channel

func InitRabbitMQ() {
	var err error
	connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ: ", err)
	}

	channel, err = connection.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel: ", err)
	}

	_, err = channel.QueueDeclare(
		"image_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue: ", err)
	}
}

func PublishMessage(queueName, message string) error {
	return channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}
