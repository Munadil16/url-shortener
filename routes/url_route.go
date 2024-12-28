package routes

import (
	"github.com/Munadil16/url-shortener-server/handlers"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	urlRouter := r.Group("/api/v1/url")

	urlRouter.POST("/shorten", handlers.Shorten)

	urlRouter.GET("/fetch/:id", handlers.Redirect)
}
