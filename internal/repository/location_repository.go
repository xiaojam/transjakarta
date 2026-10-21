package repository

import (
	"transjakarta/internal/model"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Save(location *model.VehicleLocation) error
	GetLatestLocation(vehicleID string) (*model.VehicleLocation, error)
	GetHistory(vehicleID string, start, end int64) ([]model.VehicleLocation, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) Save(location *model.VehicleLocation) error {
	return r.db.Create(location).Error
}

func (r *locationRepository) GetLatestLocation(vehicleID string) (*model.VehicleLocation, error) {
	var loc model.VehicleLocation
	err := r.db.Where("vehicle_id = ?", vehicleID).Order("timestamp DESC").First(&loc).Error
	if err != nil {
		return nil, err
	}
	return &loc, nil
}

func (r *locationRepository) GetHistory(vehicleID string, start, end int64) ([]model.VehicleLocation, error) {
	var history []model.VehicleLocation
	err := r.db.Where("vehicle_id = ? AND timestamp BETWEEN ? AND ?", vehicleID, start, end).
		Order("timestamp ASC").Find(&history).Error
	return history, err
}