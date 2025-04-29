package uploadhandler

import (
	"path/filepath"

	"github.com/espitman/jbm-hr-backend/http/dto"
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
// @Success 200 {object} UploadImageResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/upload/image [post]
func (h *UploadHandler) UploadImage(c echo.Context) error {
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return dto.BadRequestJSON(c, "Failed to get file from request")
	}

	// Validate file type
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return dto.BadRequestJSON(c, "Invalid file type. Only jpg, jpeg, png, and gif files are allowed")
	}

	// Define the path for image uploads
	path := "images"

	// Upload the file
	fileKey, err := h.uploadService.UploadFile(c.Request().Context(), file, path)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "Failed to upload file")
	}

	// Return the key of the uploaded file
	return dto.SuccessJSON(c, UploadImageData{
		Key: fileKey,
	})
}
