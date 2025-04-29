package departmentservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for department-related operations
type Service interface {
	// Create creates a new department
	Create(ctx context.Context, input *contract.DepartmentInput) (*contract.Department, error)

	// Update updates an existing department
	Update(ctx context.Context, id int, input *contract.DepartmentInput) (*contract.Department, error)

	// GetByID retrieves a department by its ID
	GetByID(ctx context.Context, id int) (*contract.Department, error)

	// List retrieves a paginated list of departments
	List(ctx context.Context, page, limit int) ([]*contract.Department, int, error)

	// Delete deletes a department by its ID
	Delete(ctx context.Context, id int) error
}
