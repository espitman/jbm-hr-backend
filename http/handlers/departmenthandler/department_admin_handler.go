package departmenthandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/departmentservice"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

// DepartmentAdminHandler handles HTTP requests for department administration
type DepartmentAdminHandler struct {
	departmentService departmentservice.Service
}

// NewDepartmentAdminHandler creates a new DepartmentAdminHandler
func NewDepartmentAdminHandler(departmentService departmentservice.Service) *DepartmentAdminHandler {
	return &DepartmentAdminHandler{
		departmentService: departmentService,
	}
}

// Create handles the creation of a new department (admin only)
// @Summary Create a new department (admin only)
// @Description Create a new department in the system (admin only)
// @Tags departments - admin
// @Accept json
// @Produce json
// @Param request body CreateDepartmentRequest true "Create Department"
// @Success 201 {object} CreateDepartmentResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/departments [post]
func (h *DepartmentAdminHandler) Create(c echo.Context) error {
	var req CreateDepartmentRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	input := &contract.DepartmentInput{
		Title:        req.Title,
		Description:  req.Description,
		Image:        req.Image,
		Icon:         req.Icon,
		Color:        req.Color,
		ShortName:    req.ShortName,
		DisplayOrder: req.DisplayOrder,
	}

	department, err := h.departmentService.Create(c.Request().Context(), input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.CreatedJSON(c, department)
}

// Update handles the update of an existing department (admin only)
// @Summary Update a department (admin only)
// @Description Update an existing department in the system (admin only)
// @Tags departments - admin
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param request body UpdateDepartmentRequest true "Update Department"
// @Success 200 {object} UpdateDepartmentResponse
// @Failure 400 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/departments/{id} [put]
func (h *DepartmentAdminHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid department id")
	}

	var req UpdateDepartmentRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	input := &contract.DepartmentInput{
		Title:        req.Title,
		Description:  req.Description,
		Image:        req.Image,
		Icon:         req.Icon,
		Color:        req.Color,
		ShortName:    req.ShortName,
		DisplayOrder: req.DisplayOrder,
	}

	department, err := h.departmentService.Update(c.Request().Context(), id, input)
	if err != nil {
		if err == contract.ErrDepartmentNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, department)
}

// Delete handles the deletion of a department (admin only)
// @Summary Delete a department (admin only)
// @Description Delete a department by its ID (admin only)
// @Tags departments - admin
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/departments/{id} [delete]
func (h *DepartmentAdminHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid department id")
	}

	err = h.departmentService.Delete(c.Request().Context(), id)
	if err != nil {
		if err == contract.ErrDepartmentNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, nil)
}
