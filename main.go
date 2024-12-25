package main

import (
	"log"
	"os"

	"github.com/Munadil16/url-shortener-server/database"
	"github.com/Munadil16/url-shortener-server/models"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from server"})
	})

	r.Run(":" + PORT)
}
