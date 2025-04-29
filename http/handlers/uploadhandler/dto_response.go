package uploadhandler

import (
	"github.com/espitman/jbm-hr-backend/http/dto"
)

// UploadImageResponse represents the response structure for image upload
type UploadImageResponse struct {
	dto.Response
	Data UploadImageData `json:"data,omitempty"`
}

// UploadImageData represents the data structure for image upload responses
type UploadImageData struct {
	Key string `json:"key"`
}

// UploadDocumentResponse represents the response structure for document upload
type UploadDocumentResponse struct {
	dto.Response
	Data UploadDocumentData `json:"data,omitempty"`
}

// UploadDocumentData represents the data structure for document upload responses
type UploadDocumentData struct {
	Key string `json:"key"`
}

// PresignedURLResponse represents the response structure for pre-signed URL
type PresignedURLResponse struct {
	dto.Response
	Data PresignedURLData `json:"data,omitempty"`
}

// PresignedURLData represents the data structure for pre-signed URL responses
type PresignedURLData struct {
	URL string `json:"url"`
}
