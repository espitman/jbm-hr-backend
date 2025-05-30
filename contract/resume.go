package contract

// User represents minimal user information for resume responses
type ResumeUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

// Resume represents a resume in the system
type Resume struct {
	ID              int        `json:"id"`
	IntroducedName  string     `json:"introduced_name"`
	IntroducedPhone string     `json:"introduced_phone"`
	Position        string     `json:"position"`
	File            string     `json:"file"`
	Status          string     `json:"status"`
	UserID          int        `json:"user_id"`
	User            ResumeUser `json:"user,omitempty"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
}

// ResumeInput represents the input for creating or updating a resume
type ResumeInput struct {
	IntroducedName  string `json:"introduced_name" validate:"required"`
	IntroducedPhone string `json:"introduced_phone" validate:"required"`
	Position        string `json:"position" validate:"required"`
	File            string `json:"file" validate:"required"`
	UserID          int    `json:"user_id" validate:"required"`
	Status          string `json:"status" validate:"omitempty,oneof=pending reviewed accepted rejected"`
}
