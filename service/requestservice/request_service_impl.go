package requestservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/database/repository/request"
)

type service struct {
	requestRepo request.Repository
}

// New creates a new request service
func New(requestRepo request.Repository) Service {
	return &service{
		requestRepo: requestRepo,
	}
}

func (s *service) CreateRequest(ctx context.Context, input *contract.CreateRequestInput) (*contract.Request, error) {
	return s.requestRepo.Create(ctx, input)
}

func (s *service) UpdateRequest(ctx context.Context, id int, input *contract.UpdateRequestInput) (*contract.Request, error) {
	return s.requestRepo.Update(ctx, id, input)
}

func (s *service) GetRequestByID(ctx context.Context, id int) (*contract.Request, error) {
	return s.requestRepo.GetByID(ctx, id)
}

func (s *service) GetRequests(ctx context.Context, filter *contract.RequestFilter) ([]*contract.Request, error) {
	return s.requestRepo.GetByFilter(ctx, filter)
}

func (s *service) DeleteRequest(ctx context.Context, id int) error {
	return s.requestRepo.Delete(ctx, id)
}
