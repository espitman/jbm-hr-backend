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

	// AdminLogin authenticates an admin user and returns a JWT token
	AdminLogin(ctx context.Context, email, password string) (string, *contract.User, error)

	// UpdateUser updates a user's information
	UpdateUser(ctx context.Context, id int, input *contract.UpdateUserInput) (*contract.User, error)

	// UpdateUserPassword updates a user's password by admin
	UpdateUserPassword(ctx context.Context, id int, input *contract.UpdatePasswordInput) error

	// UpdateAvatar updates only the avatar of a user
	UpdateAvatar(ctx context.Context, id int, avatar string) (*contract.User, error)

	// UpdateBirthdate updates a user's birthdate
	UpdateBirthdate(ctx context.Context, id int, birthdate string) (*contract.User, error)

	// UpdateCooperationStartDate updates a user's cooperation start date
	UpdateCooperationStartDate(ctx context.Context, id int, startDate string) (*contract.User, error)

	// SearchUsers searches users by term (full name, email, or phone)
	SearchUsers(ctx context.Context, term string) ([]*contract.User, error)
}
