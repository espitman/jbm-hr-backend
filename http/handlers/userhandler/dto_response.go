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
