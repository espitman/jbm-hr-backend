package hrteamservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for HR team-related operations
type Service interface {
	// GetAllHRTeam returns all HR team members
	GetAllHRTeam(ctx context.Context) ([]*contract.HRTeam, error)

	// CreateHRTeam creates a new HR team member
	CreateHRTeam(ctx context.Context, input *contract.HRTeamInput) (*contract.HRTeam, error)

	// GetHRTeamByID returns an HR team member by their ID
	GetHRTeamByID(ctx context.Context, id int) (*contract.HRTeam, error)

	// UpdateHRTeam updates an existing HR team member
	UpdateHRTeam(ctx context.Context, id int, input *contract.HRTeamInput) (*contract.HRTeam, error)

	// DeleteHRTeam deletes an HR team member by their ID
	DeleteHRTeam(ctx context.Context, id int) error
}
