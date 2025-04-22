package main

import (
	"log"

	"gin-project/database"
	"gin-project/handlers/albumhandler"
	"gin-project/router"
	"gin-project/services/albumservice"
)

func main() {
	// Initialize database
	client, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Initialize services
	albumService := albumservice.New()

	// Initialize handlers with services
	albumHandler := albumhandler.New(albumService)

	// Setup router with handlers
	r := router.NewRouter(albumHandler)
	r.SetupRoutes()

	// Run the server on port 8080
	if err := r.GetEngine().Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
