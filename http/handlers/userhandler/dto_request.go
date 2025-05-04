package userhandler

// RequestOTPRequest represents the request structure for requesting an OTP
type RequestOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// VerifyOTPRequest represents the request structure for verifying an OTP
type VerifyOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,len=6"`
}

// RegisterUserRequest represents the request structure for registering a user
type RegisterUserRequest struct {
	Email                string  `json:"email" validate:"required,email"`
	Phone                string  `json:"phone" validate:"required"`
	FirstName            string  `json:"first_name" validate:"required"`
	LastName             string  `json:"last_name" validate:"required"`
	Role                 string  `json:"role" validate:"required,oneof=admin employee"`
	Avatar               string  `json:"avatar,omitempty"`
	DepartmentID         *int    `json:"department_id,omitempty"`
	Birthdate            *string `json:"birthdate,omitempty"`
	CooperationStartDate *string `json:"cooperation_start_date,omitempty"`
}

// AdminLoginRequest represents the request body for admin login
type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UpdateUserAvatarRequest represents the request body for updating a user's avatar
type UpdateUserAvatarRequest struct {
	Avatar string `json:"avatar" validate:"required,url"`
}
