package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zumlabs/go-auth-api/config"
	"github.com/zumlabs/go-auth-api/internal/model"
)

func main() {
	cfg := config.LoadConfig()
	db := config.ConnectDatabase(cfg)

	// auto create users table
	db.AutoMigrate(&model.User{})

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
