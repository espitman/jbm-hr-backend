package otp

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for OTP data access
type Repository interface {
	// Create creates a new OTP for a user
	Create(ctx context.Context, email string, code string, expiresAt time.Time) (*contract.OTP, error)

	// GetByCode retrieves an OTP by its code
	GetByCode(ctx context.Context, code string) (*contract.OTP, error)

	// GetActiveByEmail retrieves active (unused and not expired) OTPs for a user
	GetActiveByEmail(ctx context.Context, email string) ([]*contract.OTP, error)

	// MarkAsUsed marks an OTP as used
	MarkAsUsed(ctx context.Context, id int) error

	// DeleteExpired deletes all expired OTPs
	DeleteExpired(ctx context.Context) error
}
