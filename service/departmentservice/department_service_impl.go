package departmentservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/department"
)

type service struct {
	departmentRepo department.Repository
}

// New creates a new DepartmentService instance
func New(departmentRepo department.Repository) Service {
	return &service{
		departmentRepo: departmentRepo,
	}
}

// Create creates a new department
func (s *service) Create(ctx context.Context, input *contract.DepartmentInput) (*contract.Department, error) {
	return s.departmentRepo.Create(ctx, input)
}

// Update updates an existing department
func (s *service) Update(ctx context.Context, id int, input *contract.DepartmentInput) (*contract.Department, error) {
	// Check if department exists
	_, err := s.departmentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, contract.ErrDepartmentNotFound
	}

	return s.departmentRepo.Update(ctx, id, input)
}

// GetByID retrieves a department by its ID
func (s *service) GetByID(ctx context.Context, id int) (*contract.Department, error) {
	department, err := s.departmentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, contract.ErrDepartmentNotFound
	}
	return department, nil
}

// List retrieves a paginated list of departments
func (s *service) List(ctx context.Context, page, limit int) ([]*contract.Department, int, error) {
	return s.departmentRepo.List(ctx, page, limit)
}

// Delete deletes a department by its ID
func (s *service) Delete(ctx context.Context, id int) error {
	// Check if department exists
	_, err := s.departmentRepo.GetByID(ctx, id)
	if err != nil {
		return contract.ErrDepartmentNotFound
	}

	return s.departmentRepo.Delete(ctx, id)
}
