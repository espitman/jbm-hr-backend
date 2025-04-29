package uploadservice

import (
	"context"
	"io"
	"mime/multipart"
)

// Service defines the interface for upload-related operations
type Service interface {
	// UploadFile uploads a file from a multipart form and returns the unique filename
	UploadFile(ctx context.Context, file *multipart.FileHeader) (string, error)

	// UploadFileFromReader uploads a file from an io.Reader and returns the unique filename
	UploadFileFromReader(ctx context.Context, reader io.Reader, filename string) (string, error)
}
