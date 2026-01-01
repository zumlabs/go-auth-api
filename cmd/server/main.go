package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zumlabs/go-auth-api/config"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Go Auth API is running",
		})
	})

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
