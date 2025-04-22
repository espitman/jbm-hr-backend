package album

import (
	"context"

	"github.com/espitman/jbm-hr-backend/ent"
)

// Repository defines the interface for album data access
type Repository interface {
	// GetAll retrieves all albums
	GetAll(ctx context.Context) ([]*ent.Album, error)

	// GetByID retrieves an album by its ID
	GetByID(ctx context.Context, id int) (*ent.Album, error)

	// Create creates a new album
	Create(ctx context.Context, url, caption string) (*ent.Album, error)

	// Update updates an existing album
	Update(ctx context.Context, id int, url, caption string) (*ent.Album, error)

	// Delete deletes an album by its ID
	Delete(ctx context.Context, id int) error
}
