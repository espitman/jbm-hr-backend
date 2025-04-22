package router

import (
	"github.com/espitman/jbm-hr-backend/handlers/albumhandler"
	"github.com/espitman/jbm-hr-backend/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Router holds the gin engine and handlers
type Router struct {
	*gin.Engine
	albumHandler *albumhandler.AlbumHandler
}

// NewRouter creates a new router instance
func NewRouter(albumHandler *albumhandler.AlbumHandler) *Router {
	engine := gin.Default()
	engine.Use(middleware.Logger())

	return &Router{
		Engine:       engine,
		albumHandler: albumHandler,
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

	// Add Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// GetEngine returns the gin engine
func (r *Router) GetEngine() *gin.Engine {
	return r.Engine
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Engine.ServeHTTP(w, req)
}
