package handler

import (
	"net/http"
	"strconv"

	"transjakarta/internal/service"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service service.LocationService
}

func NewLocationHandler(service service.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

func (h *LocationHandler) GetLatestLocation(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")
	res, err := h.service.GetLatestLocation(vehicleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location data not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *LocationHandler) GetLocationHistory(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")
	startStr := c.Query("start")
	endStr := c.Query("end")

	start, err1 := strconv.ParseInt(startStr, 10, 64)
	end, err2 := strconv.ParseInt(endStr, 10, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameters 'start' and 'end' must be integer UNIX timestamps"})
		return
	}

	res, err := h.service.GetLocationHistory(vehicleID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}