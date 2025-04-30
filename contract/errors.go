package contract

import "errors"

// Common errors
var (
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("resource not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
)

// User related errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidEmail      = errors.New("invalid email format")
	ErrInvalidPhone      = errors.New("invalid phone format")
	ErrInvalidRole       = errors.New("invalid role")
)

// OTP related errors
var (
	ErrOTPNotFound         = errors.New("OTP not found")
	ErrOTPInvalid          = errors.New("invalid OTP")
	ErrOTPExpired          = errors.New("OTP expired")
	ErrOTPAlreadyUsed      = errors.New("OTP already used")
	ErrActiveOTPExists     = errors.New("active OTP already exists")
	ErrOTPGenerationFailed = errors.New("failed to generate OTP")
)

// Authentication related errors
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrTokenMissing       = errors.New("token missing")
)

// Database related errors
var (
	ErrDatabaseConnection  = errors.New("database connection error")
	ErrDatabaseQuery       = errors.New("database query error")
	ErrDatabaseTransaction = errors.New("database transaction error")
)

// Validation related errors
var (
	ErrValidationFailed = errors.New("validation failed")
	ErrRequiredField    = errors.New("required field missing")
	ErrInvalidFormat    = errors.New("invalid format")
)

// HR Team related errors
var (
	ErrHRTeamNotFound = errors.New("hr team member not found")
)

// Resume related errors
var (
	ErrResumeNotFound = errors.New("resume not found")
)

// Error definitions for department-related operations
var (
	ErrDepartmentNotFound = errors.New("department not found")
	ErrDepartmentExists   = errors.New("department already exists")
)
