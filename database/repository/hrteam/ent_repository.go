package hrteam

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entHRTeam "github.com/espitman/jbm-hr-backend/ent/hrteam"
)

// EntRepository implements the Repository interface using Ent
type EntRepository struct {
	client *ent.Client
}

// NewEntRepository creates a new EntRepository
func NewEntRepository(client *ent.Client) Repository {
	return &EntRepository{
		client: client,
	}
}

// convertToContractHRTeam converts an ent.HRTeam to a contract.HRTeam
func convertToContractHRTeam(entHRTeam *ent.HRTeam) *contract.HRTeam {
	if entHRTeam == nil {
		return nil
	}
	return &contract.HRTeam{
		ID:           entHRTeam.ID,
		FullName:     entHRTeam.FullName,
		Role:         entHRTeam.Role,
		Email:        entHRTeam.Email,
		Phone:        entHRTeam.Phone,
		DisplayOrder: entHRTeam.DisplayOrder,
	}
}

// GetAll retrieves all HR team members
func (r *EntRepository) GetAll(ctx context.Context) ([]*contract.HRTeam, error) {
	entHRTeams, err := r.client.HRTeam.Query().
		Order(ent.Asc(entHRTeam.FieldDisplayOrder)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	hrTeams := make([]*contract.HRTeam, len(entHRTeams))
	for i, entHRTeam := range entHRTeams {
		hrTeams[i] = convertToContractHRTeam(entHRTeam)
	}
	return hrTeams, nil
}

// GetByID retrieves an HR team member by their ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.HRTeam, error) {
	entHRTeam, err := r.client.HRTeam.Query().Where(entHRTeam.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractHRTeam(entHRTeam), nil
}

// Create creates a new HR team member
func (r *EntRepository) Create(ctx context.Context, req *contract.HRTeamInput) (*contract.HRTeam, error) {
	entHRTeam, err := r.client.HRTeam.
		Create().
		SetFullName(req.FullName).
		SetRole(req.Role).
		SetEmail(req.Email).
		SetPhone(req.Phone).
		SetDisplayOrder(req.DisplayOrder).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractHRTeam(entHRTeam), nil
}

// Update updates an existing HR team member
func (r *EntRepository) Update(ctx context.Context, id int, req *contract.HRTeamInput) (*contract.HRTeam, error) {
	entHRTeam, err := r.client.HRTeam.
		UpdateOneID(id).
		SetFullName(req.FullName).
		SetRole(req.Role).
		SetEmail(req.Email).
		SetPhone(req.Phone).
		SetDisplayOrder(req.DisplayOrder).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractHRTeam(entHRTeam), nil
}

// Delete deletes an HR team member by their ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.HRTeam.DeleteOneID(id).Exec(ctx)
}
