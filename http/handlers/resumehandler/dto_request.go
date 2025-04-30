package resumehandler

// CreateResumeRequest represents the request body for creating a resume
type CreateResumeRequest struct {
	IntroducedName  string `json:"introduced_name" validate:"required"`
	IntroducedPhone string `json:"introduced_phone" validate:"required"`
	Position        string `json:"position" validate:"required"`
	File            string `json:"file" validate:"required"`
}

// UpdateStatusRequest represents the request body for updating a resume's status
type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending reviewed accepted rejected"`
}
