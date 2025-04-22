package contract

// Album represents the album data structure
type Album struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

// CreateAlbumRequest represents the request to create a new album
type CreateAlbumRequest struct {
	URL     string `json:"url" binding:"required"`
	Caption string `json:"caption" binding:"required"`
}

// UpdateAlbumRequest represents the request to update an existing album
type UpdateAlbumRequest struct {
	URL     string `json:"url" binding:"required"`
	Caption string `json:"caption" binding:"required"`
}

// AlbumResponse represents the response for album operations
type AlbumResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    *Album `json:"data,omitempty"`
}

// AlbumsResponse represents the response for multiple albums
type AlbumsResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    []Album `json:"data,omitempty"`
}
