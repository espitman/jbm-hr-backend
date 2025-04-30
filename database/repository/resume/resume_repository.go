package resume

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for resume data access
type Repository interface {
	// GetAll retrieves all resumes
	GetAll(ctx context.Context) ([]*contract.Resume, error)

	// GetByID retrieves a resume by their ID
	GetByID(ctx context.Context, id int) (*contract.Resume, error)

	// Create creates a new resume
	Create(ctx context.Context, req *contract.ResumeInput) (*contract.Resume, error)

	// Update updates an existing resume
	Update(ctx context.Context, id int, req *contract.ResumeInput) (*contract.Resume, error)

	// Delete deletes a resume by their ID
	Delete(ctx context.Context, id int) error

	// List retrieves a paginated list of resumes
	List(ctx context.Context, page, limit int) ([]*contract.Resume, int, error)
}
