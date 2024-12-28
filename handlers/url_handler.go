package handlers

import (
	"strings"

	"github.com/Munadil16/url-shortener-server/database"
	"github.com/Munadil16/url-shortener-server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShortenBody struct {
	Url string `json:"url"`
}

func Shorten(c *gin.Context) {
	var body ShortenBody
	c.BindJSON(&body)

	if body.Url == "" {
		c.JSON(400, gin.H{"message": "Url is required"})
		return
	}

	if !strings.HasPrefix(body.Url, "https://") {
		c.JSON(400, gin.H{"message": "Invalid Url"})
		return
	}

	shortId := uuid.NewString()[:7]

	url := models.Url{LongUrl: body.Url, Visits: 0, ShortUrlId: shortId}
	res := database.DB.Create(&url)

	if res.Error != nil {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(201, gin.H{"message": "Short url is created successfully", "shortId": shortId})

}

func Redirect(c *gin.Context) {
	id := c.Param("id")
	var url models.Url

	res := database.DB.Where("short_url_id = ?", id).First(&url)

	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"message": "Url not found"})
		} else {
			c.JSON(500, gin.H{"message": "Internal server error"})
		}

		return
	}

	database.DB.Model(&url).Update("visits", gorm.Expr("visits + ?", 1))

	c.Redirect(307, url.LongUrl)
}
