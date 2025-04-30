package resumeservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for resume-related operations
type Service interface {
	// Create creates a new resume
	Create(ctx context.Context, input *contract.ResumeInput) (*contract.Resume, error)

	// Update updates an existing resume
	Update(ctx context.Context, id int, input *contract.ResumeInput) (*contract.Resume, error)

	// GetByID retrieves a resume by its ID
	GetByID(ctx context.Context, id int) (*contract.Resume, error)

	// List retrieves a paginated list of resumes
	List(ctx context.Context, page, limit int) ([]*contract.Resume, int, error)

	// Delete deletes a resume by its ID
	Delete(ctx context.Context, id int) error
}
