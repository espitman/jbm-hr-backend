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

// UploadDocument handles document file uploads
// @Summary Upload a document
// @Description Upload a document file (PDF, DOC, DOCX) to S3 storage
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Document file to upload"
// @Success 200 {object} UploadDocumentResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/upload/document [post]
func (h *UploadHandler) UploadDocument(c echo.Context) error {
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return dto.BadRequestJSON(c, "Failed to get file from request")
	}

	// Validate file type
	ext := filepath.Ext(file.Filename)
	if ext != ".pdf" && ext != ".doc" && ext != ".docx" {
		return dto.BadRequestJSON(c, "Invalid file type. Only PDF, DOC, and DOCX files are allowed")
	}

	// Define the path for document uploads
	path := "documents"

	// Upload the file
	fileKey, err := h.uploadService.UploadFile(c.Request().Context(), file, path)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "Failed to upload file")
	}

	// Return the key of the uploaded file
	return dto.SuccessJSON(c, UploadDocumentData{
		Key: fileKey,
	})
}

// GetPresignedURL generates a pre-signed URL for a given file key
// @Summary Get a pre-signed URL for a file
// @Description Generate a pre-signed URL for accessing a file in S3 storage
// @Tags upload
// @Accept json
// @Produce json
// @Param key path string true "File key"
// @Success 200 {object} PresignedURLResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/upload/presigned/{key} [get]
func (h *UploadHandler) GetPresignedURL(c echo.Context) error {
	// Get the file key from the path parameter
	fileKey := c.Param("key")
	if fileKey == "" {
		return dto.BadRequestJSON(c, "File key is required")
	}

	// Generate the pre-signed URL
	presignedURL, err := h.uploadService.GetPresignedURL(c.Request().Context(), fileKey)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "Failed to generate pre-signed URL")
	}

	// Return the pre-signed URL
	return dto.SuccessJSON(c, PresignedURLData{
		URL: presignedURL,
	})
}

// UploadPublicImage handles public image file uploads
// @Summary Upload a public image
// @Description Upload an image file to public S3 storage
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file to upload"
// @Success 200 {object} UploadPublicImageResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/upload/image/public [post]
func (h *UploadHandler) UploadPublicImage(c echo.Context) error {
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

	// Define the path for public image uploads
	path := "public/images"

	// Upload the file to public bucket
	fileURL, err := h.uploadService.UploadFileToPublicBucket(c.Request().Context(), file, path)
	if err != nil {
		return dto.InternalServerErrorJSON(c, "Failed to upload file")
	}

	// Return the URL of the uploaded file
	return dto.SuccessJSON(c, UploadPublicImageData{
		URL: fileURL,
	})
}
