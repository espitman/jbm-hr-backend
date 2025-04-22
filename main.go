package main

import (
	"context"
	"log"
	"net/http"

	"github.com/espitman/jbm-hr-backend/database"
	"github.com/espitman/jbm-hr-backend/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/repository/album"
	"github.com/espitman/jbm-hr-backend/router"
	"github.com/espitman/jbm-hr-backend/services/albumservice"
)

func main() {
	// Initialize database connection
	client, err := database.NewClient()
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer client.Close()

	// Run database migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize repository
	albumRepo := album.NewEntRepository(client)

	// Initialize service
	albumService := albumservice.New(albumRepo)

	// Initialize handler
	albumHandler := albumhandler.New(albumService)

	// Initialize router
	r := router.NewRouter(albumHandler)
	r.SetupRoutes()

	// Start server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
