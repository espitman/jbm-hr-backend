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

// AlbumsResponse represents the response structure for multiple albums
type AlbumsResponse struct {
	dto.Response
	Data []*contract.Album `json:"data,omitempty"`
}
