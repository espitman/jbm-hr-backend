package albumservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	albumRepo "github.com/espitman/jbm-hr-backend/repository/album"
)

// AlbumService handles album-related business logic
type AlbumService struct {
	repository albumRepo.Repository
}

// New creates a new AlbumService instance with the provided repository
func New(repository albumRepo.Repository) *AlbumService {
	return &AlbumService{
		repository: repository,
	}
}

// GetAllAlbums returns all albums
func (s *AlbumService) GetAllAlbums(ctx context.Context) ([]*contract.Album, error) {
	return s.repository.GetAll(ctx)
}

// CreateAlbum creates a new album
func (s *AlbumService) CreateAlbum(ctx context.Context, req *contract.CreateAlbumInput) (*contract.Album, error) {
	return s.repository.Create(ctx, req)
}

// GetAlbumByID returns an album by its ID
func (s *AlbumService) GetAlbumByID(ctx context.Context, id int) (*contract.Album, error) {
	return s.repository.GetByID(ctx, id)
}

// UpdateAlbum updates an existing album
func (s *AlbumService) UpdateAlbum(ctx context.Context, id int, req *contract.UpdateAlbumInput) (*contract.Album, error) {
	return s.repository.Update(ctx, id, req)
}

// DeleteAlbum deletes an album by its ID
func (s *AlbumService) DeleteAlbum(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

// Add your album-related methods here that use the Ent client
// For example:
// func (s *AlbumService) CreateAlbum(ctx context.Context, album *ent.Album) (*ent.Album, error) {
//     return s.client.Album.Create()...
// }
