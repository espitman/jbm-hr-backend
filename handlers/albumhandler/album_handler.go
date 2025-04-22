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
	response, err := h.service.GetAllAlbums(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateAlbum handles POST /api/v1/albums
func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var req contract.CreateAlbumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &contract.AlbumResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	response, err := h.service.CreateAlbum(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// GetAlbumByID handles GET /api/v1/albums/:id
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.AlbumResponse{
			Success: false,
			Message: "Invalid album ID",
		})
		return
	}

	response, err := h.service.GetAlbumByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// UpdateAlbum handles PUT /api/v1/albums/:id
func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.AlbumResponse{
			Success: false,
			Message: "Invalid album ID",
		})
		return
	}

	var req contract.UpdateAlbumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &contract.AlbumResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	response, err := h.service.UpdateAlbum(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteAlbum handles DELETE /api/v1/albums/:id
func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.AlbumResponse{
			Success: false,
			Message: "Invalid album ID",
		})
		return
	}

	response, err := h.service.DeleteAlbum(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
