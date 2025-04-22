package albumservice

import (
	"context"
	"gin-project/ent"
	"gin-project/ent/album"
)

// AlbumService handles album-related business logic
type AlbumService struct {
	client *ent.Client
}

// New creates a new AlbumService instance with mock data
func New(client *ent.Client) *AlbumService {
	return &AlbumService{
		client: client,
	}
}

// GetAllAlbums returns all albums
func (s *AlbumService) GetAllAlbums(ctx context.Context) ([]*ent.Album, error) {
	return s.client.Album.Query().All(ctx)
}

// CreateAlbum creates a new album
func (s *AlbumService) CreateAlbum(ctx context.Context, url, caption string) (*ent.Album, error) {
	return s.client.Album.
		Create().
		SetURL(url).
		SetCaption(caption).
		Save(ctx)
}

// GetAlbumByID returns an album by its ID
func (s *AlbumService) GetAlbumByID(ctx context.Context, id int) (*ent.Album, error) {
	return s.client.Album.Query().Where(album.ID(id)).Only(ctx)
}

// UpdateAlbum updates an existing album
func (s *AlbumService) UpdateAlbum(ctx context.Context, id int, url, caption string) (*ent.Album, error) {
	return s.client.Album.
		UpdateOneID(id).
		SetURL(url).
		SetCaption(caption).
		Save(ctx)
}

// DeleteAlbum deletes an album by its ID
func (s *AlbumService) DeleteAlbum(ctx context.Context, id int) error {
	return s.client.Album.DeleteOneID(id).Exec(ctx)
}

// Add your album-related methods here that use the Ent client
// For example:
// func (s *AlbumService) CreateAlbum(ctx context.Context, album *ent.Album) (*ent.Album, error) {
//     return s.client.Album.Create()...
// }
