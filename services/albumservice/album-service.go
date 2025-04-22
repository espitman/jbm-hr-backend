package albumservice

import (
	"gin-project/models/albummodel"
)

// AlbumService handles album-related business logic
type AlbumService struct {
	albums []albummodel.Album
}

// New creates a new AlbumService instance with mock data
func New() *AlbumService {
	// Mock data
	albums := []albummodel.Album{
		{ID: 1, URL: "https://example.com/album1.jpg", Caption: "Beautiful sunset"},
		{ID: 2, URL: "https://example.com/album2.jpg", Caption: "Mountain view"},
		{ID: 3, URL: "https://example.com/album3.jpg", Caption: "Beach vacation"},
	}

	return &AlbumService{
		albums: albums,
	}
}

// GetAllAlbums returns all albums
func (s *AlbumService) GetAllAlbums() []albummodel.Album {
	return s.albums
}
