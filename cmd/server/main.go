package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zumlabs/go-auth-api/config"
	"github.com/zumlabs/go-auth-api/internal/handler"
	"github.com/zumlabs/go-auth-api/internal/middleware"
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
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler()

	// ===== AUTH ROUTES =====
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// ===== PROTECTED ROUTES =====
	user := r.Group("/api/user")
	user.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
	{
		user.GET("/me", userHandler.Me)
	}

	// run server
	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
