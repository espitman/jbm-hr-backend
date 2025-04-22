package albumhandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/handlers/dto"
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

// @Summary     Get all albums
// @Description Retrieves a list of all albums
// @Tags        albums
// @Produce     json
// @Success     200 {object} AlbumsResponse
// @Failure     500 {object} AlbumsResponse
// @Router      /albums [get]
func (h *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := h.service.GetAllAlbums(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumsResponse{
			Response: dto.Response{
				Success: false,
				Message: "Failed to get albums",
			},
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumsResponse{
		Response: dto.Response{
			Success: true,
		},
		Data: albums,
	})
}

// @Summary     Create a new album
// @Description Creates a new album with the provided details
// @Tags        albums
// @Accept      json
// @Produce     json
// @Param       album body contract.CreateAlbumInput true "Album details"
// @Success     201 {object} AlbumResponse
// @Failure     400 {object} AlbumResponse
// @Failure     500 {object} AlbumResponse
// @Router      /albums [post]
func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var req contract.CreateAlbumInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Invalid request body",
			},
		})
		return
	}

	album, err := h.service.CreateAlbum(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Failed to create album",
			},
		})
		return
	}
	c.JSON(http.StatusCreated, &AlbumResponse{
		Response: dto.Response{
			Success: true,
			Message: "Album created successfully",
		},
		Data: album,
	})
}

// @Summary     Get album by ID
// @Description Retrieves an album by its ID
// @Tags        albums
// @Produce     json
// @Param       id  path     int  true  "Album ID"
// @Success     200 {object} AlbumResponse
// @Failure     400 {object} AlbumResponse
// @Failure     500 {object} AlbumResponse
// @Router      /albums/{id} [get]
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Invalid album ID",
			},
		})
		return
	}

	album, err := h.service.GetAlbumByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Failed to get album",
			},
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumResponse{
		Response: dto.Response{
			Success: true,
		},
		Data: album,
	})
}

// @Summary     Update album
// @Description Updates an existing album with the provided details
// @Tags        albums
// @Accept      json
// @Produce     json
// @Param       id   path     int                    true "Album ID"
// @Param       album body    contract.UpdateAlbumInput true "Album details"
// @Success     200  {object} AlbumResponse
// @Failure     400  {object} AlbumResponse
// @Failure     500  {object} AlbumResponse
// @Router      /albums/{id} [put]
func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Invalid album ID",
			},
		})
		return
	}

	var req contract.UpdateAlbumInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Invalid request body",
			},
		})
		return
	}

	album, err := h.service.UpdateAlbum(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Failed to update album",
			},
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumResponse{
		Response: dto.Response{
			Success: true,
			Message: "Album updated successfully",
		},
		Data: album,
	})
}

// @Summary     Delete album
// @Description Deletes an album by its ID
// @Tags        albums
// @Produce     json
// @Param       id  path     int  true  "Album ID"
// @Success     200 {object} AlbumResponse
// @Failure     400 {object} AlbumResponse
// @Failure     500 {object} AlbumResponse
// @Router      /albums/{id} [delete]
func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Invalid album ID",
			},
		})
		return
	}

	err = h.service.DeleteAlbum(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Failed to delete album",
			},
		})
		return
	}
	c.JSON(http.StatusOK, &AlbumResponse{
		Response: dto.Response{
			Success: true,
			Message: "Album deleted successfully",
		},
	})
}
