package contract

// HRTeam represents the HR team member data structure
type HRTeam struct {
	ID           int    `json:"id"`
	FullName     string `json:"full_name"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	DisplayOrder int    `json:"display_order"`
}

// HRTeamInput represents the request to create or update an HR team member
type HRTeamInput struct {
	FullName     string `json:"full_name" binding:"required"`
	Role         string `json:"role" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone" binding:"required"`
	DisplayOrder int    `json:"display_order"`
}
