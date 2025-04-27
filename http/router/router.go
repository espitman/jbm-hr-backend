package router

import (
	"net/http"

	"github.com/espitman/jbm-hr-backend/http/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/http/handlers/userhandler"
	customMiddleware "github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Router holds the echo instance and handlers
type Router struct {
	*echo.Echo
	albumHandler *albumhandler.AlbumHandler
	userHandler  *userhandler.UserHandler
}

// NewRouter creates a new router instance
func NewRouter(albumHandler *albumhandler.AlbumHandler, userHandler *userhandler.UserHandler) *Router {
	e := echo.New()
	e.Use(customMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())

	return &Router{
		Echo:         e,
		albumHandler: albumHandler,
		userHandler:  userHandler,
	}
}

// SetupRoutes sets up all the routes in the application
func (r *Router) SetupRoutes() {
	// Register base routes
	r.registerBaseRoutes()

	// Create API v1 group
	apiV1 := r.Group("/api/v1")

	// Register routes
	r.registerAlbumRoutes(apiV1)
	r.registerUserRoutes(apiV1)

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
