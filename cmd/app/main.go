package main

import (
	"go_project/docs"
	"go_project/internal/db"
	"go_project/internal/logger"
	"go_project/internal/server"
	"go_project/internal/user"
)

func main() {
	// Swagger metadata
	docs.SwaggerInfo.Title = "User API"
	docs.SwaggerInfo.Description = "Create and get user details"
	docs.SwaggerInfo.Version = "1.0"

	// Initialize DB
	dbConn, err := db.New("test.db")
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect database")
	}

	// Auto migrate
	if err := dbConn.AutoMigrate(&user.User{}); err != nil {
		logger.Log.Error().Err(err).Msg("AutoMigrate failed")
	}

	// Initialize services
	srv := server.SetupServer(dbConn)

	// Run API
	if err := srv.Run(":8080"); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to run server")
	}
}
