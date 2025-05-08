package digikalacodehandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/digikalacodeservice"
	"github.com/labstack/echo/v4"
)

// DigikalaCodeAdminHandler handles HTTP requests for Digikala codes by admins
type DigikalaCodeAdminHandler struct {
	digikalaCodeService digikalacodeservice.Service
}

// NewDigikalaCodeAdminHandler creates a new DigikalaCodeAdminHandler
func NewDigikalaCodeAdminHandler(digikalaCodeService digikalacodeservice.Service) *DigikalaCodeAdminHandler {
	return &DigikalaCodeAdminHandler{
		digikalaCodeService: digikalaCodeService,
	}
}

// Create handles creating a new Digikala code
// @Summary Create a Digikala code
// @Description Create a new Digikala code
// @Tags digikala-codes - admin
// @Accept json
// @Produce json
// @Param input body CreateDigikalaCodeRequest true "Create Digikala code input"
// @Success 200 {object} CreateDigikalaCodeResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/digikala-codes [post]
func (h *DigikalaCodeAdminHandler) Create(c echo.Context) error {
	var req CreateDigikalaCodeRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid input")
	}

	input := &contract.CreateDigikalaCodeInput{
		Code: req.Code,
	}

	digikalaCode, err := h.digikalaCodeService.Create(c.Request().Context(), input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, digikalaCode)
}

// List handles retrieving all Digikala codes
// @Summary List Digikala codes
// @Description Get all Digikala codes
// @Tags digikala-codes - admin
// @Accept json
// @Produce json
// @Success 200 {object} ListDigikalaCodeResponse
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/digikala-codes [get]
func (h *DigikalaCodeAdminHandler) List(c echo.Context) error {
	digikalaCodes, err := h.digikalaCodeService.GetAll(c.Request().Context())
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	// Convert []*contract.DigikalaCode to []contract.DigikalaCode and remove code field
	convertedCodes := make([]contract.DigikalaCode, len(digikalaCodes))
	for i, code := range digikalaCodes {
		convertedCode := *code
		convertedCode.Code = "" // Remove code from response
		convertedCodes[i] = convertedCode
	}

	return dto.SuccessJSON(c, DigikalaCodeListData{
		Items: convertedCodes,
	})
}

// Get handles retrieving a Digikala code by ID
// @Summary Get a Digikala code
// @Description Get a Digikala code by ID
// @Tags digikala-codes - admin
// @Accept json
// @Produce json
// @Param id path int true "Digikala code ID"
// @Success 200 {object} GetDigikalaCodeResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/digikala-codes/{id} [get]
func (h *DigikalaCodeAdminHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid id")
	}

	digikalaCode, err := h.digikalaCodeService.GetByID(c.Request().Context(), id)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	if digikalaCode == nil {
		return dto.ErrorJSON(c, http.StatusNotFound, "digikala code not found")
	}

	// Remove code from response
	digikalaCode.Code = ""

	return dto.SuccessJSON(c, digikalaCode)
}

// Assign handles assigning a Digikala code to a user
// @Summary Assign a Digikala code
// @Description Assign a Digikala code to a user
// @Tags digikala-codes - admin
// @Accept json
// @Produce json
// @Param id path int true "Digikala code ID"
// @Param input body AssignDigikalaCodeRequest true "Assign Digikala code input"
// @Success 200 {object} AssignDigikalaCodeResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/digikala-codes/{id}/assign [put]
func (h *DigikalaCodeAdminHandler) Assign(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid id")
	}

	var req AssignDigikalaCodeRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid input")
	}

	input := &contract.AssignDigikalaCodeInput{
		ID:     id,
		UserID: req.UserID,
	}

	updatedDigikalaCode, err := h.digikalaCodeService.Assign(c.Request().Context(), "", input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, updatedDigikalaCode)
}
