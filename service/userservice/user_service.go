package userservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/user"
)

// UserService handles user-related business logic
type UserService struct {
	userRepository user.Repository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepository user.Repository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

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
	ListUsers(ctx context.Context, page, limit int, filters *contract.UserFilters) ([]*contract.User, int64, error)

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

	// UpdateConfirmed updates a user's confirmed status
	UpdateConfirmed(ctx context.Context, id int) (*contract.User, error)

	// UpdateActive updates a user's active status
	UpdateActive(ctx context.Context, id int, active bool) (*contract.User, error)

	// SearchUsers searches users by term (full name, email, or phone)
	SearchUsers(ctx context.Context, term string) ([]*contract.User, error)

	// GetUsersWithTodayBirthdate retrieves all users whose birthdate is today
	GetUsersWithTodayBirthdate(ctx context.Context) ([]*contract.User, error)

	// GetUsersWithTodayCooperationStartDate retrieves all users whose cooperation start date is today
	GetUsersWithTodayCooperationStartDate(ctx context.Context) ([]*contract.User, error)

	// GetUsersWithBirthdateInJalaliMonth retrieves all users whose birthdate is in the current Jalali month
	GetUsersWithBirthdateInJalaliMonth(ctx context.Context) ([]*contract.User, error)

	// GetUsersWithCooperationStartDateInJalaliMonth retrieves all users whose cooperation start date is in the current Jalali month
	GetUsersWithCooperationStartDateInJalaliMonth(ctx context.Context) ([]*contract.User, error)
}

// ListUsers retrieves a list of users with pagination
func (s *UserService) ListUsers(ctx context.Context, page, limit int, filters *contract.UserFilters) ([]*contract.User, int64, error) {
	// Get all users with filters
	users, err := s.userRepository.GetAll(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	// Calculate total count
	total := int64(len(users))

	// Apply pagination
	start := (page - 1) * limit
	end := start + limit
	if int64(start) >= total {
		return []*contract.User{}, total, nil
	}
	if int64(end) > total {
		end = int(total)
	}

	return users[start:end], total, nil
}
