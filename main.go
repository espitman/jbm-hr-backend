package main

import (
	"context"
	"log"
	"net/http"

	"github.com/espitman/jbm-hr-backend/database"
	_ "github.com/espitman/jbm-hr-backend/docs" // This will be generated
	"github.com/espitman/jbm-hr-backend/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/repository/album"
	"github.com/espitman/jbm-hr-backend/router"
	"github.com/espitman/jbm-hr-backend/services/albumservice"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           JBM HR Backend API
// @version         1.0
// @description     This is the backend API for JBM HR system.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

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

	// Add Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
