package main

import (
	"experiment/infra/database"
	"experiment/infra/server"
	"experiment/infra/server/router"
	"log"
)

func main() {
	// Initialize DB connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Start the server
	server := server.NewServer()
	server.Start()

	router := router.NewRouter(server)
	router.Start()
}
