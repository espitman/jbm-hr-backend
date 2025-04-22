package albummodel

// Album represents an album in our system
type Album struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}
