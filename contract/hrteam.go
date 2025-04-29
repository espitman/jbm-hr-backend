package contract

// HRTeam represents the HR team member data structure
type HRTeam struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// CreateHRTeamInput represents the request to create a new HR team member
type CreateHRTeamInput struct {
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
}

// UpdateHRTeamInput represents the request to update an existing HR team member
type UpdateHRTeamInput struct {
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
}
