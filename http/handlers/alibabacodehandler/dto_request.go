package alibabacodehandler

// CreateAlibabaCodeRequest represents the request body for creating an Alibaba code
type CreateAlibabaCodeRequest struct {
	Code string `json:"code" validate:"required"`
	Type string `json:"type" validate:"required,oneof=1m 3m 6m 12m 25m"`
}

// AssignAlibabaCodeRequest represents the request body for assigning an Alibaba code
type AssignAlibabaCodeRequest struct {
	UserID int `json:"user_id" validate:"required"`
}
