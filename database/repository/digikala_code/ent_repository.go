package digikala_code

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entDigikalaCode "github.com/espitman/jbm-hr-backend/ent/digikalacode"
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

// convertToContractDigikalaCode converts an ent.DigikalaCode to a contract.DigikalaCode
func convertToContractDigikalaCode(entDigikalaCode *ent.DigikalaCode) *contract.DigikalaCode {
	if entDigikalaCode == nil {
		return nil
	}

	var assignToUserID *int
	var assignedToUser *contract.User
	if entDigikalaCode.Edges.AssignedTo != nil {
		assignToUserID = &entDigikalaCode.Edges.AssignedTo.ID
		assignedToUser = &contract.User{
			ID:        entDigikalaCode.Edges.AssignedTo.ID,
			Email:     entDigikalaCode.Edges.AssignedTo.Email,
			Phone:     entDigikalaCode.Edges.AssignedTo.Phone,
			FirstName: entDigikalaCode.Edges.AssignedTo.FirstName,
			LastName:  entDigikalaCode.Edges.AssignedTo.LastName,
			FullName:  entDigikalaCode.Edges.AssignedTo.FullName,
			Role:      string(entDigikalaCode.Edges.AssignedTo.Role),
			Avatar:    entDigikalaCode.Edges.AssignedTo.Avatar,
		}
	}

	var assignAt *string
	if !entDigikalaCode.AssignAt.IsZero() {
		assignAtStr := entDigikalaCode.AssignAt.Format(time.RFC3339)
		assignAt = &assignAtStr
	}

	return &contract.DigikalaCode{
		ID:             entDigikalaCode.ID,
		Code:           entDigikalaCode.Code,
		Used:           entDigikalaCode.Used,
		CreatedAt:      entDigikalaCode.CreatedAt.Format(time.RFC3339),
		AssignToUserID: assignToUserID,
		AssignAt:       assignAt,
		AssignedToUser: assignedToUser,
	}
}

// Create creates a new Digikala code
func (r *EntRepository) Create(ctx context.Context, req *contract.CreateDigikalaCodeInput) (*contract.DigikalaCode, error) {
	entDigikalaCode, err := r.client.DigikalaCode.
		Create().
		SetCode(req.Code).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDigikalaCode(entDigikalaCode), nil
}

// GetAll retrieves all Digikala codes
func (r *EntRepository) GetAll(ctx context.Context) ([]*contract.DigikalaCode, error) {
	entDigikalaCodes, err := r.client.DigikalaCode.Query().
		WithAssignedTo().
		Order(ent.Asc(entDigikalaCode.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	codes := make([]*contract.DigikalaCode, len(entDigikalaCodes))
	for i, code := range entDigikalaCodes {
		codes[i] = convertToContractDigikalaCode(code)
	}
	return codes, nil
}

// GetByID retrieves a Digikala code by its ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.DigikalaCode, error) {
	entDigikalaCode, err := r.client.DigikalaCode.Query().
		Where(entDigikalaCode.ID(id)).
		WithAssignedTo().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDigikalaCode(entDigikalaCode), nil
}

// GetByCode retrieves a Digikala code by its code
func (r *EntRepository) GetByCode(ctx context.Context, code string) (*contract.DigikalaCode, error) {
	entDigikalaCode, err := r.client.DigikalaCode.Query().
		Where(entDigikalaCode.Code(code)).
		WithAssignedTo().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDigikalaCode(entDigikalaCode), nil
}

// Assign assigns a Digikala code to a user
func (r *EntRepository) Assign(ctx context.Context, code string, req *contract.AssignDigikalaCodeInput) (*contract.DigikalaCode, error) {
	codeEntity, err := r.client.DigikalaCode.Query().
		Where(entDigikalaCode.Code(code)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	codeEntity, err = codeEntity.Update().
		SetAssignToUserID(req.UserID).
		SetAssignAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Reload with edges
	codeEntity, err = r.client.DigikalaCode.Query().
		Where(entDigikalaCode.ID(codeEntity.ID)).
		WithAssignedTo().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return convertToContractDigikalaCode(codeEntity), nil
}

// Delete deletes a Digikala code by its ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.DigikalaCode.DeleteOneID(id).Exec(ctx)
}
