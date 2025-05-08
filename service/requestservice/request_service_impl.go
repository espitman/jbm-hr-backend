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

func (s *service) GetRequests(ctx context.Context, filter *contract.RequestFilter) ([]*contract.Request, int, error) {
	// Create channels for results
	type result struct {
		requests []*contract.Request
		err      error
	}
	type countResult struct {
		count int
		err   error
	}

	reqChan := make(chan result)
	countChan := make(chan countResult)

	// Get total count in a goroutine
	go func() {
		total, err := s.requestRepo.GetTotalCount(ctx, filter)
		countChan <- countResult{count: total, err: err}
	}()

	// Get paginated results in a goroutine
	go func() {
		requests, err := s.requestRepo.GetByFilter(ctx, filter)
		reqChan <- result{requests: requests, err: err}
	}()

	// Wait for both results
	var requests []*contract.Request
	var total int
	var reqErr, countErr error

	for i := 0; i < 2; i++ {
		select {
		case req := <-reqChan:
			requests = req.requests
			reqErr = req.err
		case count := <-countChan:
			total = count.count
			countErr = count.err
		}
	}

	// Check for errors
	if reqErr != nil {
		return nil, 0, reqErr
	}
	if countErr != nil {
		return nil, 0, countErr
	}

	return requests, total, nil
}

func (s *service) DeleteRequest(ctx context.Context, id int) error {
	return s.requestRepo.Delete(ctx, id)
}
