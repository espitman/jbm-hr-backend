package contract

// User represents a user in the system
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar,omitempty"`
}

// CreateUserInput represents the input for creating a new user
type CreateUserInput struct {
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=admin employee"`
	Avatar    string `json:"avatar,omitempty"`
}

// UpdateUserInput represents the input for updating a user
type UpdateUserInput struct {
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=admin employee"`
	Avatar    string `json:"avatar,omitempty"`
}

// RegisterUserInput represents the input for registering a new user
type RegisterUserInput struct {
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=admin employee"`
	Avatar    string `json:"avatar,omitempty"`
}
