package contract

// Department represents a department in the system
type Department struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Icon         string `json:"icon"`
	Color        string `json:"color"`
	ShortName    string `json:"short_name"`
	DisplayOrder int    `json:"display_order"`
}

// DepartmentInput represents the input for creating or updating a department
type DepartmentInput struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Image        string `json:"image" validate:"required"`
	Icon         string `json:"icon" validate:"required"`
	Color        string `json:"color" validate:"required"`
	ShortName    string `json:"short_name" validate:"required"`
	DisplayOrder int    `json:"display_order"`
}
