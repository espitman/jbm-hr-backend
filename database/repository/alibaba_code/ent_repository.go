package alibaba_code

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entAlibabaCode "github.com/espitman/jbm-hr-backend/ent/alibabacode"
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

// convertToContractAlibabaCode converts an ent.AlibabaCode to a contract.AlibabaCode
func convertToContractAlibabaCode(entCode *ent.AlibabaCode, includeCode bool) (*contract.AlibabaCode, error) {
	if entCode == nil {
		return nil, nil
	}
	var assignToUserID *int
	if entCode.Edges.AssignedToUser != nil {
		assignToUserID = &entCode.Edges.AssignedToUser.ID
	}
	var assignAt *string
	if entCode.AssignAt != nil {
		assignAtStr := entCode.AssignAt.Format(time.RFC3339)
		assignAt = &assignAtStr
	}

	var code string
	if includeCode {
		// Decrypt the code only if it should be included
		decryptedCode, err := encryption.Decrypt(entCode.Code)
		if err != nil {
			return nil, err
		}
		code = decryptedCode
	}

	return &contract.AlibabaCode{
		ID:             entCode.ID,
		Code:           code,
		Used:           entCode.Used,
		CreatedAt:      entCode.CreatedAt.Format(time.RFC3339),
		AssignToUserID: assignToUserID,
		AssignAt:       assignAt,
		Type:           string(entCode.Type),
	}, nil
}

// GetAll retrieves all Alibaba codes
func (r *EntRepository) GetAll(ctx context.Context, filters *contract.AlibabaCodeFilters) ([]*contract.AlibabaCode, error) {
	query := r.client.AlibabaCode.Query().
		WithAssignedToUser()

	// Apply filters
	if filters != nil {
		if filters.Used != nil {
			query = query.Where(entAlibabaCode.UsedEQ(*filters.Used))
		}
		if filters.AssignToUserID != nil {
			query = query.Where(entAlibabaCode.AssignToUserIDEQ(*filters.AssignToUserID))
		}
		if filters.Type != nil {
			query = query.Where(entAlibabaCode.TypeEQ(entAlibabaCode.Type(*filters.Type)))
		}
	}

	entCodes, err := query.Order(ent.Desc(entAlibabaCode.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}

	codes := make([]*contract.AlibabaCode, len(entCodes))
	for i, entCode := range entCodes {
		contractCode, err := convertToContractAlibabaCode(entCode, false)
		if err != nil {
			return nil, err
		}
		codes[i] = contractCode
	}
	return codes, nil
}

// GetByID retrieves an Alibaba code by its ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.AlibabaCode, error) {
	entCode, err := r.client.AlibabaCode.Query().
		Where(entAlibabaCode.ID(id)).
		WithAssignedToUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractAlibabaCode(entCode, false)
}

// GetByCode retrieves an Alibaba code by its code
func (r *EntRepository) GetByCode(ctx context.Context, code string) (*contract.AlibabaCode, error) {
	// Encrypt the code before querying
	encryptedCode, err := encryption.Encrypt(code)
	if err != nil {
		return nil, err
	}

	entCode, err := r.client.AlibabaCode.Query().
		Where(entAlibabaCode.Code(encryptedCode)).
		WithAssignedToUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractAlibabaCode(entCode, true)
}

// Create creates a new Alibaba code
func (r *EntRepository) Create(ctx context.Context, req *contract.CreateAlibabaCodeInput) (*contract.AlibabaCode, error) {
	// Encrypt the code before storing
	encryptedCode, err := encryption.Encrypt(req.Code)
	if err != nil {
		return nil, err
	}

	entCode, err := r.client.AlibabaCode.
		Create().
		SetCode(encryptedCode).
		SetType(entAlibabaCode.Type(req.Type)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractAlibabaCode(entCode, true)
}

// Update updates an existing Alibaba code
func (r *EntRepository) Update(ctx context.Context, id int, req *contract.UpdateAlibabaCodeInput) (*contract.AlibabaCode, error) {
	update := r.client.AlibabaCode.
		UpdateOneID(id).
		SetUsed(req.Used)

	if req.AssignToUserID != nil {
		update = update.SetAssignToUserID(*req.AssignToUserID)
	}
	if req.AssignAt != "" {
		assignAt, err := time.Parse(time.RFC3339, req.AssignAt)
		if err != nil {
			return nil, err
		}
		update = update.SetAssignAt(assignAt)
	}

	entCode, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractAlibabaCode(entCode, false)
}

// Delete deletes an Alibaba code by its ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.AlibabaCode.DeleteOneID(id).Exec(ctx)
}
