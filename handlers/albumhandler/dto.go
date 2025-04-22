package albumhandler

import (
	"github.com/espitman/jbm-hr-backend/contract"
)

// AlbumResponse represents the response structure for album operations
type AlbumResponse struct {
	Success bool            `json:"success" example:"true"`
	Message string          `json:"message,omitempty" example:"Album created successfully"`
	Data    *contract.Album `json:"data,omitempty"`
}

// AlbumsResponse represents the response structure for multiple albums
type AlbumsResponse struct {
	Success bool              `json:"success" example:"true"`
	Message string            `json:"message,omitempty" example:"Albums retrieved successfully"`
	Data    []*contract.Album `json:"data,omitempty"`
}
