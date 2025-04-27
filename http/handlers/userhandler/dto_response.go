package userhandler

import (
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// RequestOTPData represents the data structure for OTP request response
type RequestOTPData struct {
	OTP string `json:"otp"`
}

// RequestOTPResponse represents the response structure for OTP request
type RequestOTPResponse struct {
	dto.Response
	Data RequestOTPData `json:"data,omitempty"`
}

// VerifyOTPData represents the data structure for OTP verification response
type VerifyOTPData struct {
	Token string `json:"token,omitempty"`
}

// VerifyOTPResponse represents the response structure for OTP verification
type VerifyOTPResponse struct {
	dto.Response
	Data VerifyOTPData `json:"data,omitempty"`
}

// RegisterUserData represents the data structure for user registration response
type RegisterUserData struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Avatar string `json:"avatar,omitempty"`
}

// RegisterUserResponse represents the response structure for user registration
type RegisterUserResponse struct {
	dto.Response
	Data RegisterUserData `json:"data,omitempty"`
}
