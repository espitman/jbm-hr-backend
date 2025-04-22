package albumhandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/services/albumservice"
	"github.com/gin-gonic/gin"
)

// AlbumHandler handles HTTP requests for albums
type AlbumHandler struct {
	service *albumservice.AlbumService
}

// New creates a new AlbumHandler
func New(service *albumservice.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		service: service,
	}
}

// GetAllAlbums handles GET /api/v1/albums
func (h *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := h.service.GetAllAlbums(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumsResponse{
			Success: false,
			Message: "Failed to get albums",
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumsResponse{
		Success: true,
		Data:    albums,
	})
}

// CreateAlbum handles POST /api/v1/albums
func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var req contract.CreateAlbumInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	album, err := h.service.CreateAlbum(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Success: false,
			Message: "Failed to create album",
		})
		return
	}
	c.JSON(http.StatusCreated, &AlbumResponse{
		Success: true,
		Message: "Album created successfully",
		Data:    album,
	})
}

// GetAlbumByID handles GET /api/v1/albums/:id
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Success: false,
			Message: "Invalid album ID",
		})
		return
	}

	album, err := h.service.GetAlbumByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Success: false,
			Message: "Failed to get album",
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumResponse{
		Success: true,
		Data:    album,
	})
}

// UpdateAlbum handles PUT /api/v1/albums/:id
func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Success: false,
			Message: "Invalid album ID",
		})
		return
	}

	var req contract.UpdateAlbumInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	album, err := h.service.UpdateAlbum(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Success: false,
			Message: "Failed to update album",
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumResponse{
		Success: true,
		Message: "Album updated successfully",
		Data:    album,
	})
}

// DeleteAlbum handles DELETE /api/v1/albums/:id
func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Success: false,
			Message: "Invalid album ID",
		})
		return
	}

	err = h.service.DeleteAlbum(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Success: false,
			Message: "Failed to delete album",
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumResponse{
		Success: true,
		Message: "Album deleted successfully",
	})
}
