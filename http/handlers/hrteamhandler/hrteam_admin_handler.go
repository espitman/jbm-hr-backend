package hrteamhandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/hrteamservice"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

// HRTeamAdminHandler handles HTTP requests for HR team administration
type HRTeamAdminHandler struct {
	hrteamService hrteamservice.Service
}

// NewHRTeamAdminHandler creates a new HRTeamAdminHandler
func NewHRTeamAdminHandler(hrteamService hrteamservice.Service) *HRTeamAdminHandler {
	return &HRTeamAdminHandler{
		hrteamService: hrteamService,
	}
}

// Create handles the creation of a new HR team member (admin only)
// @Summary Create a new HR team member (admin only)
// @Description Create a new HR team member in the system (admin only)
// @Tags admin - hr-team
// @Accept json
// @Produce json
// @Param request body HRTeamInput true "Create HR Team Member"
// @Success 201 {object} CreateHRTeamResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/hr-team [post]
func (h *HRTeamAdminHandler) Create(c echo.Context) error {
	var req HRTeamInput
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	input := &contract.HRTeamInput{
		FullName: req.FullName,
		Role:     req.Role,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	hrTeam, err := h.hrteamService.CreateHRTeam(c.Request().Context(), input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.CreatedJSON(c, hrTeam)
}

// Update handles the update of an existing HR team member (admin only)
// @Summary Update an HR team member (admin only)
// @Description Update an existing HR team member in the system (admin only)
// @Tags admin - hr-team
// @Accept json
// @Produce json
// @Param id path int true "HR Team Member ID"
// @Param request body HRTeamInput true "Update HR Team Member"
// @Success 200 {object} UpdateHRTeamResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/hr-team/{id} [put]
func (h *HRTeamAdminHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid hr team member id")
	}

	var req HRTeamInput
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	input := &contract.HRTeamInput{
		FullName: req.FullName,
		Role:     req.Role,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	hrTeam, err := h.hrteamService.UpdateHRTeam(c.Request().Context(), id, input)
	if err != nil {
		if err == contract.ErrHRTeamNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, hrTeam)
}

// Delete handles the deletion of an HR team member (admin only)
// @Summary Delete an HR team member (admin only)
// @Description Delete an HR team member by their ID (admin only)
// @Tags admin - hr-team
// @Accept json
// @Produce json
// @Param id path int true "HR Team Member ID"
// @Success 200 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/hr-team/{id} [delete]
func (h *HRTeamAdminHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid hr team member id")
	}

	err = h.hrteamService.DeleteHRTeam(c.Request().Context(), id)
	if err != nil {
		if err == contract.ErrHRTeamNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, nil)
}
