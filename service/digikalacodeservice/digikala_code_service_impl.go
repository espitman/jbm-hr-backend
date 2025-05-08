package digikalacodeservice

import (
	"context"
	"fmt"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/digikala_code"
)

// service implements the Service interface
type service struct {
	repo digikala_code.Repository
}

// New creates a new Digikala code service
func New(repo digikala_code.Repository) Service {
	return &service{repo: repo}
}

// Create creates a new Digikala code
func (s *service) Create(ctx context.Context, req *contract.CreateDigikalaCodeInput) (*contract.DigikalaCode, error) {
	if req.Code == "" {
		return nil, fmt.Errorf("code is required")
	}
	return s.repo.Create(ctx, req)
}

// GetAll retrieves all Digikala codes with pagination
func (s *service) GetAll(ctx context.Context, page, pageSize int) ([]*contract.DigikalaCode, int, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // Default page size
	}
	if pageSize > 100 {
		pageSize = 100 // Maximum page size
	}

	return s.repo.GetAll(ctx, page, pageSize)
}

// GetByID retrieves a Digikala code by its ID
func (s *service) GetByID(ctx context.Context, id int) (*contract.DigikalaCode, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id")
	}
	return s.repo.GetByID(ctx, id)
}

// GetByCode retrieves a Digikala code by its code
func (s *service) GetByCode(ctx context.Context, code string) (*contract.DigikalaCode, error) {
	if code == "" {
		return nil, fmt.Errorf("code is required")
	}
	return s.repo.GetByCode(ctx, code)
}

// Assign assigns a Digikala code to a user
func (s *service) Assign(ctx context.Context, code string, req *contract.AssignDigikalaCodeInput) (*contract.DigikalaCode, error) {
	if req.ID <= 0 {
		return nil, fmt.Errorf("invalid id")
	}
	if req.UserID <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	return s.repo.Assign(ctx, code, req)
}

// Delete deletes a Digikala code by its ID
func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
