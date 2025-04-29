package hrteamhandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// HRTeamResponse represents the response structure for HR team operations
type HRTeamResponse struct {
	dto.Response
	Data *contract.HRTeam `json:"data,omitempty"`
}

// HRTeamListResponse represents the response structure for multiple HR team members
type HRTeamListResponse struct {
	dto.Response
	Data []*contract.HRTeam `json:"data,omitempty"`
}

// CreateHRTeamResponse represents the response structure for creating an HR team member
type CreateHRTeamResponse struct {
	dto.Response
	Data *contract.HRTeam `json:"data,omitempty"`
}

// UpdateHRTeamResponse represents the response structure for updating an HR team member
type UpdateHRTeamResponse struct {
	dto.Response
	Data *contract.HRTeam `json:"data,omitempty"`
}

// GetHRTeamResponse represents the response structure for getting an HR team member
type GetHRTeamResponse struct {
	dto.Response
	Data *contract.HRTeam `json:"data,omitempty"`
}
