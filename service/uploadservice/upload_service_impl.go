package uploadservice

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/espitman/jbm-hr-backend/utils/config"
	"github.com/google/uuid"
)

type service struct {
	s3Client *s3.Client
	bucket   string
	endpoint string
}

// New creates a new UploadService instance
func New() (Service, error) {
	// Get configuration from environment variables using config utility
	endpoint := config.GetConfig("S3_ENDPOINT", "https://storage.c2.liara.space")
	accessKey := config.GetConfig("S3_ACCESS_KEY", "icvivvt5uv1g2l7s")
	secretKey := config.GetConfig("S3_SECRET_KEY", "2049dea6-0bf7-4204-8109-bd474c5df834")
	bucket := config.GetConfig("S3_BUCKET", "your-bucket-name")
	region := config.GetConfig("S3_REGION", "us-east-1")

	// Create custom endpoint resolver
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: endpoint,
		}, nil
	})

	// Load AWS configuration
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithEndpointResolverWithOptions(customResolver),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			"",
		)),
		awsconfig.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)

	return &service{
		s3Client: client,
		bucket:   bucket,
		endpoint: endpoint,
	}, nil
}

func (s *service) UploadFile(ctx context.Context, file *multipart.FileHeader, path string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Generate a unique filename using UUID
	ext := filepath.Ext(file.Filename)
	uniqueFilename := uuid.New().String() + ext

	// Combine path and filename
	fullPath := filepath.Join(path, uniqueFilename)

	// Upload the file to S3
	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fullPath),
		Body:   src,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	return fullPath, nil
}

func (s *service) UploadFileFromReader(ctx context.Context, reader io.Reader, filename string, path string) (string, error) {
	// Generate a unique filename using UUID
	ext := filepath.Ext(filename)
	uniqueFilename := uuid.New().String() + ext

	// Combine path and filename
	fullPath := filepath.Join(path, uniqueFilename)

	// Upload the file to S3
	_, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fullPath),
		Body:   reader,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	return fullPath, nil
}

// GetPresignedURL generates a pre-signed URL for a given file key
func (s *service) GetPresignedURL(ctx context.Context, fileKey string) (string, error) {
	// Create a pre-signed URL request
	presignedClient := s3.NewPresignClient(s.s3Client)

	// Create the request
	request, err := presignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileKey),
	}, s3.WithPresignExpires(time.Hour*1))

	if err != nil {
		return "", fmt.Errorf("failed to generate pre-signed URL: %v", err)
	}

	return request.URL, nil
}
