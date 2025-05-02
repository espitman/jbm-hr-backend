package albumhandler

import (
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/albumservice"
	"github.com/labstack/echo/v4"
)

// AlbumAdminHandler handles HTTP requests for album admin operations
type AlbumAdminHandler struct {
	service albumservice.Service
}

// NewAlbumAdminHandler creates a new AlbumAdminHandler
func NewAlbumAdminHandler(service albumservice.Service) *AlbumAdminHandler {
	return &AlbumAdminHandler{
		service: service,
	}
}

// CreateAlbum godoc
// @Summary Create a new album
// @Description Create a new album with the provided details (Admin only)
// @Tags albums - admin
// @Accept json
// @Produce json
// @Param album body CreateAlbumRequest true "Album creation details"
// @Success 201 {object} AlbumResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/albums [post]
func (h *AlbumAdminHandler) CreateAlbum(c echo.Context) error {
	var req CreateAlbumRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	input := contract.CreateAlbumInput{
		URL:     req.URL,
		Caption: req.Caption,
	}

	album, err := h.service.CreateAlbum(c.Request().Context(), &input)
	if err != nil {
		return dto.InternalServerErrorJSON(c, err.Error())
	}

	return dto.CreatedJSON(c, album)
}

// UpdateAlbum godoc
// @Summary Update an album
// @Description Update an album with the provided details (Admin only)
// @Tags albums - admin
// @Accept json
// @Produce json
// @Param id path int true "Album ID"
// @Param album body UpdateAlbumRequest true "Album update details"
// @Success 200 {object} AlbumResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/albums/{id} [put]
func (h *AlbumAdminHandler) UpdateAlbum(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "Invalid album ID")
	}

	var req UpdateAlbumRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	input := contract.UpdateAlbumInput{
		URL:     req.URL,
		Caption: req.Caption,
	}

	album, err := h.service.UpdateAlbum(c.Request().Context(), id, &input)
	if err != nil {
		return dto.InternalServerErrorJSON(c, err.Error())
	}

	return dto.SuccessJSON(c, album)
}

// DeleteAlbum godoc
// @Summary Delete an album
// @Description Delete an album by ID (Admin only)
// @Tags albums - admin
// @Accept json
// @Produce json
// @Param id path int true "Album ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/albums/{id} [delete]
func (h *AlbumAdminHandler) DeleteAlbum(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "Invalid album ID")
	}

	err = h.service.DeleteAlbum(c.Request().Context(), id)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "Failed to delete album")
	}
	return dto.SuccessJSON(c, "Album deleted successfully")
}
