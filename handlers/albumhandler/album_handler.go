package albumhandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/handlers/dto"
	"github.com/espitman/jbm-hr-backend/services/albumservice"
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
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("Failed to retrieve albums"))
	}
	return c.JSON(http.StatusOK, NewAlbumsResponse(albums))
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
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	input := contract.CreateAlbumInput{
		URL:     req.URL,
		Caption: req.Caption,
	}

	album, err := h.service.CreateAlbum(c.Request().Context(), &input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, NewAlbumResponse(album))
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
		return c.JSON(http.StatusBadRequest, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Invalid album ID",
			},
		})
	}

	album, err := h.service.GetAlbumByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Failed to get album",
			},
		})
	}
	return c.JSON(http.StatusOK, &AlbumResponse{
		Response: dto.Response{
			Success: true,
		},
		Data: album,
	})
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
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse("Invalid album ID"))
	}

	var req UpdateAlbumRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	input := contract.UpdateAlbumInput{
		URL:     req.URL,
		Caption: req.Caption,
	}

	album, err := h.service.UpdateAlbum(c.Request().Context(), id, &input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, NewAlbumResponse(album))
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
		return c.JSON(http.StatusBadRequest, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Invalid album ID",
			},
		})
	}

	err = h.service.DeleteAlbum(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &AlbumResponse{
			Response: dto.Response{
				Success: false,
				Message: "Failed to delete album",
			},
		})
	}
	return c.JSON(http.StatusOK, &AlbumResponse{
		Response: dto.Response{
			Success: true,
			Message: "Album deleted successfully",
		},
	})
}
