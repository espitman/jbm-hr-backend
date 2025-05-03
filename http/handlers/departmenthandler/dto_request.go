package departmenthandler

// CreateDepartmentRequest represents the request structure for creating a department
type CreateDepartmentRequest struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Image        string `json:"image" validate:"required,url"`
	Icon         string `json:"icon" validate:"required"`
	Color        string `json:"color" validate:"required"`
	ShortName    string `json:"short_name" validate:"required"`
	DisplayOrder int    `json:"display_order"`
}

// UpdateDepartmentRequest represents the request structure for updating a department
type UpdateDepartmentRequest struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Image        string `json:"image" validate:"required,url"`
	Icon         string `json:"icon" validate:"required"`
	Color        string `json:"color" validate:"required"`
	ShortName    string `json:"short_name" validate:"required"`
	DisplayOrder int    `json:"display_order"`
}
