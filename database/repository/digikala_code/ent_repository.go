package digikala_code

import (
	"context"
	"fmt"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entDigikalaCode "github.com/espitman/jbm-hr-backend/ent/digikalacode"
	"github.com/espitman/jbm-hr-backend/utils/encryption"
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
func convertToContractDigikalaCode(entDigikalaCode *ent.DigikalaCode, includeCode bool) (*contract.DigikalaCode, error) {
	if entDigikalaCode == nil {
		return nil, nil
	}

	var code string
	if includeCode {
		// Decrypt the code only if it should be included
		decryptedCode, err := encryption.Decrypt(entDigikalaCode.Code)
		if err != nil {
			return nil, err
		}
		code = decryptedCode
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

	digikalaCode := &contract.DigikalaCode{
		ID:             entDigikalaCode.ID,
		Used:           entDigikalaCode.Used,
		CreatedAt:      entDigikalaCode.CreatedAt.Format(time.RFC3339),
		AssignToUserID: assignToUserID,
		AssignAt:       assignAt,
		AssignedToUser: assignedToUser,
	}

	if includeCode {
		digikalaCode.Code = code
	}

	return digikalaCode, nil
}

// Create creates a new Digikala code
func (r *EntRepository) Create(ctx context.Context, req *contract.CreateDigikalaCodeInput) (*contract.DigikalaCode, error) {
	// Encrypt the code before storing
	encryptedCode, err := encryption.Encrypt(req.Code)
	if err != nil {
		return nil, err
	}

	entDigikalaCode, err := r.client.DigikalaCode.
		Create().
		SetCode(encryptedCode).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDigikalaCode(entDigikalaCode, true)
}

// GetAll retrieves all Digikala codes with pagination
func (r *EntRepository) GetAll(ctx context.Context, page, pageSize int, used *bool, userID *int) ([]*contract.DigikalaCode, int, error) {
	// Calculate offset
	offset := (page - 1) * pageSize

	// Build query
	query := r.client.DigikalaCode.Query()

	// Apply filters
	if used != nil {
		query = query.Where(entDigikalaCode.Used(*used))
	}
	if userID != nil {
		query = query.Where(entDigikalaCode.AssignToUserID(*userID))
	}

	// Get total count with filters
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	entDigikalaCodes, err := query.
		WithAssignedTo().
		Order(ent.Desc(entDigikalaCode.FieldID)).
		Limit(pageSize).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	codes := make([]*contract.DigikalaCode, len(entDigikalaCodes))
	for i, code := range entDigikalaCodes {
		contractCode, err := convertToContractDigikalaCode(code, false)
		if err != nil {
			return nil, 0, err
		}
		codes[i] = contractCode
	}
	return codes, total, nil
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
	return convertToContractDigikalaCode(entDigikalaCode, false)
}

// GetByCode retrieves a Digikala code by its code
func (r *EntRepository) GetByCode(ctx context.Context, code string) (*contract.DigikalaCode, error) {
	// Encrypt the code before querying
	encryptedCode, err := encryption.Encrypt(code)
	if err != nil {
		return nil, err
	}

	entDigikalaCode, err := r.client.DigikalaCode.Query().
		Where(entDigikalaCode.Code(encryptedCode)).
		WithAssignedTo().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractDigikalaCode(entDigikalaCode, true)
}

// Assign assigns a Digikala code to a user
func (r *EntRepository) Assign(ctx context.Context, code string, req *contract.AssignDigikalaCodeInput) (*contract.DigikalaCode, error) {
	// Get the DigikalaCode by ID
	codeEntity, err := r.client.DigikalaCode.Query().
		Where(entDigikalaCode.ID(req.ID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Check if code is already used
	if codeEntity.Used {
		return nil, fmt.Errorf("digikala code is already used")
	}

	codeEntity, err = codeEntity.Update().
		SetAssignToUserID(req.UserID).
		SetAssignAt(time.Now()).
		SetUsed(true).
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

	return convertToContractDigikalaCode(codeEntity, true)
}

// Delete deletes a Digikala code by its ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.DigikalaCode.DeleteOneID(id).Exec(ctx)
}
