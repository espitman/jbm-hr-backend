package alibabacodehandler

import (
	"net/http"
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/alibabacodeservice"
	"github.com/labstack/echo/v4"
)

// AlibabaCodeAdminHandler handles HTTP requests for Alibaba codes by admins
type AlibabaCodeAdminHandler struct {
	alibabaCodeService alibabacodeservice.Service
}

// NewAlibabaCodeAdminHandler creates a new AlibabaCodeAdminHandler
func NewAlibabaCodeAdminHandler(alibabaCodeService alibabacodeservice.Service) *AlibabaCodeAdminHandler {
	return &AlibabaCodeAdminHandler{
		alibabaCodeService: alibabaCodeService,
	}
}

// Create handles creating a new Alibaba code
// @Summary Create an Alibaba code
// @Description Create a new Alibaba code
// @Tags alibaba-codes - admin
// @Accept json
// @Produce json
// @Param input body CreateAlibabaCodeRequest true "Create Alibaba code input"
// @Success 200 {object} CreateAlibabaCodeResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/alibaba-codes [post]
func (h *AlibabaCodeAdminHandler) Create(c echo.Context) error {
	var req CreateAlibabaCodeRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid input")
	}

	input := &contract.CreateAlibabaCodeInput{
		Code: req.Code,
		Type: req.Type,
	}

	alibabaCode, err := h.alibabaCodeService.Create(c.Request().Context(), input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, alibabaCode)
}

// List handles the request to get all Alibaba codes
// @Summary List all Alibaba codes
// @Description Get a list of all Alibaba codes
// @Tags alibaba-codes - admin
// @Accept json
// @Produce json
// @Param used query bool false "Filter by used status"
// @Param assign_to_user_id query int false "Filter by assigned user ID"
// @Param type query string false "Filter by code type"
// @Success 200 {object} ListAlibabaCodeResponse
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/alibaba-codes [get]
func (h *AlibabaCodeAdminHandler) List(c echo.Context) error {
	filters := &contract.AlibabaCodeFilters{}

	// Parse query parameters
	if usedStr := c.QueryParam("used"); usedStr != "" {
		used, err := strconv.ParseBool(usedStr)
		if err == nil {
			filters.Used = &used
		}
	}

	if assignToUserIDStr := c.QueryParam("assign_to_user_id"); assignToUserIDStr != "" {
		if id, err := strconv.Atoi(assignToUserIDStr); err == nil {
			filters.AssignToUserID = &id
		}
	}

	if codeType := c.QueryParam("type"); codeType != "" {
		filters.Type = &codeType
	}

	codes, err := h.alibabaCodeService.GetAll(c.Request().Context(), filters)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	// Convert []*contract.AlibabaCode to []contract.AlibabaCode and remove code field
	convertedCodes := make([]contract.AlibabaCode, len(codes))
	for i, code := range codes {
		convertedCode := *code
		convertedCode.Code = "" // Remove code from response
		convertedCodes[i] = convertedCode
	}

	return dto.SuccessJSON(c, AlibabaCodeListData{
		Items: convertedCodes,
		Total: len(convertedCodes),
	})
}

// Get handles retrieving an Alibaba code by ID
// @Summary Get an Alibaba code
// @Description Get an Alibaba code by ID
// @Tags alibaba-codes - admin
// @Accept json
// @Produce json
// @Param id path int true "Alibaba code ID"
// @Success 200 {object} GetAlibabaCodeResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/alibaba-codes/{id} [get]
func (h *AlibabaCodeAdminHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid id")
	}

	alibabaCode, err := h.alibabaCodeService.GetByID(c.Request().Context(), id)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	if alibabaCode == nil {
		return dto.ErrorJSON(c, http.StatusNotFound, "alibaba code not found")
	}

	// Remove code from response
	alibabaCode.Code = ""

	return dto.SuccessJSON(c, alibabaCode)
}

// Assign handles assigning an Alibaba code to a user
// @Summary Assign an Alibaba code
// @Description Assign an Alibaba code to a user
// @Tags alibaba-codes - admin
// @Accept json
// @Produce json
// @Param id path int true "Alibaba code ID"
// @Param input body AssignAlibabaCodeRequest true "Assign Alibaba code input"
// @Success 200 {object} AssignAlibabaCodeResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Security BearerAuth
// @Router /api/v1/admin/alibaba-codes/{id}/assign [put]
func (h *AlibabaCodeAdminHandler) Assign(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid id")
	}

	var req AssignAlibabaCodeRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid input")
	}

	input := &contract.UpdateAlibabaCodeInput{
		Used:           true,
		AssignToUserID: &req.UserID,
	}

	updatedAlibabaCode, err := h.alibabaCodeService.Update(c.Request().Context(), id, input)
	if err != nil {
		return dto.ErrorJSON(c, http.StatusInternalServerError, err.Error())
	}

	return dto.SuccessJSON(c, updatedAlibabaCode)
}
