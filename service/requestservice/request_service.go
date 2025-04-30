package requestservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for request-related operations
type Service interface {
	// CreateRequest creates a new request
	CreateRequest(ctx context.Context, input *contract.CreateRequestInput) (*contract.Request, error)

	// UpdateRequest updates an existing request
	UpdateRequest(ctx context.Context, id int, input *contract.UpdateRequestInput) (*contract.Request, error)

	// GetRequestByID retrieves a request by its ID
	GetRequestByID(ctx context.Context, id int) (*contract.Request, error)

	// GetRequests retrieves requests based on filter criteria
	GetRequests(ctx context.Context, filter *contract.RequestFilter) ([]*contract.Request, error)

	// DeleteRequest deletes a request by its ID
	DeleteRequest(ctx context.Context, id int) error
}
