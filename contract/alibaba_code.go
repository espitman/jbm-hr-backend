package contract

// AlibabaCode represents an Alibaba code in the system
type AlibabaCode struct {
	ID             int     `json:"id"`
	Code           string  `json:"code"`
	Used           bool    `json:"used"`
	CreatedAt      string  `json:"created_at"`
	AssignToUserID *int    `json:"assign_to_user_id,omitempty"`
	AssignAt       *string `json:"assign_at,omitempty"`
	Type           string  `json:"type"`
}

// CreateAlibabaCodeInput represents the input for creating a new Alibaba code
type CreateAlibabaCodeInput struct {
	Code string `json:"code" validate:"required"`
	Type string `json:"type" validate:"required,oneof=1m 3m 6m 12m 25m"`
}

// UpdateAlibabaCodeInput represents the input for updating an Alibaba code
type UpdateAlibabaCodeInput struct {
	Used           bool   `json:"used"`
	AssignToUserID *int   `json:"assign_to_user_id,omitempty"`
	AssignAt       string `json:"assign_at,omitempty"`
}

// AlibabaCodeFilters represents the filters for listing Alibaba codes
type AlibabaCodeFilters struct {
	Used           *bool   `json:"used,omitempty"`
	AssignToUserID *int    `json:"assign_to_user_id,omitempty"`
	Type           *string `json:"type,omitempty"`
}
