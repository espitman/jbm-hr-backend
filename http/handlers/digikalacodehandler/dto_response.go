package digikalacodehandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// CreateDigikalaCodeResponse represents the response for creating a Digikala code
type CreateDigikalaCodeResponse struct {
	dto.Response
	Data contract.DigikalaCode `json:"data"`
}

// GetDigikalaCodeResponse represents the response for getting a Digikala code
type GetDigikalaCodeResponse struct {
	dto.Response
	Data contract.DigikalaCode `json:"data"`
}

// DigikalaCodeListData represents the data structure for listing Digikala codes
type DigikalaCodeListData struct {
	Items []contract.DigikalaCode `json:"items"`
}

// ListDigikalaCodeResponse represents the response for listing Digikala codes
type ListDigikalaCodeResponse struct {
	dto.Response
	Data DigikalaCodeListData `json:"data"`
}

// AssignDigikalaCodeResponse represents the response for assigning a Digikala code
type AssignDigikalaCodeResponse struct {
	dto.Response
	Data contract.DigikalaCode `json:"data"`
}
