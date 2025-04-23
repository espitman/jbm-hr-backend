package albumhandler

import (
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/albumservice"
	"github.com/labstack/echo/v4"
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

// GetAllAlbums godoc
// @Summary Get all albums
// @Description Retrieves a list of all albums
// @Tags albums
// @Produce json
// @Success 200 {object} AlbumsResponse
// @Failure 500 {object} dto.Response
// @Router /albums [get]
func (h *AlbumHandler) GetAllAlbums(c echo.Context) error {
	albums, err := h.service.GetAllAlbums(c.Request().Context())
	if err != nil {
		return dto.InternalServerErrorJSON(c, "Failed to retrieve albums")
	}
	return dto.SuccessJSON(c, albums)
}

// CreateAlbum godoc
// @Summary Create a new album
// @Description Create a new album with the provided details
// @Tags albums
// @Accept json
// @Produce json
// @Param album body CreateAlbumRequest true "Album creation details"
// @Success 201 {object} AlbumResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /albums [post]
func (h *AlbumHandler) CreateAlbum(c echo.Context) error {
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

// @Summary     Get album by ID
// @Description Retrieves an album by its ID
// @Tags        albums
// @Produce     json
// @Param       id  path     int  true  "Album ID"
// @Success     200 {object} AlbumResponse
// @Failure     400 {object} AlbumResponse
// @Failure     500 {object} AlbumResponse
// @Router      /albums/{id} [get]
func (h *AlbumHandler) GetAlbumByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "Invalid album ID")
	}

	album, err := h.service.GetAlbumByID(c.Request().Context(), id)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "Failed to get album")
	}
	return dto.SuccessJSON(c, album)
}

// UpdateAlbum godoc
// @Summary Update an existing album
// @Description Update an album's details by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path int true "Album ID"
// @Param album body UpdateAlbumRequest true "Album update details"
// @Success 200 {object} AlbumResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /albums/{id} [put]
func (h *AlbumHandler) UpdateAlbum(c echo.Context) error {
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

// @Summary     Delete album
// @Description Deletes an album by its ID
// @Tags        albums
// @Produce     json
// @Param       id  path     int  true  "Album ID"
// @Success     200 {object} AlbumResponse
// @Failure     400 {object} AlbumResponse
// @Failure     500 {object} AlbumResponse
// @Router      /albums/{id} [delete]
func (h *AlbumHandler) DeleteAlbum(c echo.Context) error {
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
