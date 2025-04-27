package albumhandler

import (
	"strconv"

	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/albumservice"
	"github.com/labstack/echo/v4"
)

// AlbumHandler handles HTTP requests for album operations
type AlbumHandler struct {
	service albumservice.Service
}

// NewAlbumHandler creates a new AlbumHandler
func NewAlbumHandler(service albumservice.Service) *AlbumHandler {
	return &AlbumHandler{
		service: service,
	}
}

// GetAllAlbums godoc
// @Summary Get all albums
// @Description Get all albums
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} AlbumResponse
// @Failure 500 {object} dto.Response
// @Router /api/v1/albums [get]
func (h *AlbumHandler) GetAllAlbums(c echo.Context) error {
	albums, err := h.service.GetAllAlbums(c.Request().Context())
	if err != nil {
		return dto.InternalServerErrorJSON(c, err.Error())
	}

	return dto.SuccessJSON(c, albums)
}

// GetAlbumByID godoc
// @Summary Get an album by ID
// @Description Get an album by its ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path int true "Album ID"
// @Success 200 {object} AlbumResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/albums/{id} [get]
func (h *AlbumHandler) GetAlbumByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "Invalid album ID")
	}

	album, err := h.service.GetAlbumByID(c.Request().Context(), id)
	if err != nil {
		return dto.InternalServerErrorJSON(c, err.Error())
	}

	return dto.SuccessJSON(c, album)
}
