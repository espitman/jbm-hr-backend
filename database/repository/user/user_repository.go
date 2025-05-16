package user

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for user data access
type Repository interface {
	// GetAll retrieves all users
	GetAll(ctx context.Context, filters *contract.UserFilters) ([]*contract.User, error)

	// GetByID retrieves a user by their ID
	GetByID(ctx context.Context, id int) (*contract.User, error)

	// GetByEmail retrieves a user by their email
	GetByEmail(ctx context.Context, email string) (*contract.User, error)

	// Create creates a new user
	Create(ctx context.Context, req *contract.CreateUserInput) (*contract.User, error)

	// Update updates an existing user
	Update(ctx context.Context, id int, req *contract.UpdateUserInput) (*contract.User, error)

	// UpdateAvatar updates only the avatar of a user
	UpdateAvatar(ctx context.Context, id int, avatar string) (*contract.User, error)

	// UpdatePassword updates a user's password
	UpdatePassword(ctx context.Context, id int, req *contract.UpdatePasswordInput) error

	// UpdateBirthdate updates a user's birthdate
	UpdateBirthdate(ctx context.Context, id int, birthdate string) (*contract.User, error)

	// UpdateCooperationStartDate updates a user's cooperation start date
	UpdateCooperationStartDate(ctx context.Context, id int, startDate string) (*contract.User, error)

	// UpdateConfirmed updates a user's confirmed status
	UpdateConfirmed(ctx context.Context, id int) (*contract.User, error)

	// Delete deletes a user by their ID
	Delete(ctx context.Context, id int) error

	// SearchUsers searches users by term (full name, email, or phone)
	SearchUsers(ctx context.Context, term string) ([]*contract.User, error)
}
