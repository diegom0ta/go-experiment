package main

import (
	"experiment/infra/cache"
	"experiment/infra/database"
	"experiment/infra/logger"
	"experiment/infra/server"
	"experiment/infra/server/router"
)

func main() {
	// Initialize logger
	log := logger.GetLogger()
	log.Info("Starting application...")

	// Initialize DB connection
	if err := database.InitDB(); err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}
	logger.Info("Database connection established")

	// Initialize Redis connection
	if err := cache.InitRedis(); err != nil {
		logger.Fatalf("failed to connect to redis: %v", err)
	}
	defer cache.Close()
	logger.Info("Redis connection established")

	// Create server and router
	server := server.NewServer()
	router := router.NewRouter(server)

	// Setup routes first
	router.Start()
	logger.Info("Routes configured")

	// Then start the server (this will block)
	server.Start()
}
