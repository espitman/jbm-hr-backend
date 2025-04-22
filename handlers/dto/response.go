package dto

// Response represents the base response structure for all API responses
type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message,omitempty" example:"Operation completed successfully"`
}
