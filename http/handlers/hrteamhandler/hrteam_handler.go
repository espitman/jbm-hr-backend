package hrteamhandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/hrteamservice"
	"github.com/labstack/echo/v4"
)

// HRTeamHandler handles HTTP requests for HR team members
type HRTeamHandler struct {
	hrteamService hrteamservice.Service
}

// NewHRTeamHandler creates a new HRTeamHandler
func NewHRTeamHandler(hrteamService hrteamservice.Service) *HRTeamHandler {
	return &HRTeamHandler{
		hrteamService: hrteamService,
	}
}

// Get handles retrieving an HR team member by ID
// @Summary Get an HR team member
// @Description Get an HR team member by their ID
// @Tags hr-team
// @Accept json
// @Produce json
// @Param id path int true "HR Team Member ID"
// @Success 200 {object} GetHRTeamResponse
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/hr-team/{id} [get]
func (h *HRTeamHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid hr team member id")
	}

	hrTeam, err := h.hrteamService.GetHRTeamByID(c.Request().Context(), id)
	if err != nil {
		if err == contract.ErrHRTeamNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, hrTeam)
}

// List handles retrieving all HR team members
// @Summary List HR team members
// @Description Get all HR team members
// @Tags hr-team
// @Accept json
// @Produce json
// @Success 200 {object} HRTeamListResponse
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/hr-team [get]
func (h *HRTeamHandler) List(c echo.Context) error {
	hrTeams, err := h.hrteamService.GetAllHRTeam(c.Request().Context())
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, hrTeams)
}
