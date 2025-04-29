package departmenthandler

import (
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// DepartmentData represents the data structure for department responses
type DepartmentData struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	ShortName   string `json:"short_name"`
}

// CreateDepartmentResponse represents the response structure for department creation
type CreateDepartmentResponse struct {
	dto.Response
	Data DepartmentData `json:"data,omitempty"`
}

// UpdateDepartmentResponse represents the response structure for department update
type UpdateDepartmentResponse struct {
	dto.Response
	Data DepartmentData `json:"data,omitempty"`
}

// GetDepartmentResponse represents the response structure for getting a department
type GetDepartmentResponse struct {
	dto.Response
	Data DepartmentData `json:"data,omitempty"`
}

// ListDepartmentsResponse represents the response structure for listing departments
type ListDepartmentsResponse struct {
	dto.Response
	Data  []DepartmentData `json:"data,omitempty"`
	Total int              `json:"total"`
}
