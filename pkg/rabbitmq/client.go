package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(url string) *RabbitMQ {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("[RabbitMQ] Failed to connect: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[RabbitMQ] Failed to open channel: %v", err)
	}

	err = ch.ExchangeDeclare(
		"fleet.events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("[RabbitMQ] Failed to declare exchange fleet.events: %v", err)
	}

	q, err := ch.QueueDeclare(
		"geofence_alerts",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("[RabbitMQ] Failed to declare queue geofence_alerts: %v", err)
	}

	err = ch.QueueBind(
		q.Name,
		"geofence.#",
		"fleet.events",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("[RabbitMQ] Failed to bind queue to exchange: %v", err)
	}

	return &RabbitMQ{Conn: conn, Channel: ch}
}