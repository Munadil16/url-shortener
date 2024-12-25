package main

import (
	"log"
	"os"

	"github.com/Munadil16/url-shortener-server/database"
	"github.com/Munadil16/url-shortener-server/models"
	"github.com/Munadil16/url-shortener-server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT is not defined")
	}

	database.ConnectDatabase()
	if err := database.AutoMigrate(&models.Url{}); err != nil {
		log.Fatal("Failed to migrate database")
	}

	r := gin.Default()

	routes.Router(r)

	r.Run(":" + PORT)
}
