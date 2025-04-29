package departmentservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for department-related operations
type Service interface {
	// CreateDepartment creates a new department
	CreateDepartment(ctx context.Context, input *contract.DepartmentInput) (*contract.Department, error)

	// UpdateDepartment updates an existing department
	UpdateDepartment(ctx context.Context, id int, input *contract.DepartmentInput) (*contract.Department, error)

	// GetDepartmentByID retrieves a department by its ID
	GetDepartmentByID(ctx context.Context, id int) (*contract.Department, error)

	// ListDepartments retrieves a list of departments with optional pagination
	ListDepartments(ctx context.Context, offset, limit int) ([]*contract.Department, int, error)

	// DeleteDepartment deletes a department by its ID
	DeleteDepartment(ctx context.Context, id int) error
}
