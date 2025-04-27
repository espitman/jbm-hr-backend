package albumservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
	albumRepo "github.com/espitman/jbm-hr-backend/database/repository/album"
)

// albumService implements the Service interface
type albumService struct {
	repository albumRepo.Repository
}

// New creates a new album service instance
func New(repository albumRepo.Repository) Service {
	return &albumService{
		repository: repository,
	}
}

// GetAllAlbums returns all albums
func (s *albumService) GetAllAlbums(ctx context.Context) ([]*contract.Album, error) {
	return s.repository.GetAll(ctx)
}

// CreateAlbum creates a new album
func (s *albumService) CreateAlbum(ctx context.Context, input *contract.CreateAlbumInput) (*contract.Album, error) {
	return s.repository.Create(ctx, input)
}

// GetAlbumByID returns an album by its ID
func (s *albumService) GetAlbumByID(ctx context.Context, id int) (*contract.Album, error) {
	return s.repository.GetByID(ctx, id)
}

// UpdateAlbum updates an existing album
func (s *albumService) UpdateAlbum(ctx context.Context, id int, input *contract.UpdateAlbumInput) (*contract.Album, error) {
	return s.repository.Update(ctx, id, input)
}

// DeleteAlbum deletes an album by its ID
func (s *albumService) DeleteAlbum(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
