package userservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for user-related operations
type Service interface {
	// RequestOTP generates and sends a new OTP for a user
	RequestOTP(ctx context.Context, email string) (*contract.OTP, error)

	// VerifyOTP verifies an OTP code for a user
	VerifyOTP(ctx context.Context, email string, code string) (bool, error)
}
