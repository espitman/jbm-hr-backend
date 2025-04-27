package contract

import "time"

// OTP represents an OTP entity
type OTP struct {
	ID        int        `json:"id"`
	Code      string     `json:"code"`
	ExpiresAt time.Time  `json:"expires_at"`
	Used      bool       `json:"used"`
	CreatedAt time.Time  `json:"created_at"`
	UsedAt    *time.Time `json:"used_at,omitempty"`
	UserID    int        `json:"user_id"`
}

// CreateOTPInput represents the input for creating an OTP
type CreateOTPInput struct {
	UserID    int       `json:"user_id"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
}
