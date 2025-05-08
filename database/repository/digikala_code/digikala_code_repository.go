package digikala_code

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for Digikala code operations
type Repository interface {
	// Create creates a new Digikala code
	Create(ctx context.Context, req *contract.CreateDigikalaCodeInput) (*contract.DigikalaCode, error)

	// GetAll retrieves all Digikala codes with pagination and filters
	GetAll(ctx context.Context, page, pageSize int, used *bool, userID *int) ([]*contract.DigikalaCode, int, error)

	// GetByID retrieves a Digikala code by its ID
	GetByID(ctx context.Context, id int) (*contract.DigikalaCode, error)

	// GetByCode retrieves a Digikala code by its code
	GetByCode(ctx context.Context, code string) (*contract.DigikalaCode, error)

	// Assign assigns a Digikala code to a user
	Assign(ctx context.Context, code string, req *contract.AssignDigikalaCodeInput) (*contract.DigikalaCode, error)

	// Delete deletes a Digikala code by its ID
	Delete(ctx context.Context, id int) error
}
