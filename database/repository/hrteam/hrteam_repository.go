package hrteam

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for HR team data operations
type Repository interface {
	// GetAll retrieves all HR team members
	GetAll(ctx context.Context) ([]*contract.HRTeam, error)

	// GetByID retrieves an HR team member by their ID
	GetByID(ctx context.Context, id int) (*contract.HRTeam, error)

	// Create creates a new HR team member
	Create(ctx context.Context, req *contract.HRTeamInput) (*contract.HRTeam, error)

	// Update updates an existing HR team member
	Update(ctx context.Context, id int, req *contract.HRTeamInput) (*contract.HRTeam, error)

	// Delete deletes an HR team member by their ID
	Delete(ctx context.Context, id int) error
}
