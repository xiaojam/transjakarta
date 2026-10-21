package main

import (
	"log"

	"transjakarta/config"

	httpDelivery "transjakarta/internal/delivery/http"
	httpHandler "transjakarta/internal/delivery/http/handler"
	mqttDelivery "transjakarta/internal/delivery/mqtt"
	workerDelivery "transjakarta/internal/delivery/worker"

	"transjakarta/internal/repository"
	"transjakarta/internal/service"

	"transjakarta/pkg/database"
	pkgMqtt "transjakarta/pkg/mqtt"
	pkgRmq "transjakarta/pkg/rabbitmq"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitDB(cfg)
	mqttClient := pkgMqtt.NewMQTTClient(cfg.MQTTBroker, "transjakarta_backend_core")
	rmq := pkgRmq.NewRabbitMQ(cfg.RabbitMQURL)

	defer rmq.Conn.Close()
	defer rmq.Channel.Close()

	locRepo := repository.NewLocationRepository(db)
	geofenceService := service.NewGeofenceService(rmq)
	locService := service.NewLocationService(locRepo, geofenceService)

	mqttSub := mqttDelivery.NewLocationSubscriber(mqttClient, locService)
	mqttSub.Start()

	geofenceWorker := workerDelivery.NewGeofenceWorker(rmq)
	geofenceWorker.Start()

	locHandler := httpHandler.NewLocationHandler(locService)
	router := httpDelivery.SetupRouter(locHandler)

	log.Printf("[Server] Starting HTTP server on port %s...", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("[Server] Failed to start server: %v", err)
	}
}