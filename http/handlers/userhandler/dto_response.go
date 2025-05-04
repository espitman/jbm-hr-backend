package userhandler

import (
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// RequestOTPResponse represents the response structure for OTP request
type RequestOTPResponse struct {
	dto.Response
}

// DepartmentDTO represents a department in user responses
type DepartmentDTO struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	ShortName string `json:"short_name"`
}

// VerifyOTPUserData represents the user data structure for OTP verification response
type VerifyOTPUserData struct {
	ID                   int            `json:"id"`
	Email                string         `json:"email"`
	Phone                string         `json:"phone"`
	FirstName            string         `json:"first_name"`
	LastName             string         `json:"last_name"`
	Role                 string         `json:"role"`
	Avatar               string         `json:"avatar,omitempty"`
	Department           *DepartmentDTO `json:"department,omitempty"`
	Birthdate            *string        `json:"birthdate,omitempty"`
	CooperationStartDate *string        `json:"cooperation_start_date,omitempty"`
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
	ID                   int            `json:"id"`
	Email                string         `json:"email"`
	Phone                string         `json:"phone"`
	FirstName            string         `json:"first_name"`
	LastName             string         `json:"last_name"`
	Role                 string         `json:"role"`
	Avatar               string         `json:"avatar,omitempty"`
	Department           *DepartmentDTO `json:"department,omitempty"`
	Birthdate            *string        `json:"birthdate,omitempty"`
	CooperationStartDate *string        `json:"cooperation_start_date,omitempty"`
}

// RegisterUserResponse represents the response structure for user registration
type RegisterUserResponse struct {
	dto.Response
	Data RegisterUserData `json:"data,omitempty"`
}

// GetMeData represents the data structure for the /me endpoint response
type GetMeData struct {
	ID                   int            `json:"id"`
	Email                string         `json:"email"`
	Phone                string         `json:"phone"`
	FirstName            string         `json:"first_name"`
	LastName             string         `json:"last_name"`
	Role                 string         `json:"role"`
	Avatar               string         `json:"avatar,omitempty"`
	Department           *DepartmentDTO `json:"department,omitempty"`
	Birthdate            *string        `json:"birthdate,omitempty"`
	CooperationStartDate *string        `json:"cooperation_start_date,omitempty"`
}

// GetMeResponse represents the response structure for the /me endpoint
type GetMeResponse struct {
	dto.Response
	Data GetMeData `json:"data,omitempty"`
}

// UserData represents the user data structure for list users response
type UserData struct {
	ID                   int            `json:"id"`
	Email                string         `json:"email"`
	Phone                string         `json:"phone"`
	FirstName            string         `json:"first_name"`
	LastName             string         `json:"last_name"`
	Role                 string         `json:"role"`
	Avatar               string         `json:"avatar,omitempty"`
	Department           *DepartmentDTO `json:"department,omitempty"`
	Birthdate            *string        `json:"birthdate,omitempty"`
	CooperationStartDate *string        `json:"cooperation_start_date,omitempty"`
}

// ListUsersData represents the data structure for list users response
type ListUsersData struct {
	Users []UserData `json:"users"`
	Total int64      `json:"total"`
}

// ListUsersResponse represents the response structure for list users
type ListUsersResponse struct {
	dto.Response
	Data ListUsersData `json:"data,omitempty"`
}

// AdminLoginResponse represents the response structure for admin login
type AdminLoginResponse struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}

// UpdateUserResponse represents the response for updating a user
type UpdateUserResponse struct {
	dto.Response
	Data UserData `json:"data"`
}
