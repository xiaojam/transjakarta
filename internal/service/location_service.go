package service

import (
	"transjakarta/internal/dto"
	"transjakarta/internal/model"
	"transjakarta/internal/repository"
)

type LocationService interface {
	ProcessLocation(payload dto.LocationPayload) error
	GetLatestLocation(vehicleID string) (*dto.LocationResponse, error)
	GetLocationHistory(vehicleID string, start, end int64) ([]dto.LocationResponse, error)
}

type locationService struct {
	repo            repository.LocationRepository
	geofenceService GeofenceService
}

func NewLocationService(repo repository.LocationRepository, geofenceService GeofenceService) LocationService {
	return &locationService{
		repo:            repo,
		geofenceService: geofenceService,
	}
}

func (s *locationService) ProcessLocation(payload dto.LocationPayload) error {
	loc := &model.VehicleLocation{
		VehicleID: payload.VehicleID,
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		Timestamp: payload.Timestamp,
	}

	if err := s.repo.Save(loc); err != nil {
		return err
	}

	_ = s.geofenceService.CheckAndPublishGeofence(payload)
	return nil
}

func (s *locationService) GetLatestLocation(vehicleID string) (*dto.LocationResponse, error) {
	loc, err := s.repo.GetLatestLocation(vehicleID)
	if err != nil {
		return nil, err
	}

	return &dto.LocationResponse{
		ID:        loc.ID.String(),
		VehicleID: loc.VehicleID,
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Timestamp: loc.Timestamp,
	}, nil
}

func (s *locationService) GetLocationHistory(vehicleID string, start, end int64) ([]dto.LocationResponse, error) {
	history, err := s.repo.GetHistory(vehicleID, start, end)
	if err != nil {
		return nil, err
	}

	var response []dto.LocationResponse
	for _, h := range history {
		response = append(response, dto.LocationResponse{
			ID:        h.ID.String(),
			VehicleID: h.VehicleID,
			Latitude:  h.Latitude,
			Longitude: h.Longitude,
			Timestamp: h.Timestamp,
		})
	}
	return response, nil
}