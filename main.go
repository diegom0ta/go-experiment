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
	defer func() {
		if err := cache.Close(); err != nil {
			logger.Errorf("failed to close redis connection: %v", err)
		}
	}()

	logger.Info("Redis connection established")

	// Create server and router
	srv := server.NewServer()
	routers := router.NewRouter(srv)

	// Setup routes first
	routers.Start()
	logger.Info("Routes configured")

	// Then start the server (this will block)
	srv.Start()
}
