package albumhandler

import (
	"net/http"

	"gin-project/services/albumservice"

	"github.com/gin-gonic/gin"
)

// AlbumHandler handles album-related routes
type AlbumHandler struct {
	albumService *albumservice.AlbumService
}

// New creates a new AlbumHandler instance with the provided service
func New(albumService *albumservice.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		albumService: albumService,
	}
}

// GetAlbums returns all albums
func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums := h.albumService.GetAllAlbums()
	c.JSON(http.StatusOK, albums)
}
