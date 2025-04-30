package requesthandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// CreateRequestResponse represents the response for creating a request
type CreateRequestResponse struct {
	dto.Response
	Data contract.Request `json:"data"`
}

// GetRequestResponse represents the response for getting a request
type GetRequestResponse struct {
	dto.Response
	Data contract.Request `json:"data"`
}

// RequestListData represents the data structure for listing requests
type RequestListData struct {
	Items []*contract.Request `json:"items"`
	Total int                 `json:"total"`
}

// ListRequestResponse represents the response for listing requests
type ListRequestResponse struct {
	dto.Response
	Data RequestListData `json:"data"`
}

// UpdateRequestStatusResponse represents the response for updating a request's status
type UpdateRequestStatusResponse struct {
	dto.Response
	Data contract.Request `json:"data"`
}
