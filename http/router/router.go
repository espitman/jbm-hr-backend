package router

import (
	"net/http"
	"path/filepath"

	_ "github.com/espitman/jbm-hr-backend/docs" // This is important for Swagger
	"github.com/espitman/jbm-hr-backend/http/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/fronthandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/userhandler"
	customMiddleware "github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Router holds the echo instance and handlers
type Router struct {
	*echo.Echo
	albumHandler      *albumhandler.AlbumHandler
	albumAdminHandler *albumhandler.AlbumAdminHandler
	userHandler       *userhandler.UserHandler
	frontHandler      *fronthandler.FrontHandler
}

// NewRouter creates a new router instance
func NewRouter(albumHandler *albumhandler.AlbumHandler, albumAdminHandler *albumhandler.AlbumAdminHandler, userHandler *userhandler.UserHandler) *Router {
	e := echo.New()
	e.Use(customMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())

	// Get the absolute path to the frontend directory
	frontendPath, _ := filepath.Abs("frontend")

	return &Router{
		Echo:              e,
		albumHandler:      albumHandler,
		albumAdminHandler: albumAdminHandler,
		userHandler:       userHandler,
		frontHandler:      fronthandler.NewFrontHandler(frontendPath),
	}
}

// SetupRoutes sets up all the routes in the application
func (r *Router) SetupRoutes() {
	// Serve frontend static files first
	r.GET("/*", r.frontHandler.ServeFrontend)

	// Create API v1 group
	apiV1 := r.Group("/api/v1")

	// Admin API v1 group with admin middleware
	apiV1Admin := apiV1.Group("/admin")
	apiV1Admin.Use(customMiddleware.JWT())
	apiV1Admin.Use(customMiddleware.Admin())

	// Register routes
	r.registerAlbumRoutes(apiV1)
	r.registerAlbumAdminRoutes(apiV1Admin)
	r.registerUserRoutes(apiV1)
	r.registerUserAdminRoutes(apiV1Admin)

	// Add Swagger
	r.GET("/swagger/*", echoSwagger.WrapHandler)
}

// GetEcho returns the echo instance
func (r *Router) GetEcho() *echo.Echo {
	return r.Echo
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Echo.ServeHTTP(w, req)
}
