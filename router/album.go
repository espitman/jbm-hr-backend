package router

import (
	"github.com/labstack/echo/v4"
)

// registerAlbumRoutes registers all album-related routes
func (r *Router) registerAlbumRoutes(group *echo.Group) {
	albums := group.Group("/albums")
	{
		albums.GET("", r.albumHandler.GetAllAlbums)
		albums.POST("", r.albumHandler.CreateAlbum)
		albums.GET("/:id", r.albumHandler.GetAlbumByID)
		albums.PUT("/:id", r.albumHandler.UpdateAlbum)
		albums.DELETE("/:id", r.albumHandler.DeleteAlbum)
	}
}
