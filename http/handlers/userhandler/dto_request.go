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
	PersonnelNumber      string  `json:"personnel_number,omitempty"`
	NationalCode         string  `json:"national_code,omitempty"`
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

// UpdateUserBirthdateRequest represents the request body for updating a user's birthdate
type UpdateUserBirthdateRequest struct {
	Birthdate string `json:"birthdate" validate:"required,datetime=2006-01-02"`
}

// UpdateUserCooperationStartDateRequest represents the request body for updating a user's cooperation start date
type UpdateUserCooperationStartDateRequest struct {
	CooperationStartDate string `json:"cooperation_start_date" validate:"required,datetime=2006-01-02"`
}

// ListUsersRequest represents the request structure for listing users with filters
type ListUsersRequest struct {
	Page            int     `query:"page" validate:"required,min=1"`
	Limit           int     `query:"limit" validate:"required,min=1,max=100"`
	FullName        *string `query:"full_name"`
	Role            *string `query:"role" validate:"omitempty,oneof=admin employee"`
	PersonnelNumber *string `query:"personnel_number"`
	NationalCode    *string `query:"national_code"`
	Phone           *string `query:"phone"`
	DepartmentID    *int    `query:"department_id"`
}
