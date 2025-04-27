package router

import (
	"github.com/espitman/jbm-hr-backend/http/middleware"
	"github.com/labstack/echo/v4"
)

// registerAlbumRoutes registers all album-related routes
func (r *Router) registerAlbumRoutes(group *echo.Group) {
	albums := group.Group("/albums")
	albums.Use(middleware.JWT())
	{
		// All album routes require JWT
		albums.GET("", r.albumHandler.GetAllAlbums)
		albums.GET("/:id", r.albumHandler.GetAlbumByID)
	}
}

// registerAlbumAdminRoutes registers all admin album-related routes
func (r *Router) registerAlbumAdminRoutes(group *echo.Group) {
	albums := group.Group("/albums")
	{
		// Admin routes (protected by admin middleware)
		albums.POST("", r.albumAdminHandler.CreateAlbum)
		albums.PUT("/:id", r.albumAdminHandler.UpdateAlbum)
		albums.DELETE("/:id", r.albumAdminHandler.DeleteAlbum)
	}
}
