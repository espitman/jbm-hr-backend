package contract

// Request represents a request in the system
type Request struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	FullName    string `json:"full_name"`
	Kind        string `json:"kind"` // employment, payroll_stamped, salary_deduction, introduction_letter, good_conduct_letter, confirmation_letter, embassy_letter
	Description string `json:"description,omitempty"`
	Status      string `json:"status"` // pending, doing, done, rejected
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// CreateRequestInput represents the input for creating a new request
type CreateRequestInput struct {
	UserID      int    `json:"user_id" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Kind        string `json:"kind" validate:"required,oneof=employment payroll_stamped salary_deduction introduction_letter good_conduct_letter confirmation_letter embassy_letter"`
	Description string `json:"description,omitempty"`
}

// UpdateRequestInput represents the input for updating a request
type UpdateRequestInput struct {
	Status      string `json:"status" validate:"required,oneof=pending doing done rejected"`
	Description string `json:"description,omitempty"`
}

// RequestFilter represents the filter criteria for querying requests
type RequestFilter struct {
	UserID int    `json:"user_id,omitempty"`
	Kind   string `json:"kind,omitempty"`
	Status string `json:"status,omitempty"`
}
