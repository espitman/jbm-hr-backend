package uihandler

import (
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// UIHandler handles serving the UI static files
type UIHandler struct {
	webPath   string
	adminPath string
}

// NewUIHandler creates a new UI handler
func NewUIHandler(webPath, adminPath string) *UIHandler {
	return &UIHandler{
		webPath:   webPath,
		adminPath: adminPath,
	}
}

// ServeUI serves the web UI static files
func (h *UIHandler) ServeUI(c echo.Context) error {
	// Get the requested path from the URL
	requestPath := c.Param("*")
	if requestPath == "" {
		requestPath = "index.html"
	}

	// Construct the full path to the requested file
	filePath := filepath.Join(h.webPath, requestPath)

	// Serve the file
	return c.File(filePath)
}

// ServeAdminUI serves the admin UI static files
func (h *UIHandler) ServeAdminUI(c echo.Context) error {
	// Get the requested path from the URL
	requestPath := c.Param("*")

	// If the path is empty or ends with a slash, serve index.html
	if requestPath == "" || requestPath[len(requestPath)-1] == '/' {
		requestPath = "index.html"
	}

	// Construct the full path to the requested file
	filePath := filepath.Join(h.adminPath, requestPath)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If the file doesn't exist, serve index.html
		filePath = filepath.Join(h.adminPath, "index.html")
	}

	// Serve the file
	return c.File(filePath)
}
