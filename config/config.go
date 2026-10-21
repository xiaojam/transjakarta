package config

import "os"

type Config struct {
	AppPort     string
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	DBSSLMode	string
	DBTimezone	string
	MQTTBroker  string
	RabbitMQURL string
}

func LoadConfig() *Config {
	return &Config{
		AppPort:     getEnv("APP_PORT", "8080"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "postgres"),
		DBName:      getEnv("DB_NAME", "transjakarta"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBSSLMode:   getEnv("DB_SSL", "disable"),
		DBTimezone:  getEnv("DB_TZ", "Asia/Jakarta"),
		MQTTBroker:  getEnv("MQTT_BROKER", "tcp://localhost:1883"),
		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}