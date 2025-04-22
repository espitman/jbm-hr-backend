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
func (s *AlbumService) GetAllAlbums(ctx context.Context) (*contract.AlbumsResponse, error) {
	albums, err := s.repository.GetAll(ctx)
	if err != nil {
		return &contract.AlbumsResponse{
			Success: false,
			Message: "Failed to get albums",
		}, err
	}

	// Convert []*contract.Album to []contract.Album
	albumSlice := make([]contract.Album, len(albums))
	for i, a := range albums {
		albumSlice[i] = *a
	}

	return &contract.AlbumsResponse{
		Success: true,
		Data:    albumSlice,
	}, nil
}

// CreateAlbum creates a new album
func (s *AlbumService) CreateAlbum(ctx context.Context, req *contract.CreateAlbumRequest) (*contract.AlbumResponse, error) {
	a, err := s.repository.Create(ctx, req)
	if err != nil {
		return &contract.AlbumResponse{
			Success: false,
			Message: "Failed to create album",
		}, err
	}

	return &contract.AlbumResponse{
		Success: true,
		Message: "Album created successfully",
		Data:    a,
	}, nil
}

// GetAlbumByID returns an album by its ID
func (s *AlbumService) GetAlbumByID(ctx context.Context, id int) (*contract.AlbumResponse, error) {
	a, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return &contract.AlbumResponse{
			Success: false,
			Message: "Failed to get album",
		}, err
	}

	return &contract.AlbumResponse{
		Success: true,
		Data:    a,
	}, nil
}

// UpdateAlbum updates an existing album
func (s *AlbumService) UpdateAlbum(ctx context.Context, id int, req *contract.UpdateAlbumRequest) (*contract.AlbumResponse, error) {
	a, err := s.repository.Update(ctx, id, req)
	if err != nil {
		return &contract.AlbumResponse{
			Success: false,
			Message: "Failed to update album",
		}, err
	}

	return &contract.AlbumResponse{
		Success: true,
		Message: "Album updated successfully",
		Data:    a,
	}, nil
}

// DeleteAlbum deletes an album by its ID
func (s *AlbumService) DeleteAlbum(ctx context.Context, id int) (*contract.AlbumResponse, error) {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return &contract.AlbumResponse{
			Success: false,
			Message: "Failed to delete album",
		}, err
	}

	return &contract.AlbumResponse{
		Success: true,
		Message: "Album deleted successfully",
	}, nil
}

// Add your album-related methods here that use the Ent client
// For example:
// func (s *AlbumService) CreateAlbum(ctx context.Context, album *ent.Album) (*ent.Album, error) {
//     return s.client.Album.Create()...
// }
