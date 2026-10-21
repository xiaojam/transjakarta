package main

import (
	"encoding/json"
	"fmt"
	"log"

	"time"

	"transjakarta/config"
	"transjakarta/internal/dto"
	pkgMqtt "transjakarta/pkg/mqtt"

	"github.com/shopspring/decimal"
)

func main() {
	cfg := config.LoadConfig()
	client := pkgMqtt.NewMQTTClient(cfg.MQTTBroker, "mock_publisher_script")

	vehicleID := "B1234XYZ"
	topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)

	coords := []struct {
		lat string
		lng string
	}{
		{"-6.22000000", "106.84000000"}, 
		{"-6.21000000", "106.84500000"}, 
		{"-6.20880000", "106.84560000"}, 
		{"-6.20879999", "106.84559999"}, 
		{"-6.20000000", "106.85000000"}, 
	}

	idx := 0
	log.Println("[Mock Publisher] Starting to publish mock location data")

	for {
		c := coords[idx%len(coords)]
		latDec, _ := decimal.NewFromString(c.lat)
		lngDec, _ := decimal.NewFromString(c.lng)

		payload := dto.LocationPayload{
			VehicleID: vehicleID,
			Latitude:  latDec,
			Longitude: lngDec,
			Timestamp: time.Now().Unix(),
		}

		body, _ := json.Marshal(payload)
		token := client.Publish(topic, 1, false, body)
		token.Wait()

		log.Printf("[Mock Publisher] Sent -> %s", string(body))

		idx++
		time.Sleep(2 * time.Second)
	}
}