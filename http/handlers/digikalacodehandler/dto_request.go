package digikalacodehandler

// CreateDigikalaCodeRequest represents the request body for creating a Digikala code
type CreateDigikalaCodeRequest struct {
	Code string `json:"code" validate:"required"`
}

// AssignDigikalaCodeRequest represents the request body for assigning a Digikala code
type AssignDigikalaCodeRequest struct {
	UserID int `json:"user_id" validate:"required"`
}
