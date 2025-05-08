package request

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for request data access
type Repository interface {
	// Create creates a new request
	Create(ctx context.Context, req *contract.CreateRequestInput) (*contract.Request, error)

	// Update updates an existing request
	Update(ctx context.Context, id int, req *contract.UpdateRequestInput) (*contract.Request, error)

	// GetByID retrieves a request by its ID
	GetByID(ctx context.Context, id int) (*contract.Request, error)

	// GetByFilter retrieves requests based on filter criteria
	GetByFilter(ctx context.Context, filter *contract.RequestFilter) ([]*contract.Request, error)

	// GetTotalCount returns the total number of requests matching the filter criteria
	GetTotalCount(ctx context.Context, filter *contract.RequestFilter) (int, error)

	// Delete deletes a request by its ID
	Delete(ctx context.Context, id int) error
}
