package userservice

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/otp"
	"github.com/espitman/jbm-hr-backend/database/repository/user"
	"github.com/espitman/jbm-hr-backend/utils"
)

var (
	ErrOTPNotFound     = errors.New("OTP not found")
	ErrOTPExpired      = errors.New("OTP has expired")
	ErrOTPAlreadyUsed  = errors.New("OTP has already been used")
	ErrOTPInvalid      = errors.New("invalid OTP")
	ErrActiveOTPExists = errors.New("active OTP already exists")
)

type service struct {
	userRepo user.Repository
	otpRepo  otp.Repository
}

// New creates a new UserService instance
func New(userRepo user.Repository, otpRepo otp.Repository) Service {
	return &service{
		userRepo: userRepo,
		otpRepo:  otpRepo,
	}
}

// RequestOTP generates and sends a new OTP for a user
func (s *service) RequestOTP(ctx context.Context, email string) (*contract.OTP, error) {
	// Check if user exists
	_, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, contract.ErrUserNotFound
	}

	// Check for existing active OTPs
	activeOTPs, err := s.otpRepo.GetActiveByEmail(ctx, email)
	if err != nil {
		return nil, contract.ErrDatabaseQuery
	}
	if len(activeOTPs) > 0 {
		return nil, contract.ErrActiveOTPExists
	}

	// Generate new OTP
	code, err := generateOTP()
	if err != nil {
		return nil, contract.ErrOTPGenerationFailed
	}

	// Set expiration time (5 minutes from now)
	expiresAt := time.Now().Add(5 * time.Minute)

	// Create OTP
	otp, err := s.otpRepo.Create(ctx, email, code, expiresAt)
	if err != nil {
		return nil, contract.ErrDatabaseQuery
	}

	// Send OTP via email
	if err := utils.SendOTPEmail(email, code); err != nil {
		// Log the error but don't return it to the user
		fmt.Printf("Failed to send OTP email: %v\n", err)
	}

	return otp, nil
}

// VerifyOTP verifies an OTP code for a user and returns a JWT token and user data if valid
func (s *service) VerifyOTP(ctx context.Context, email string, code string) (string, *contract.User, error) {
	// Get OTP by code
	otp, err := s.otpRepo.GetByCode(ctx, code)
	if err != nil {
		return "", nil, contract.ErrOTPNotFound
	}

	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, contract.ErrUserNotFound
	}

	// Verify user ID matches
	if otp.UserID != user.ID {
		return "", nil, contract.ErrOTPInvalid
	}

	// Check if OTP is expired
	if otp.ExpiresAt.Before(time.Now()) {
		return "", nil, contract.ErrOTPExpired
	}

	// Check if OTP is already used
	if otp.Used {
		return "", nil, contract.ErrOTPAlreadyUsed
	}

	// Mark OTP as used
	err = s.otpRepo.MarkAsUsed(ctx, otp.ID)
	if err != nil {
		return "", nil, contract.ErrDatabaseQuery
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, user, nil
}

// RegisterUser registers a new user in the system
func (s *service) RegisterUser(ctx context.Context, input *contract.RegisterUserInput) (*contract.User, error) {
	// Create a CreateUserInput from the RegisterUserInput
	createInput := &contract.CreateUserInput{
		Email:     input.Email,
		Phone:     input.Phone,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Role:      input.Role,
		Avatar:    input.Avatar,
	}

	// Create the user using the repository
	user, err := s.userRepo.Create(ctx, createInput)
	if err != nil {
		return nil, contract.ErrDatabaseQuery
	}

	return user, nil
}

// GetUserByID retrieves a user by their ID
func (s *service) GetUserByID(ctx context.Context, id int) (*contract.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// ListUsers retrieves a paginated list of users
func (s *service) ListUsers(ctx context.Context, page, limit int) ([]*contract.User, int64, error) {
	// Get all users
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Calculate total count
	total := int64(len(users))

	// Calculate pagination
	start := (page - 1) * limit
	end := start + limit

	// Handle out of bounds
	if start >= len(users) {
		return []*contract.User{}, total, nil
	}
	if end > len(users) {
		end = len(users)
	}

	// Return paginated results
	return users[start:end], total, nil
}

// UpdatePassword updates a user's password
func (s *service) UpdatePassword(ctx context.Context, id int, input *contract.UpdatePasswordInput) error {
	// Check if user exists and get their role
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return contract.ErrUserNotFound
	}

	// Check if user has admin role
	if user.Role != "admin" {
		return errors.New("only admin users can set passwords")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create a new input with the hashed password
	hashedInput := &contract.UpdatePasswordInput{
		Password: hashedPassword,
	}

	// Update the password
	return s.userRepo.UpdatePassword(ctx, id, hashedInput)
}

// UpdateUser updates a user's information
func (s *service) UpdateUser(ctx context.Context, id int, input *contract.UpdateUserInput) (*contract.User, error) {
	// Check if user exists
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, contract.ErrUserNotFound
	}

	// Update user
	updatedUser, err := s.userRepo.Update(ctx, id, input)
	if err != nil {
		return nil, contract.ErrDatabaseQuery
	}

	return updatedUser, nil
}

// Helper functions
func generateOTP() (string, error) {
	// Generate a 6-digit OTP
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	return otp, nil
}

// AdminLogin authenticates an admin user and returns a JWT token
func (s *service) AdminLogin(ctx context.Context, email, password string) (string, *contract.User, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, contract.ErrUserNotFound
	}

	// Check if user is an admin
	if user.Role != "admin" {
		return "", nil, errors.New("only admin users can login with password")
	}

	// Check if user has a password set
	if user.Password == "" {
		return "", nil, errors.New("password not set for this user")
	}

	// Compare passwords
	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", nil, errors.New("invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, user, nil
}

// UpdateUserPassword updates a user's password by admin
func (s *service) UpdateUserPassword(ctx context.Context, id int, input *contract.UpdatePasswordInput) error {
	// Check if user exists and get their role
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return contract.ErrUserNotFound
	}

	// Check if user has admin role
	if user.Role != "admin" {
		return errors.New("only admin users can set passwords")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create a new input with the hashed password
	hashedInput := &contract.UpdatePasswordInput{
		Password: hashedPassword,
	}

	// Update the password
	return s.userRepo.UpdatePassword(ctx, id, hashedInput)
}
