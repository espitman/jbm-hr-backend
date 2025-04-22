package router

import (
	"gin-project/handlers/albumhandler"
	"gin-project/middleware"

	"github.com/gin-gonic/gin"
)

// Router holds the gin engine and handlers
type Router struct {
	engine       *gin.Engine
	albumHandler *albumhandler.AlbumHandler
}

// NewRouter creates a new router instance
func NewRouter(albumHandler *albumhandler.AlbumHandler) *Router {
	engine := gin.Default()
	engine.Use(middleware.Logger())

	return &Router{
		engine:       engine,
		albumHandler: albumHandler,
	}
}

// SetupRoutes sets up all the routes in the application
func (r *Router) SetupRoutes() {
	// Register base routes
	r.registerBaseRoutes()

	// Create API v1 group
	apiV1 := r.engine.Group("/api/v1")

	// Register routes
	r.registerAlbumRoutes(apiV1)
}

// registerBaseRoutes registers all base routes
func (r *Router) registerBaseRoutes() {
	r.engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin API",
		})
	})
}

// registerAlbumRoutes registers all album-related routes
func (r *Router) registerAlbumRoutes(group *gin.RouterGroup) {
	albums := group.Group("/albums")
	{
		albums.GET("", r.albumHandler.GetAlbums)
		albums.POST("", r.albumHandler.CreateAlbum)
		albums.GET("/:id", r.albumHandler.GetAlbum)
		albums.PUT("/:id", r.albumHandler.UpdateAlbum)
		albums.DELETE("/:id", r.albumHandler.DeleteAlbum)
	}
}

// GetEngine returns the gin engine
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
