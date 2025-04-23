package dto

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response represents a generic API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string) *Response {
	return &Response{
		Success: true,
		Message: message,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string) *Response {
	return &Response{
		Success: false,
		Message: message,
	}
}

// NewDataResponse creates a new response with data
func NewDataResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Message: "ok",
		Data:    data,
	}
}

// JSON sends a JSON response with the given status code and data
func JSON(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}

// SuccessJSON sends a success JSON response with the given data
func SuccessJSON(c echo.Context, data interface{}) error {
	return JSON(c, http.StatusOK, NewDataResponse(data))
}

// CreatedJSON sends a created JSON response with the given data
func CreatedJSON(c echo.Context, data interface{}) error {
	return JSON(c, http.StatusCreated, NewDataResponse(data))
}

// ErrorJSON sends an error JSON response with the given message and status code
func ErrorJSON(c echo.Context, status int, message string) error {
	return JSON(c, status, NewErrorResponse(message))
}

// BadRequestJSON sends a bad request JSON response with the given message
func BadRequestJSON(c echo.Context, message string) error {
	return ErrorJSON(c, http.StatusBadRequest, message)
}

// InternalServerErrorJSON sends an internal server error JSON response with the given message
func InternalServerErrorJSON(c echo.Context, message string) error {
	return ErrorJSON(c, http.StatusInternalServerError, message)
}
