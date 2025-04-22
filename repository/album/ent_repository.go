package album

import (
	"context"

	"github.com/espitman/jbm-hr-backend/ent"
	"github.com/espitman/jbm-hr-backend/ent/album"
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

// GetAll retrieves all albums
func (r *EntRepository) GetAll(ctx context.Context) ([]*ent.Album, error) {
	return r.client.Album.Query().All(ctx)
}

// GetByID retrieves an album by its ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*ent.Album, error) {
	return r.client.Album.Query().Where(album.ID(id)).Only(ctx)
}

// Create creates a new album
func (r *EntRepository) Create(ctx context.Context, url, caption string) (*ent.Album, error) {
	return r.client.Album.
		Create().
		SetURL(url).
		SetCaption(caption).
		Save(ctx)
}

// Update updates an existing album
func (r *EntRepository) Update(ctx context.Context, id int, url, caption string) (*ent.Album, error) {
	return r.client.Album.
		UpdateOneID(id).
		SetURL(url).
		SetCaption(caption).
		Save(ctx)
}

// Delete deletes an album by its ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.Album.DeleteOneID(id).Exec(ctx)
}
