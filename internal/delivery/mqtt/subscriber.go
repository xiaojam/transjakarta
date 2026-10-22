package mqtt

import (
	"encoding/json"
	"log"

	"transjakarta/internal/dto"
	"transjakarta/internal/service"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

type LocationSubscriber struct {
	client  pahoMqtt.Client
	service service.LocationService
}

func NewLocationSubscriber(client pahoMqtt.Client, service service.LocationService) *LocationSubscriber {
	return &LocationSubscriber{
		client:  client,
		service: service,
	}
}

func (s *LocationSubscriber) Start() {
	topic := "/fleet/vehicle/+/location"
	token := s.client.Subscribe(topic, 1, s.handleMessage)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("[MQTT Subscriber] Failed to subscribe to topic '%s': %v", topic, token.Error())
	}
	log.Printf("[MQTT Subscriber] Listening on topic: %s", topic)
}

func (s *LocationSubscriber) handleMessage(client pahoMqtt.Client, msg pahoMqtt.Message) {
	var payload dto.LocationPayload
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Printf("[MQTT Error] Invalid JSON payload format: %v", err)
		return
	}

	if payload.VehicleID == "" || payload.Latitude.IsZero() || payload.Longitude.IsZero() || payload.Timestamp == 0 {
		log.Println("[MQTT Error] Missing or invalid payload fields")
		return
	}

	if err := s.service.ProcessLocation(payload); err != nil {
		log.Printf("[MQTT Error] Failed to process vehicle location: %v", err)
	}
}