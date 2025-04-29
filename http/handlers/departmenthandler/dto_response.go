package departmenthandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// CreateDepartmentResponse represents the response structure for department creation
type CreateDepartmentResponse struct {
	dto.Response
	Data contract.Department `json:"data,omitempty"`
}

// UpdateDepartmentResponse represents the response structure for department update
type UpdateDepartmentResponse struct {
	dto.Response
	Data contract.Department `json:"data,omitempty"`
}

// GetDepartmentResponse represents the response structure for getting a department
type GetDepartmentResponse struct {
	dto.Response
	Data contract.Department `json:"data,omitempty"`
}

// DepartmentListData represents the data structure for department list responses
type DepartmentListData struct {
	Departments []*contract.Department `json:"departments,omitempty"`
	Total       int                    `json:"total"`
}

// ListDepartmentsResponse represents the response structure for listing departments
type ListDepartmentsResponse struct {
	dto.Response
	Data DepartmentListData `json:"data,omitempty"`
}
