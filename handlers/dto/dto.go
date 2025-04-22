package dto

// Response represents the standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string) Response {
	return Response{
		Success: false,
		Message: message,
	}
}

// NewAlbumResponse creates a new album response
func NewAlbumResponse(album interface{}) Response {
	return Response{
		Success: true,
		Message: "Operation successful",
		Data:    album,
	}
}
