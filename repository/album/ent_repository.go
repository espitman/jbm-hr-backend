package album

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/ent"
	entAlbum "github.com/espitman/jbm-hr-backend/ent/album"
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

// convertToContractAlbum converts an ent.Album to a contract.Album
func convertToContractAlbum(entAlbum *ent.Album) *contract.Album {
	if entAlbum == nil {
		return nil
	}
	return &contract.Album{
		ID:      entAlbum.ID,
		URL:     entAlbum.URL,
		Caption: entAlbum.Caption,
	}
}

// GetAll retrieves all albums
func (r *EntRepository) GetAll(ctx context.Context) ([]*contract.Album, error) {
	entAlbums, err := r.client.Album.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	albums := make([]*contract.Album, len(entAlbums))
	for i, entAlbum := range entAlbums {
		albums[i] = convertToContractAlbum(entAlbum)
	}
	return albums, nil
}

// GetByID retrieves an album by its ID
func (r *EntRepository) GetByID(ctx context.Context, id int) (*contract.Album, error) {
	entAlbum, err := r.client.Album.Query().Where(entAlbum.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractAlbum(entAlbum), nil
}

// Create creates a new album
func (r *EntRepository) Create(ctx context.Context, req *contract.CreateAlbumInput) (*contract.Album, error) {
	entAlbum, err := r.client.Album.
		Create().
		SetURL(req.URL).
		SetCaption(req.Caption).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractAlbum(entAlbum), nil
}

// Update updates an existing album
func (r *EntRepository) Update(ctx context.Context, id int, req *contract.UpdateAlbumInput) (*contract.Album, error) {
	entAlbum, err := r.client.Album.
		UpdateOneID(id).
		SetURL(req.URL).
		SetCaption(req.Caption).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return convertToContractAlbum(entAlbum), nil
}

// Delete deletes an album by its ID
func (r *EntRepository) Delete(ctx context.Context, id int) error {
	return r.client.Album.DeleteOneID(id).Exec(ctx)
}
