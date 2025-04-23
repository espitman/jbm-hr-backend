package main

import (
	"context"
	"log"
	"net/http"

	"github.com/espitman/jbm-hr-backend/database"
	"github.com/espitman/jbm-hr-backend/database/repository/album"
	_ "github.com/espitman/jbm-hr-backend/docs"
	"github.com/espitman/jbm-hr-backend/http/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/http/router"
	"github.com/espitman/jbm-hr-backend/service/albumservice"
	"github.com/espitman/jbm-hr-backend/utils/config"
)

// @title           JBM HR Backend API
// @version         1.01
// @description     This is the backend API for JBM HR system.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  s.heidari@jabama.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	// Load environment variables
	config.LoadEnv()

	// Get port from environment variables
	port := config.GetConfig("PORT", "8080")

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
	log.Printf("Server starting on port %s...", port)
	if err := r.Start(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed starting server: %v", err)
	}
}
