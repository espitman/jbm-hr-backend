package router

import (
	"github.com/gin-gonic/gin"
)

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
