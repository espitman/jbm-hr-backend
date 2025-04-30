package request

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entrequest "github.com/espitman/jbm-hr-backend/ent/request"
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

// convertToContractRequest converts an ent.Request to a contract.Request
func convertToContractRequest(entReq *ent.Request) *contract.Request {
	if entReq == nil {
		return nil
	}
	var description *string
	if entReq.Description != "" {
		description = &entReq.Description
	}
	return &contract.Request{
		ID:          entReq.ID,
		UserID:      entReq.UserID,
		FullName:    entReq.FullName,
		Kind:        string(entReq.Kind),
		Description: description,
		Status:      string(entReq.Status),
		CreatedAt:   entReq.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   entReq.UpdatedAt.Format(time.RFC3339),
	}
}

// Create creates a new request
func (r *EntRepository) Create(ctx context.Context, req *contract.CreateRequestInput) (*contract.Request, error) {
	entReq, err := r.client.Request.
		Create().
		SetUserID(req.UserID).
		SetFullName(req.FullName).
		SetKind(entrequest.Kind(req.Kind)).
		SetNillableDescription(req.Description).
		SetStatus(entrequest.StatusPending).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractRequest(entReq), nil
}

// Update updates an existing request
func (r *EntRepository) Update(ctx context.Context, id int, req *contract.UpdateRequestInput) (*contract.Request, error) {
	entReq, err := r.client.Request.
		UpdateOneID(id).
		SetStatus(entrequest.Status(req.Status)).
		SetNillableDescription(req.Description).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractRequest(entReq), nil
}

// GetByID retrieves a request by its ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.Request, error) {
	entReq, err := r.client.Request.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertToContractRequest(entReq), nil
}

// GetByFilter retrieves requests based on filter criteria
func (r *EntRepository) GetByFilter(ctx context.Context, filter *contract.RequestFilter) ([]*contract.Request, error) {
	query := r.client.Request.Query()

	if filter.UserID != 0 {
		query = query.Where(entrequest.UserID(filter.UserID))
	}
	if filter.Kind != "" {
		query = query.Where(entrequest.KindEQ(entrequest.Kind(filter.Kind)))
	}
	if filter.Status != "" {
		query = query.Where(entrequest.StatusEQ(entrequest.Status(filter.Status)))
	}

	entReqs, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	reqs := make([]*contract.Request, len(entReqs))
	for i, entReq := range entReqs {
		reqs[i] = convertToContractRequest(entReq)
	}
	return reqs, nil
}

// Delete deletes a request by its ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.Request.DeleteOneID(id).Exec(ctx)
}
