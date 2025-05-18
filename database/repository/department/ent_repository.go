package department

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entDepartment "github.com/espitman/jbm-hr-backend/ent/department"
)

// EntRepository implements the Repository interface using Ent
type EntRepository struct {
	client *ent.Client
}

// NewEntRepository creates a new EntRepository
func NewEntRepository(client *ent.Client) *EntRepository {
	return &EntRepository{
		client: client,
	}
}

// convertToContractDepartment converts an ent.Department to a contract.Department
func convertToContractDepartment(entDepartment *ent.Department) *contract.Department {
	if entDepartment == nil {
		return nil
	}
	return &contract.Department{
		ID:           entDepartment.ID,
		Title:        entDepartment.Title,
		Description:  entDepartment.Description,
		Image:        entDepartment.Image,
		Icon:         entDepartment.Icon,
		Color:        entDepartment.Color,
		ShortName:    entDepartment.ShortName,
		DisplayOrder: entDepartment.DisplayOrder,
	}
}

// GetAll retrieves all departments
func (r *EntRepository) GetAll(ctx context.Context) ([]*contract.Department, error) {
	entDepartments, err := r.client.Department.Query().
		Order(ent.Asc(entDepartment.FieldDisplayOrder)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	departments := make([]*contract.Department, len(entDepartments))
	for i, entDepartment := range entDepartments {
		departments[i] = convertToContractDepartment(entDepartment)
	}
	return departments, nil
}

// GetByID retrieves a department by their ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.Department, error) {
	entDepartment, err := r.client.Department.Query().Where(entDepartment.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDepartment(entDepartment), nil
}

// Create creates a new department
func (r *EntRepository) Create(ctx context.Context, req *contract.DepartmentInput) (*contract.Department, error) {
	entDepartment, err := r.client.Department.
		Create().
		SetTitle(req.Title).
		SetDescription(req.Description).
		SetImage(req.Image).
		SetIcon(req.Icon).
		SetColor(req.Color).
		SetShortName(req.ShortName).
		SetDisplayOrder(req.DisplayOrder).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDepartment(entDepartment), nil
}

// Update updates an existing department
func (r *EntRepository) Update(ctx context.Context, id int, req *contract.DepartmentInput) (*contract.Department, error) {
	entDepartment, err := r.client.Department.
		UpdateOneID(id).
		SetTitle(req.Title).
		SetDescription(req.Description).
		SetImage(req.Image).
		SetIcon(req.Icon).
		SetColor(req.Color).
		SetShortName(req.ShortName).
		SetDisplayOrder(req.DisplayOrder).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDepartment(entDepartment), nil
}

// Delete deletes a department by their ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.Department.DeleteOneID(id).Exec(ctx)
}

// List retrieves a paginated list of departments
func (r *EntRepository) List(ctx context.Context, page, limit int) ([]*contract.Department, int, error) {
	// Calculate offset
	offset := (page - 1) * limit

	// Get total count
	total, err := r.client.Department.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated departments
	entDepartments, err := r.client.Department.Query().
		Order(ent.Asc(entDepartment.FieldDisplayOrder)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Convert to contract departments
	departments := make([]*contract.Department, len(entDepartments))
	for i, entDepartment := range entDepartments {
		departments[i] = convertToContractDepartment(entDepartment)
	}

	return departments, total, nil
}

// GetTotalCount returns the total number of departments
func (r *EntRepository) GetTotalCount(ctx context.Context) (int, error) {
	return r.client.Department.Query().Count(ctx)
}
