package album

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Repository defines the interface for album data access
type Repository interface {
	// GetAll retrieves all albums
	GetAll(ctx context.Context) ([]*contract.Album, error)

	// GetByID retrieves an album by its ID
	GetByID(ctx context.Context, id int) (*contract.Album, error)

	// Create creates a new album
	Create(ctx context.Context, req *contract.CreateAlbumRequest) (*contract.Album, error)

	// Update updates an existing album
	Update(ctx context.Context, id int, req *contract.UpdateAlbumRequest) (*contract.Album, error)

	// Delete deletes an album by its ID
	Delete(ctx context.Context, id int) error
}
