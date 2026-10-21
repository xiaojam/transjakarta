package service

import (
	"encoding/json"
	"math"

	"transjakarta/internal/dto"
	"transjakarta/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shopspring/decimal"
)

type GeofenceService interface {
	CheckAndPublishGeofence(payload dto.LocationPayload) error
}

type geofenceService struct {
	rmq          *rabbitmq.RabbitMQ
	targetLat    decimal.Decimal
	targetLng    decimal.Decimal
	radiusMeters float64
}

func NewGeofenceService(rmq *rabbitmq.RabbitMQ) GeofenceService {
	return &geofenceService{
		rmq:          rmq,
		targetLat:    decimal.NewFromFloat(-6.20880000),
		targetLng:    decimal.NewFromFloat(106.84560000),
		radiusMeters: 50.0,
	}
}

func (g *geofenceService) CheckAndPublishGeofence(payload dto.LocationPayload) error {
	lat1, _ := payload.Latitude.Float64()
	lon1, _ := payload.Longitude.Float64()
	lat2, _ := g.targetLat.Float64()
	lon2, _ := g.targetLng.Float64()

	dist := calculateHaversine(lat1, lon1, lat2, lon2)

	if dist <= g.radiusMeters {
		event := dto.GeofenceEventPayload{
			VehicleID: payload.VehicleID,
			Event:     "geofence_entry",
			Location: dto.GeofenceLocationSpec{
				Latitude:  payload.Latitude,
				Longitude: payload.Longitude,
			},
			Timestamp: payload.Timestamp,
		}

		body, err := json.Marshal(event)
		if err != nil {
			return err
		}

		return g.rmq.Channel.Publish(
			"fleet.events",
			"geofence.entry",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
	}
	return nil
}

func calculateHaversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}