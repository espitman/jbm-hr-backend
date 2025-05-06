package request

import (
	"context"
	"time"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entrequest "github.com/espitman/jbm-hr-backend/ent/request"
	"github.com/espitman/jbm-hr-backend/ent/user"
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
	req := &contract.Request{
		ID:          entReq.ID,
		UserID:      entReq.UserID,
		FullName:    entReq.FullName,
		Kind:        string(entReq.Kind),
		Description: description,
		Status:      string(entReq.Status),
		CreatedAt:   entReq.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   entReq.UpdatedAt.Format(time.RFC3339),
	}

	if entReq.Edges.User != nil {
		req.User = contract.RequestUser{
			ID:        entReq.Edges.User.ID,
			Email:     entReq.Edges.User.Email,
			FirstName: entReq.Edges.User.FirstName,
			LastName:  entReq.Edges.User.LastName,
			Avatar:    entReq.Edges.User.Avatar,
		}
	}

	// Convert metadata
	if entReq.Edges.Meta != nil {
		req.Meta = make([]contract.RequestMeta, len(entReq.Edges.Meta))
		for i, meta := range entReq.Edges.Meta {
			req.Meta[i] = contract.RequestMeta{
				Key:   meta.Key,
				Value: meta.Value,
			}
		}
	}

	return req
}

// Create creates a new request
func (r *EntRepository) Create(ctx context.Context, req *contract.CreateRequestInput) (*contract.Request, error) {
	// Start a transaction
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Create the request
	entReq, err := tx.Request.
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
		tx.Rollback()
		return nil, err
	}

	// Create metadata if provided
	if req.Meta != nil {
		metaBuilders := make([]*ent.RequestMetaCreate, len(req.Meta))
		for i, meta := range req.Meta {
			metaBuilders[i] = tx.RequestMeta.
				Create().
				SetRequestID(entReq.ID).
				SetKey(meta.Key).
				SetValue(meta.Value)
		}
		_, err = tx.RequestMeta.CreateBulk(metaBuilders...).Save(ctx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Return the created request with all its relations
	return r.GetByID(ctx, entReq.ID)
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
	return r.GetByID(ctx, entReq.ID)
}

// GetByID retrieves a request by its ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.Request, error) {
	entReq, err := r.client.Request.Query().
		WithUser(func(q *ent.UserQuery) {
			q.Select(user.FieldID, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldAvatar)
		}).
		WithMeta().
		Where(entrequest.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractRequest(entReq), nil
}

// GetByFilter retrieves requests based on filter criteria
func (r *EntRepository) GetByFilter(ctx context.Context, filter *contract.RequestFilter) ([]*contract.Request, error) {
	query := r.client.Request.Query().
		WithUser(func(q *ent.UserQuery) {
			q.Select(user.FieldID, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldAvatar)
		}).
		WithMeta().
		Order(ent.Desc(entrequest.FieldID))

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
