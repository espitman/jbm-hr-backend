package contract

// Album represents the album data structure
type Album struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	Caption      string `json:"caption"`
	DisplayOrder int    `json:"display_order"`
}

// CreateAlbumInput represents the request to create a new album
type CreateAlbumInput struct {
	URL          string `json:"url" binding:"required"`
	Caption      string `json:"caption" binding:"required"`
	DisplayOrder int    `json:"display_order"`
}

// UpdateAlbumInput represents the request to update an existing album
type UpdateAlbumInput struct {
	URL          string `json:"url" binding:"required"`
	Caption      string `json:"caption" binding:"required"`
	DisplayOrder int    `json:"display_order"`
}
