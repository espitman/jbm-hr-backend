package alibaba_code

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for Alibaba code data access
type Repository interface {
	// GetAll retrieves all Alibaba codes
	GetAll(ctx context.Context, filters *contract.AlibabaCodeFilters) ([]*contract.AlibabaCode, error)

	// GetByID retrieves an Alibaba code by its ID
	GetByID(ctx context.Context, id int) (*contract.AlibabaCode, error)

	// GetByCode retrieves an Alibaba code by its code
	GetByCode(ctx context.Context, code string) (*contract.AlibabaCode, error)

	// Create creates a new Alibaba code
	Create(ctx context.Context, req *contract.CreateAlibabaCodeInput) (*contract.AlibabaCode, error)

	// Update updates an existing Alibaba code
	Update(ctx context.Context, id int, req *contract.UpdateAlibabaCodeInput) (*contract.AlibabaCode, error)

	// Delete deletes an Alibaba code by its ID
	Delete(ctx context.Context, id int) error
}
