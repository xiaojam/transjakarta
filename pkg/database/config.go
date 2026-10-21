package database

import (
	"fmt"
	"log"
	"time"

	"transjakarta/config"
	"transjakarta/internal/model"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.DBTimezone)

	var db *gorm.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("[Database] Failed to connect to PostgreSQL (attempt %d/10): %v", i, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("[Database] Could not connect to PostgreSQL after 10 attempts: %v", err)
	}

	if err := db.AutoMigrate(&model.VehicleLocation{}); err != nil {
		log.Fatalf("[Database] Failed to perform AutoMigrate: %v", err)
	}

	runGolangMigrate(cfg)

	log.Println("[Database] Successfully connected to PostgreSQL, performed AutoMigrate, and ran SQL migrations.")
	return db
}

func runGolangMigrate(cfg *config.Config) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Printf("[Migration] Failed to initialize migrate: %v", err)
		return
	}

	// Menjalankan migration UP
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("[Migration] Failed to execute SQL migration: %v", err)
	} else if err == migrate.ErrNoChange {
		log.Println("[Migration] No new SQL migration files to execute.")
	} else {
		log.Println("[Migration] New SQL migration files successfully executed.")
	}
}