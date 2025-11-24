package main

import (
	"go_project/logger"

	"go_project/db"
	"go_project/docs"
	"go_project/models"
	"go_project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	docs.SwaggerInfo.Title = "User API"
	docs.SwaggerInfo.Description = "This is application to create and get user details from sqlite."
	docs.SwaggerInfo.Version = "1.0"
	// Optional: Auto migrate the User model
	if err := db.Connection().AutoMigrate(&models.User{}); err != nil {
		// handle error appropriately, e.g., log or return
		logger.Log.Error().Msgf("AutoMigrate failed: %v", err)
	}
	// Setup Gin router
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.JsonLoggerMiddleware())
	routes.RegisterRoutes(r)
	// Start server on port 8080
	if err := r.Run(":8080"); err != nil {
		logger.Log.Error().Msgf("Failed to run server: %v", err)
	}
}
