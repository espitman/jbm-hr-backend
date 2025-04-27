package userhandler

import (
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// RequestOTPResponse represents the response structure for OTP request
type RequestOTPResponse struct {
	dto.Response
}

// VerifyOTPUserData represents the user data structure for OTP verification response
type VerifyOTPUserData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar,omitempty"`
}

// VerifyOTPData represents the data structure for OTP verification response
type VerifyOTPData struct {
	Token string            `json:"token,omitempty"`
	User  VerifyOTPUserData `json:"user,omitempty"`
}

// VerifyOTPResponse represents the response structure for OTP verification
type VerifyOTPResponse struct {
	dto.Response
	Data VerifyOTPData `json:"data,omitempty"`
}

// RegisterUserData represents the data structure for user registration response
type RegisterUserData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar,omitempty"`
}

// RegisterUserResponse represents the response structure for user registration
type RegisterUserResponse struct {
	dto.Response
	Data RegisterUserData `json:"data,omitempty"`
}

// GetMeData represents the data structure for the /me endpoint response
type GetMeData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar,omitempty"`
}

// GetMeResponse represents the response structure for the /me endpoint
type GetMeResponse struct {
	dto.Response
	Data GetMeData `json:"data,omitempty"`
}
