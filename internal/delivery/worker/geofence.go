package worker

import (
	"log"

	"transjakarta/pkg/rabbitmq"
)

type GeofenceWorker struct {
	rmq *rabbitmq.RabbitMQ
}

func NewGeofenceWorker(rmq *rabbitmq.RabbitMQ) *GeofenceWorker {
	return &GeofenceWorker{rmq: rmq}
}

func (w *GeofenceWorker) Start() {
	msgs, err := w.rmq.Channel.Consume(
		"geofence_alerts",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("[RabbitMQ Worker] Failed to consume queue 'geofence_alerts': %v", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("[RabbitMQ Worker] Geofence Event Alert Received: %s", d.Body)
		}
	}()

	log.Println("[RabbitMQ Worker] Ready to consume messages from 'geofence_alerts' queue...")
}