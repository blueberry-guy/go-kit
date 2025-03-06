package s3

import (
	"context"
	"io"
)

// Reader defines methods for reading objects from S3
type Reader interface {
	// GetObject downloads an object from S3 and returns its content as a byte slice
	GetObject(ctx context.Context, bucket, key string) ([]byte, error)

	// GetObjectStream downloads an object from S3 and returns it as an io.ReadCloser
	GetObjectStream(ctx context.Context, bucket, key string) (io.ReadCloser, error)

	// ListObjects lists objects in a bucket with the given prefix
	ListObjects(ctx context.Context, bucket, prefix string) ([]string, error)

	// ObjectExists checks if an object exists in the bucket
	ObjectExists(ctx context.Context, bucket, key string) (bool, error)

	// GetObjectMetadata retrieves metadata of an object without downloading its contents
	GetObjectMetadata(ctx context.Context, bucket, key string) (map[string]string, error)
}
