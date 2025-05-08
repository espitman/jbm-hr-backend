package digikala_code

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/digikala_code"
)

// Service defines the interface for Digikala code operations
type Service interface {
	// Create creates a new Digikala code
	Create(ctx context.Context, req *contract.CreateDigikalaCodeInput) (*contract.DigikalaCode, error)

	// GetAll retrieves all Digikala codes
	GetAll(ctx context.Context) ([]*contract.DigikalaCode, error)

	// GetByID retrieves a Digikala code by its ID
	GetByID(ctx context.Context, id int) (*contract.DigikalaCode, error)

	// GetByCode retrieves a Digikala code by its code
	GetByCode(ctx context.Context, code string) (*contract.DigikalaCode, error)

	// Assign assigns a Digikala code to a user
	Assign(ctx context.Context, code string, req *contract.AssignDigikalaCodeInput) (*contract.DigikalaCode, error)
}

// service implements the Service interface
type service struct {
	repo digikala_code.Repository
}

// NewService creates a new Digikala code service
func NewService(repo digikala_code.Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new Digikala code
func (s *service) Create(ctx context.Context, req *contract.CreateDigikalaCodeInput) (*contract.DigikalaCode, error) {
	return s.repo.Create(ctx, req)
}

// GetAll retrieves all Digikala codes
func (s *service) GetAll(ctx context.Context) ([]*contract.DigikalaCode, error) {
	return s.repo.GetAll(ctx)
}

// GetByID retrieves a Digikala code by its ID
func (s *service) GetByID(ctx context.Context, id int) (*contract.DigikalaCode, error) {
	return s.repo.GetByID(ctx, id)
}

// GetByCode retrieves a Digikala code by its code
func (s *service) GetByCode(ctx context.Context, code string) (*contract.DigikalaCode, error) {
	return s.repo.GetByCode(ctx, code)
}

// Assign assigns a Digikala code to a user
func (s *service) Assign(ctx context.Context, code string, req *contract.AssignDigikalaCodeInput) (*contract.DigikalaCode, error) {
	return s.repo.Assign(ctx, code, req)
}
