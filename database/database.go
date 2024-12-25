package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL is not defined")
	}

	connection, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	}

	DB = connection
}

func AutoMigrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
