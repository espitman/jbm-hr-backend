package infohandler

import (
	"net/http"

	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/infoservice"
	"github.com/labstack/echo/v4"
)

// InfoHandler handles HTTP requests for info operations
type InfoHandler struct {
	infoService infoservice.Service
}

// NewInfoHandler creates a new InfoHandler
func NewInfoHandler(infoService infoservice.Service) *InfoHandler {
	return &InfoHandler{
		infoService: infoService,
	}
}

// GetDashboardInfo handles getting dashboard information
// @Summary Get dashboard information
// @Description Get counts of users, requests, resumes, and departments
// @Tags info - admin
// @Accept json
// @Produce json
// @Success 200 {object} DashboardInfoResponse
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/info/dashboard [get]
func (h *InfoHandler) GetDashboardInfo(c echo.Context) error {
	info, err := h.infoService.GetDashboardInfo(c.Request().Context())
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, &info)
}
