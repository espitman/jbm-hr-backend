package user

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for user data access
type Repository interface {
	// GetAll retrieves all users
	GetAll(ctx context.Context) ([]*contract.User, error)

	// GetByID retrieves a user by their ID
	GetByID(ctx context.Context, id int) (*contract.User, error)

	// GetByEmail retrieves a user by their email
	GetByEmail(ctx context.Context, email string) (*contract.User, error)

	// Create creates a new user
	Create(ctx context.Context, req *contract.CreateUserInput) (*contract.User, error)

	// Update updates an existing user
	Update(ctx context.Context, id int, req *contract.UpdateUserInput) (*contract.User, error)

	// Delete deletes a user by their ID
	Delete(ctx context.Context, id int) error
}
