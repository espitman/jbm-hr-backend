package upload

import (
	"net/http"
	"path/filepath"

	"github.com/espitman/jbm-hr-backend/service/uploadservice"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	uploadService *uploadservice.UploadService
}

func NewUploadHandler(uploadService *uploadservice.UploadService) *UploadHandler {
	return &UploadHandler{
		uploadService: uploadService,
	}
}

// UploadImage godoc
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
