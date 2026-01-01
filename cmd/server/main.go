package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zumlabs/go-auth-api/config"
)

func main() {
	cfg := config.LoadConfig()
	db := config.ConnectDatabase(cfg)
	_ = db // dipakai nanti, biar tidak error

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
