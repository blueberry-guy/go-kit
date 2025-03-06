package s3

import (
	"context"
	"time"
)

// Presigner defines methods for generating pre-signed URLs for S3 operations
type Presigner interface {
	// GetPresignedURL generates a pre-signed URL for downloading an object
	GetPresignedURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error)

	// PutPresignedURL generates a pre-signed URL for uploading an object
	PutPresignedURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error)

	// DeletePresignedURL generates a pre-signed URL for deleting an object
	DeletePresignedURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error)
}
