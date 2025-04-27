package albumservice

import (
	"context"

	"github.com/espitman/jbm-hr-backend/contract"
)

// Service defines the interface for album-related operations
type Service interface {
	// GetAllAlbums returns all albums
	GetAllAlbums(ctx context.Context) ([]*contract.Album, error)

	// CreateAlbum creates a new album
	CreateAlbum(ctx context.Context, input *contract.CreateAlbumInput) (*contract.Album, error)

	// GetAlbumByID returns an album by its ID
	GetAlbumByID(ctx context.Context, id int) (*contract.Album, error)

	// UpdateAlbum updates an existing album
	UpdateAlbum(ctx context.Context, id int, input *contract.UpdateAlbumInput) (*contract.Album, error)

	// DeleteAlbum deletes an album by its ID
	DeleteAlbum(ctx context.Context, id int) error
}
