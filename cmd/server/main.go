package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zumlabs/go-auth-api/config"
	"github.com/zumlabs/go-auth-api/internal/handler"
	"github.com/zumlabs/go-auth-api/internal/model"
	"github.com/zumlabs/go-auth-api/internal/repository"
	"github.com/zumlabs/go-auth-api/internal/service"
)

func main() {
	// load config
	cfg := config.LoadConfig()

	// connect database
	db := config.ConnectDatabase(cfg)

	// auto migrate table
	db.AutoMigrate(&model.User{})

	// init gin
	r := gin.Default()

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// ===== DEPENDENCY INJECTION =====
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// ===== ROUTES =====
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
	}

	// run server
	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
