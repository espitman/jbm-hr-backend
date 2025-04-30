package requesthandler

import (
	"net/http"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/requestservice"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	requestService requestservice.Service
}

func NewHandler(requestService requestservice.Service) *Handler {
	return &Handler{
		requestService: requestService,
	}
}

// CreateRequest handles the creation of a new request
// @Summary Create a new request
// @Description Create a new request with the provided details
// @Tags requests
// @Accept json
// @Produce json
// @Param request body CreateRequestRequest true "Request details"
// @Success 201 {object} CreateRequestResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/requests [post]
func (h *Handler) CreateRequest(c echo.Context) error {
	var req CreateRequestRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request body")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	// Get user ID from JWT claims
	claims, ok := c.Get("user").(*utils.Claims)
	if !ok {
		return dto.ErrorJSON(c, http.StatusUnauthorized, "Invalid token claims")
	}
	userID := claims.UserID

	request, err := h.requestService.CreateRequest(c.Request().Context(), &contract.CreateRequestInput{
		UserID:      userID,
		FullName:    req.FullName,
		Kind:        req.Kind,
		Description: req.Description,
	})
	if err != nil {
		return dto.InternalServerErrorJSON(c, "failed to create request")
	}

	return dto.SuccessJSON(c, request)
}
