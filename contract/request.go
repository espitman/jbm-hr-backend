package contract

// RequestUser represents minimal user information for request responses
type RequestUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

// RequestMeta represents a key-value pair for request metadata
type RequestMeta struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// Request represents a request in the system
type Request struct {
	ID          int           `json:"id"`
	UserID      int           `json:"user_id"`
	User        RequestUser   `json:"user,omitempty"`
	FullName    string        `json:"full_name"`
	Kind        string        `json:"kind"` // employment, payroll_stamped, salary_deduction, introduction_letter, good_conduct_letter, confirmation_letter, embassy_letter
	Description *string       `json:"description,omitempty"`
	Status      string        `json:"status"` // pending, doing, done, rejected
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	Meta        []RequestMeta `json:"meta,omitempty"`
}

// CreateRequestInput represents the input for creating a new request
type CreateRequestInput struct {
	UserID      int           `json:"user_id" validate:"required"`
	FullName    string        `json:"full_name" validate:"required"`
	Kind        string        `json:"kind" validate:"required,oneof=employment payroll_stamped salary_deduction introduction_letter good_conduct_letter confirmation_letter embassy_letter development_learning marriage_gift childbirth_gift travel_credit supplementary_insurance"`
	Description *string       `json:"description,omitempty"`
	Meta        []RequestMeta `json:"meta,omitempty"`
}

// UpdateRequestInput represents the input for updating a request
type UpdateRequestInput struct {
	Status      string  `json:"status" validate:"required,oneof=pending doing done rejected"`
	Description *string `json:"description,omitempty"`
}

// RequestFilter represents the filter criteria for querying requests
type RequestFilter struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	UserID   int    `json:"user_id,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Status   string `json:"status,omitempty"`
}
