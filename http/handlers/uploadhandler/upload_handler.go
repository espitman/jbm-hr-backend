package uploadhandler

import (
	"net/http"
	"path/filepath"

	"github.com/espitman/jbm-hr-backend/service/uploadservice"
	"github.com/labstack/echo/v4"
)

// UploadHandler handles HTTP requests for file uploads
type UploadHandler struct {
	uploadService uploadservice.Service
}

// NewUploadHandler creates a new UploadHandler
func NewUploadHandler(uploadService uploadservice.Service) *UploadHandler {
	return &UploadHandler{
		uploadService: uploadService,
	}
}

// UploadImage handles image file uploads
// @Summary Upload an image
// @Description Upload an image file to S3 storage
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file to upload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/upload/image [post]
func (h *UploadHandler) UploadImage(c echo.Context) error {
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to get file from request",
		})
	}

	// Validate file type
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid file type. Only jpg, jpeg, png, and gif files are allowed",
		})
	}

	// Upload the file
	fileURL, err := h.uploadService.UploadFile(c.Request().Context(), file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to upload file",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"url": fileURL,
	})
}
