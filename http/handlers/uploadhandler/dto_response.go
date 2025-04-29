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
