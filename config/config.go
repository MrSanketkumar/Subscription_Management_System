package config

import (
	"Subscription_Management_System/constants"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func NewConfig() *Config {
	dsn := "host=localhost user=postgres password=S@nket123 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(constants.ErrDB, err)
	}

	return &Config{DB: db}
}
