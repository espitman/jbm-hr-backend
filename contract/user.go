package contract

// User represents a user in the system
type User struct {
	ID                   int     `json:"id"`
	Email                string  `json:"email"`
	Phone                string  `json:"phone"`
	FirstName            string  `json:"first_name"`
	LastName             string  `json:"last_name"`
	FullName             string  `json:"full_name"`
	Role                 string  `json:"role"`
	Avatar               string  `json:"avatar,omitempty"`
	Password             string  `json:"-"`
	DepartmentID         *int    `json:"department_id,omitempty"`
	DepartmentTitle      *string `json:"department_title,omitempty"`
	DepartmentIcon       *string `json:"department_icon,omitempty"`
	DepartmentShortName  *string `json:"department_short_name,omitempty"`
	Birthdate            *string `json:"birthdate,omitempty"`
	CooperationStartDate *string `json:"cooperation_start_date,omitempty"`
	PersonnelNumber      string  `json:"personnel_number,omitempty"`
	NationalCode         string  `json:"national_code,omitempty"`
	Confirmed            bool    `json:"confirmed"`
	Active               bool    `json:"active"`
	Age                  *int    `json:"age,omitempty"`
	CooperationDuration  *int    `json:"cooperation_duration,omitempty"`
}

// CreateUserInput represents the input for creating a new user
type CreateUserInput struct {
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

// UpdateUserInput represents the input for updating a user
type UpdateUserInput struct {
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

// RegisterUserInput represents the input for registering a new user
type RegisterUserInput struct {
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

// UpdatePasswordInput represents the input for updating a user's password
type UpdatePasswordInput struct {
	Password string `json:"password" validate:"required"`
}

// UserFilters represents the filters for listing users
type UserFilters struct {
	FullName        *string `json:"full_name,omitempty"`
	Role            *string `json:"role,omitempty"`
	PersonnelNumber *string `json:"personnel_number,omitempty"`
	NationalCode    *string `json:"national_code,omitempty"`
	Phone           *string `json:"phone,omitempty"`
	DepartmentID    *int    `json:"department_id,omitempty"`
}
