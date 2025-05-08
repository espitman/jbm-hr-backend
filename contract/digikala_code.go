package contract

// DigikalaCode represents a Digikala code in the system
type DigikalaCode struct {
	ID             int     `json:"id"`
	Code           string  `json:"code,omitempty"`
	Used           bool    `json:"used"`
	CreatedAt      string  `json:"created_at"`
	AssignToUserID *int    `json:"assign_to_user_id,omitempty"`
	AssignAt       *string `json:"assign_at,omitempty"`
	AssignedToUser *User   `json:"assigned_to_user,omitempty"`
}

// CreateDigikalaCodeInput represents the input for creating a new Digikala code
type CreateDigikalaCodeInput struct {
	Code string `json:"code" validate:"required"`
}

// AssignDigikalaCodeInput represents the input for assigning a Digikala code to a user
type AssignDigikalaCodeInput struct {
	ID     int `json:"id" validate:"required"`
	UserID int `json:"user_id" validate:"required"`
}

// UseDigikalaCodeInput represents the input for marking a Digikala code as used
type UseDigikalaCodeInput struct {
	Code string `json:"code" validate:"required"`
}
