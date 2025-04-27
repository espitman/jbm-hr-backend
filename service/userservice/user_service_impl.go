package userservice

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/otp"
	"github.com/espitman/jbm-hr-backend/database/repository/user"
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

// generateOTP generates a random 6-digit numeric OTP code
func generateOTP() (string, error) {
	var n uint32
	err := binary.Read(rand.Reader, binary.BigEndian, &n)
	if err != nil {
		return "", err
	}
	// Ensure the number is between 100000 and 999999
	code := (n % 900000) + 100000
	return fmt.Sprintf("%06d", code), nil
}

// RequestOTP generates and sends a new OTP for a user
func (s *service) RequestOTP(ctx context.Context, email string) (*contract.OTP, error) {
	// Check for existing active OTPs
	activeOTPs, err := s.otpRepo.GetActiveByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if len(activeOTPs) > 0 {
		return nil, ErrActiveOTPExists
	}

	// Generate new OTP
	code, err := generateOTP()
	if err != nil {
		return nil, err
	}

	// Set expiration time (5 minutes from now)
	expiresAt := time.Now().Add(5 * time.Minute)

	// Create OTP
	otp, err := s.otpRepo.Create(ctx, email, code, expiresAt)
	if err != nil {
		return nil, err
	}

	// TODO: Send OTP via email/SMS
	// This would be implemented based on your notification service

	return otp, nil
}

// VerifyOTP verifies an OTP code for a user
func (s *service) VerifyOTP(ctx context.Context, email string, code string) (bool, error) {
	// Get OTP by code
	otp, err := s.otpRepo.GetByCode(ctx, code)
	if err != nil {
		return false, ErrOTPNotFound
	}

	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	// Verify user ID matches
	if otp.UserID != user.ID {
		return false, ErrOTPInvalid
	}

	// Check if OTP is expired
	if otp.ExpiresAt.Before(time.Now()) {
		return false, ErrOTPExpired
	}

	// Check if OTP is already used
	if otp.Used {
		return false, ErrOTPAlreadyUsed
	}

	// Mark OTP as used
	err = s.otpRepo.MarkAsUsed(ctx, otp.ID)
	if err != nil {
		return false, err
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
		return nil, err
	}

	return user, nil
}
