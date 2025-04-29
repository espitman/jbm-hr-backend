package uploadservice

import (
	"context"
	"io"
	"mime/multipart"
)

// Service defines the interface for upload-related operations
type Service interface {
	// UploadFile uploads a file from a multipart form
	UploadFile(ctx context.Context, file *multipart.FileHeader) (string, error)

	// UploadFileFromReader uploads a file from an io.Reader
	UploadFileFromReader(ctx context.Context, reader io.Reader, filename string) (string, error)
}
