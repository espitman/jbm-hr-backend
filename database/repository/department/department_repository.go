package department

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for department data access
type Repository interface {
	// GetAll retrieves all departments
	GetAll(ctx context.Context) ([]*contract.Department, error)

	// GetByID retrieves a department by their ID
	GetByID(ctx context.Context, id int) (*contract.Department, error)

	// Create creates a new department
	Create(ctx context.Context, req *contract.DepartmentInput) (*contract.Department, error)

	// Update updates an existing department
	Update(ctx context.Context, id int, req *contract.DepartmentInput) (*contract.Department, error)

	// Delete deletes a department by their ID
	Delete(ctx context.Context, id int) error

	// List retrieves a paginated list of departments
	List(ctx context.Context, page, limit int) ([]*contract.Department, int, error)

	// GetTotalCount returns the total number of departments
	GetTotalCount(ctx context.Context) (int, error)
}
