package dto

import "github.com/shopspring/decimal"

type LocationPayload struct {
	VehicleID string          `json:"vehicle_id" binding:"required"`
	Latitude  decimal.Decimal `json:"latitude" binding:"required"`
	Longitude decimal.Decimal `json:"longitude" binding:"required"`
	Timestamp int64           `json:"timestamp" binding:"required"`
}

type LocationResponse struct {
	ID        string          `json:"id"`
	VehicleID string          `json:"vehicle_id"`
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
	Timestamp int64           `json:"timestamp"`
}

type GeofenceEventPayload struct {
	VehicleID string               `json:"vehicle_id"`
	Event     string               `json:"event"`
	Location  GeofenceLocationSpec `json:"location"`
	Timestamp int64                `json:"timestamp"`
}

type GeofenceLocationSpec struct {
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
}