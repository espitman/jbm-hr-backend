package otp

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for OTP data access
type Repository interface {
	// Create creates a new OTP for a user
	Create(ctx context.Context, userID int, code string, expiresAt time.Time) (*contract.OTP, error)

	// GetByCode retrieves an OTP by its code
	GetByCode(ctx context.Context, code string) (*contract.OTP, error)

	// GetActiveByUserID retrieves active (unused and not expired) OTPs for a user
	GetActiveByUserID(ctx context.Context, userID int) ([]*contract.OTP, error)

	// MarkAsUsed marks an OTP as used
	MarkAsUsed(ctx context.Context, id int) error

	// DeleteExpired deletes all expired OTPs
	DeleteExpired(ctx context.Context) error
}
