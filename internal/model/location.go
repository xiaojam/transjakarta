package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type VehicleLocation struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	VehicleID string          `gorm:"type:varchar(50);not null;index:idx_vehicle_time" json:"vehicle_id"`
	Latitude  decimal.Decimal `gorm:"type:decimal(10,8);not null" json:"latitude"`
	Longitude decimal.Decimal `gorm:"type:decimal(11,8);not null" json:"longitude"`
	Timestamp int64           `gorm:"not null;index:idx_vehicle_time" json:"timestamp"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
}

func (VehicleLocation) TableName() string {
	return "vehicle_locations"
}