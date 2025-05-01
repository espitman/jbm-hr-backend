package userservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for user-related operations
type Service interface {
	// RequestOTP generates and sends a new OTP for a user
	RequestOTP(ctx context.Context, email string) (*contract.OTP, error)

	// VerifyOTP verifies an OTP code for a user and returns a JWT token and user data if valid
	VerifyOTP(ctx context.Context, email string, code string) (string, *contract.User, error)

	// RegisterUser registers a new user in the system
	RegisterUser(ctx context.Context, input *contract.RegisterUserInput) (*contract.User, error)

	// GetUserByID retrieves a user by their ID
	GetUserByID(ctx context.Context, id int) (*contract.User, error)

	// ListUsers retrieves a paginated list of users
	ListUsers(ctx context.Context, page, limit int) ([]*contract.User, int64, error)

	// UpdatePassword updates a user's password
	UpdatePassword(ctx context.Context, id int, input *contract.UpdatePasswordInput) error
}
