package albumhandler

import (
	"net/http"
	"strconv"

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
	albums, err := h.albumService.GetAllAlbums(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

// CreateAlbum creates a new album
func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var input struct {
		URL     string `json:"url" binding:"required"`
		Caption string `json:"caption"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := h.albumService.CreateAlbum(c.Request.Context(), input.URL, input.Caption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, album)
}

// GetAlbum returns an album by ID
func (h *AlbumHandler) GetAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	album, err := h.albumService.GetAlbumByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

// UpdateAlbum updates an existing album
func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input struct {
		URL     string `json:"url" binding:"required"`
		Caption string `json:"caption"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := h.albumService.UpdateAlbum(c.Request.Context(), id, input.URL, input.Caption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

// DeleteAlbum deletes an album
func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.albumService.DeleteAlbum(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
