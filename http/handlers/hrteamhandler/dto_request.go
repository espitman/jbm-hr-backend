package hrteamhandler

// HRTeamInput represents the input data for creating or updating an HR team member
type HRTeamInput struct {
	FullName string `json:"fullName" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
}
