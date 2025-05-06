package requesthandler

// CreateRequestRequest represents the request for creating a request
type CreateRequestRequest struct {
	FullName    string  `json:"full_name" validate:"required"`
	Kind        string  `json:"kind" validate:"required,oneof=employment payroll_stamped salary_deduction introduction_letter good_conduct_letter confirmation_letter embassy_letter development_learning"`
	Description *string `json:"description,omitempty"`
	Meta        []struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	} `json:"meta,omitempty"`
}

// UpdateRequestStatusRequest represents the request for updating a request's status
type UpdateRequestStatusRequest struct {
	Status      string  `json:"status" validate:"required,oneof=pending doing done rejected"`
	Description *string `json:"description,omitempty"`
}

// GetRequestsRequest represents the request for getting requests
type GetRequestsRequest struct {
	Page     int    `query:"page" validate:"required"`
	PageSize int    `query:"page_size" validate:"required"`
	Status   string `query:"status"`
	Kind     string `query:"kind"`
	UserID   int    `query:"user_id"`
}
