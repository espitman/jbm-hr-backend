package requesthandler

import (
	"strconv"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
	"github.com/espitman/jbm-hr-backend/service/requestservice"
	"github.com/espitman/jbm-hr-backend/utils"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	requestService requestservice.Service
}

func NewAdminHandler(requestService requestservice.Service) *AdminHandler {
	return &AdminHandler{
		requestService: requestService,
	}
}

// GetRequest handles getting a single request by ID
// @Summary Get a request by ID
// @Description Get a request by its ID
// @Tags requests - admin
// @Accept json
// @Produce json
// @Param id path int true "Request ID"
// @Success 200 {object} GetRequestResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/admin/requests/{id} [get]
func (h *AdminHandler) GetRequest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid request ID")
	}

	request, err := h.requestService.GetRequestByID(c.Request().Context(), id)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "failed to get request")
	}

	return dto.SuccessJSON(c, *request)
}

// GetRequests handles listing requests with optional filters
// @Summary List requests
// @Description Get a list of requests with optional filters
// @Tags requests - admin
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param page_size query int true "Number of items per page"
// @Param status query string false "Filter by status"
// @Param kind query string false "Filter by kind"
// @Param user_id query int false "Filter by user ID"
// @Success 200 {object} ListRequestResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/admin/requests [get]
func (h *AdminHandler) GetRequests(c echo.Context) error {
	var req GetRequestsRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid query parameters")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	requests, total, err := h.requestService.GetRequests(c.Request().Context(), &contract.RequestFilter{
		Page:     req.Page,
		PageSize: req.PageSize,
		UserID:   req.UserID,
		Kind:     req.Kind,
		Status:   req.Status,
	})
	if err != nil {
		return dto.InternalServerErrorJSON(c, "failed to get requests")
	}

	return dto.SuccessJSON(c, RequestListData{
		Items: requests,
		Total: total,
	})
}

// UpdateRequestStatus handles updating a request's status
// @Summary Update request status
// @Description Update the status of a request
// @Tags requests - admin
// @Accept json
// @Produce json
// @Param id path int true "Request ID"
// @Param request body UpdateRequestStatusRequest true "Status update details"
// @Success 200 {object} UpdateRequestStatusResponse
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/admin/requests/{id}/status [put]
func (h *AdminHandler) UpdateRequestStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return dto.BadRequestJSON(c, "invalid request ID")
	}

	var req UpdateRequestStatusRequest
	if err := c.Bind(&req); err != nil {
		return dto.BadRequestJSON(c, "invalid request body")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return dto.BadRequestJSON(c, err.Error())
	}

	request, err := h.requestService.UpdateRequest(c.Request().Context(), id, &contract.UpdateRequestInput{
		Status:      req.Status,
		Description: req.Description,
	})
	if err != nil {
		return dto.InternalServerErrorJSON(c, "failed to update request status")
	}

	return dto.SuccessJSON(c, request)
}
