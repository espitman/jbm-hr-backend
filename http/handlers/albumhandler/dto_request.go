package albumhandler

// CreateAlbumRequest represents the request structure for creating a new album
type CreateAlbumRequest struct {
	URL     string `json:"url" binding:"required" example:"https://example.com/image.jpg"`
	Caption string `json:"caption" binding:"required" example:"Album caption"`
}

// UpdateAlbumRequest represents the request structure for updating an existing album
type UpdateAlbumRequest struct {
	URL     string `json:"url" binding:"required" example:"https://example.com/updated-image.jpg"`
	Caption string `json:"caption" binding:"required" example:"Updated album caption"`
}
