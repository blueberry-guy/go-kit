package s3

import (
	"context"
	"io"
)

// Writer defines methods for writing objects to S3
type Writer interface {
	// PutObject uploads data to S3
	PutObject(ctx context.Context, bucket, key string, data []byte) error

	// PutObjectStream uploads data from a reader to S3
	PutObjectStream(ctx context.Context, bucket, key string, reader io.Reader) error

	// PutObjectWithMetadata uploads data with custom metadata
	PutObjectWithMetadata(ctx context.Context, bucket, key string, data []byte, metadata map[string]string) error

	// CopyObject copies an object within S3 (same or different bucket)
	CopyObject(ctx context.Context, sourceBucket, sourceKey, destBucket, destKey string) error
}
