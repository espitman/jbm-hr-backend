package contract

import "errors"

// Department represents a department in the system
type Department struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	ShortName   string `json:"short_name"`
}

// DepartmentInput represents the input for creating or updating a department
type DepartmentInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
	Color       string `json:"color" validate:"required"`
	ShortName   string `json:"short_name" validate:"required"`
}

// Error definitions for department-related operations
var (
	ErrDepartmentNotFound = errors.New("department not found")
	ErrDepartmentExists   = errors.New("department already exists")
)
