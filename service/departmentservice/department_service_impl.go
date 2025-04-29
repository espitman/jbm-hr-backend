package departmentservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/department"
)

type service struct {
	repo department.Repository
}

// New creates a new instance of the department service
func New(repo department.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateDepartment(ctx context.Context, input *contract.DepartmentInput) (*contract.Department, error) {
	return s.repo.Create(ctx, input)
}

func (s *service) UpdateDepartment(ctx context.Context, id int, input *contract.DepartmentInput) (*contract.Department, error) {
	return s.repo.Update(ctx, id, input)
}

func (s *service) GetDepartmentByID(ctx context.Context, id int) (*contract.Department, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) ListDepartments(ctx context.Context, offset, limit int) ([]*contract.Department, int, error) {
	depts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	total := len(depts)
	start := offset
	end := offset + limit
	if start >= total {
		return []*contract.Department{}, total, nil
	}
	if end > total {
		end = total
	}

	return depts[start:end], total, nil
}

func (s *service) DeleteDepartment(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
