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

// VerifyOTP verifies an OTP code for a user
func (s *service) VerifyOTP(ctx context.Context, email string, code string) (bool, error) {
	// Get OTP by code
	otp, err := s.otpRepo.GetByCode(ctx, code)
	if err != nil {
		return false, contract.ErrOTPNotFound
	}

	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return false, contract.ErrUserNotFound
	}

	// Verify user ID matches
	if otp.UserID != user.ID {
		return false, contract.ErrOTPInvalid
	}

	// Check if OTP is expired
	if otp.ExpiresAt.Before(time.Now()) {
		return false, contract.ErrOTPExpired
	}

	// Check if OTP is already used
	if otp.Used {
		return false, contract.ErrOTPAlreadyUsed
	}

	// Mark OTP as used
	err = s.otpRepo.MarkAsUsed(ctx, otp.ID)
	if err != nil {
		return false, contract.ErrDatabaseQuery
	}

	return true, nil
}

// RegisterUser registers a new user in the system
func (s *service) RegisterUser(ctx context.Context, input *contract.RegisterUserInput) (*contract.User, error) {
	// Create a CreateUserInput from the RegisterUserInput
	createInput := &contract.CreateUserInput{
		Email:  input.Email,
		Phone:  input.Phone,
		Role:   input.Role,
		Avatar: input.Avatar,
	}

	// Create the user using the repository
	user, err := s.userRepo.Create(ctx, createInput)
	if err != nil {
		return nil, contract.ErrDatabaseQuery
	}

	return user, nil
}

// Helper functions
func generateOTP() (string, error) {
	// Generate a 6-digit OTP
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	return otp, nil
}
