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
func convertToContractDepartment(entDept *ent.Department) *contract.Department {
	if entDept == nil {
		return nil
	}
	return &contract.Department{
		ID:          entDept.ID,
		Title:       entDept.Title,
		Description: entDept.Description,
		Image:       entDept.Image,
		Icon:        entDept.Icon,
		Color:       entDept.Color,
		ShortName:   entDept.ShortName,
	}
}

// GetAll retrieves all departments
func (r *EntRepository) GetAll(ctx context.Context) ([]*contract.Department, error) {
	entDepts, err := r.client.Department.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	depts := make([]*contract.Department, len(entDepts))
	for i, entDept := range entDepts {
		depts[i] = convertToContractDepartment(entDept)
	}
	return depts, nil
}

// GetByID retrieves a department by its ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.Department, error) {
	entDept, err := r.client.Department.Query().Where(entDepartment.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDepartment(entDept), nil
}

// Create creates a new department
func (r *EntRepository) Create(ctx context.Context, req *contract.DepartmentInput) (*contract.Department, error) {
	entDept, err := r.client.Department.
		Create().
		SetTitle(req.Title).
		SetDescription(req.Description).
		SetImage(req.Image).
		SetIcon(req.Icon).
		SetColor(req.Color).
		SetShortName(req.ShortName).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDepartment(entDept), nil
}

// Update updates an existing department
func (r *EntRepository) Update(ctx context.Context, id int, req *contract.DepartmentInput) (*contract.Department, error) {
	entDept, err := r.client.Department.
		UpdateOneID(id).
		SetTitle(req.Title).
		SetDescription(req.Description).
		SetImage(req.Image).
		SetIcon(req.Icon).
		SetColor(req.Color).
		SetShortName(req.ShortName).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDepartment(entDept), nil
}

// Delete deletes a department by its ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.Department.DeleteOneID(id).Exec(ctx)
}
