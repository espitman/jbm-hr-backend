package albumhandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
	"github.com/espitman/jbm-hr-backend/handlers/dto"
)

// AlbumResponse represents the response structure for album operations
type AlbumResponse struct {
	dto.Response
	Data *contract.Album `json:"data,omitempty"`
}

// NewAlbumResponse creates a new album response
func NewAlbumResponse(album *contract.Album) *AlbumResponse {
	return &AlbumResponse{
		Response: dto.Response{
			Success: true,
			Message: "Album operation completed successfully",
		},
		Data: album,
	}
}

// AlbumsResponse represents the response structure for multiple albums
type AlbumsResponse struct {
	dto.Response
	Data []*contract.Album `json:"data,omitempty"`
}

// NewAlbumsResponse creates a new albums response
func NewAlbumsResponse(albums []*contract.Album) *AlbumsResponse {
	return &AlbumsResponse{
		Response: dto.Response{
			Success: true,
			Message: "Albums retrieved successfully",
		},
		Data: albums,
	}
}
