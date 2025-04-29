package uihandler

import (
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// UIHandler handles serving the UI static files
type UIHandler struct {
	uiPath string
}

// NewUIHandler creates a new UI handler
func NewUIHandler(uiPath string) *UIHandler {
	return &UIHandler{
		uiPath: uiPath,
	}
}

// ServeUI serves the UI static files
func (h *UIHandler) ServeUI(c echo.Context) error {
	// Get the requested path from the URL
	requestPath := c.Param("*")
	if requestPath == "" {
		requestPath = "index.html"
	}

	// Construct the full path to the requested file
	filePath := filepath.Join(h.uiPath, requestPath)

	// Serve the file
	return c.File(filePath)
}
