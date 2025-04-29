package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/espitman/jbm-hr-backend/database"
	"github.com/espitman/jbm-hr-backend/database/repository/album"
	"github.com/espitman/jbm-hr-backend/database/repository/department"
	"github.com/espitman/jbm-hr-backend/database/repository/otp"
	"github.com/espitman/jbm-hr-backend/database/repository/user"
	"github.com/espitman/jbm-hr-backend/ent/migrate"
	"github.com/espitman/jbm-hr-backend/http/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/departmenthandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/fronthandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/userhandler"
	"github.com/espitman/jbm-hr-backend/http/router"
	"github.com/espitman/jbm-hr-backend/service/albumservice"
	"github.com/espitman/jbm-hr-backend/service/departmentservice"
	"github.com/espitman/jbm-hr-backend/service/userservice"
	"github.com/espitman/jbm-hr-backend/utils/config"
	_ "github.com/swaggo/files"
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

// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize repository
	albumRepo := album.NewEntRepository(client)
	userRepo := user.NewEntRepository(client)
	otpRepo := otp.NewEntRepository(client)
	departmentRepo := department.NewEntRepository(client)

	// Initialize service
	albumService := albumservice.New(albumRepo)
	userService := userservice.New(userRepo, otpRepo)
	departmentService := departmentservice.New(departmentRepo)

	// Initialize handlers
	albumHandler := albumhandler.NewAlbumHandler(albumService)
	albumAdminHandler := albumhandler.NewAlbumAdminHandler(albumService)
	userHandler := userhandler.NewUserHandler(userService)
	departmentHandler := departmenthandler.NewDepartmentHandler(departmentService)
	departmentAdminHandler := departmenthandler.NewDepartmentAdminHandler(departmentService)

	// Initialize front handler
	frontendPath, _ := filepath.Abs("ui/web")
	frontHandler := fronthandler.NewFrontHandler(frontendPath)

	// Initialize router
	r := router.NewRouter(albumHandler, albumAdminHandler, userHandler, departmentHandler, departmentAdminHandler, frontHandler)
	r.SetupRoutes()

	// Start server
	log.Printf("Server starting on port %s...", port)
	if err := r.Start(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed starting server: %v", err)
	}
}
