package http

import (
	"transjakarta/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(locationHandler *handler.LocationHandler) *gin.Engine {
	r := gin.Default()

	// Route Grouping langsung berawalan /vehicles
	vehicles := r.Group("/vehicles")
	{
		vehicles.GET("/:vehicle_id/location", locationHandler.GetLatestLocation)
		vehicles.GET("/:vehicle_id/history", locationHandler.GetLocationHistory)
	}

	return r
}