package albummodel

// Album represents an album entity
type Album struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}
