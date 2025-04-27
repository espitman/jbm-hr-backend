package userhandler

import (
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// RequestOTPResponse represents the response structure for OTP request
type RequestOTPResponse struct {
	dto.Response
	Data struct{} `json:"data,omitempty"`
}

// VerifyOTPResponse represents the response structure for OTP verification
type VerifyOTPResponse struct {
	dto.Response
	Data struct {
		Token string `json:"token,omitempty"`
	} `json:"data,omitempty"`
}
