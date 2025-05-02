package uploadservice

import (
	"context"
	"io"
	"mime/multipart"
)

// Service defines the interface for upload-related operations
type Service interface {
	// UploadFile uploads a file from a multipart form and returns the unique filename
	UploadFile(ctx context.Context, file *multipart.FileHeader, path string) (string, error)

	// UploadFileFromReader uploads a file from an io.Reader and returns the unique filename
	UploadFileFromReader(ctx context.Context, reader io.Reader, filename string, path string) (string, error)

	// UploadFileToPublicBucket uploads a file to the public bucket and returns the unique filename
	UploadFileToPublicBucket(ctx context.Context, file *multipart.FileHeader, path string) (string, error)

	// GetPresignedURL generates a pre-signed URL for a given file key
	GetPresignedURL(ctx context.Context, fileKey string) (string, error)
}
