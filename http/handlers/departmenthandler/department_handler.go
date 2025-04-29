package departmenthandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/departmentservice"
	"github.com/labstack/echo/v4"
)

// DepartmentHandler handles HTTP requests for departments
type DepartmentHandler struct {
	departmentService departmentservice.Service
}

// NewDepartmentHandler creates a new DepartmentHandler
func NewDepartmentHandler(departmentService departmentservice.Service) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: departmentService,
	}
}

// Get handles retrieving a department by ID
// @Summary Get a department
// @Description Get a department by its ID
// @Tags departments
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} GetDepartmentResponse
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/departments/{id} [get]
func (h *DepartmentHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid department id")
	}

	department, err := h.departmentService.GetByID(c.Request().Context(), id)
	if err != nil {
		if err == contract.ErrDepartmentNotFound {
			return dto.ErrorJSON(c, http.StatusNotFound, err.Error())
		}
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, DepartmentData{
		ID:          department.ID,
		Title:       department.Title,
		Description: department.Description,
		Image:       department.Image,
		Icon:        department.Icon,
		Color:       department.Color,
		ShortName:   department.ShortName,
	})
}

// List handles retrieving a paginated list of departments
// @Summary List departments
// @Description Get a paginated list of departments
// @Tags departments
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} ListDepartmentsResponse
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/departments [get]
func (h *DepartmentHandler) List(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 10
	}

	departments, total, err := h.departmentService.List(c.Request().Context(), page, limit)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	data := make([]DepartmentData, len(departments))
	for i, department := range departments {
		data[i] = DepartmentData{
			ID:          department.ID,
			Title:       department.Title,
			Description: department.Description,
			Image:       department.Image,
			Icon:        department.Icon,
			Color:       department.Color,
			ShortName:   department.ShortName,
		}
	}

	return dto.SuccessJSON(c, struct {
		Data  []DepartmentData `json:"data"`
		Total int              `json:"total"`
	}{
		Data:  data,
		Total: total,
	})
}
