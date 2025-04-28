package fronthandler

import (
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// FrontHandler handles serving the frontend static files
type FrontHandler struct {
	frontendPath string
}

// NewFrontHandler creates a new front handler
func NewFrontHandler(frontendPath string) *FrontHandler {
	return &FrontHandler{
		frontendPath: frontendPath,
	}
}

// ServeFrontend serves the frontend static files
func (h *FrontHandler) ServeFrontend(c echo.Context) error {
	// Get the requested path from the URL
	requestPath := c.Param("*")
	if requestPath == "" {
		requestPath = "index.html"
	}

	// Construct the full path to the requested file
	filePath := filepath.Join(h.frontendPath, requestPath)

	// Serve the file
	return c.File(filePath)
}
