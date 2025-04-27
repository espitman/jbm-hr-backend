package contract

// User represents a user in the system
type User struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Avatar string `json:"avatar,omitempty"`
}

// CreateUserInput represents the input for creating a new user
type CreateUserInput struct {
	Email  string `json:"email" validate:"required,email"`
	Phone  string `json:"phone" validate:"required"`
	Role   string `json:"role" validate:"required,oneof=admin employee"`
	Avatar string `json:"avatar,omitempty"`
}

// UpdateUserInput represents the input for updating a user
type UpdateUserInput struct {
	Email  string `json:"email" validate:"required,email"`
	Phone  string `json:"phone" validate:"required"`
	Role   string `json:"role" validate:"required,oneof=admin employee"`
	Avatar string `json:"avatar,omitempty"`
}
