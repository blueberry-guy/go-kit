package s3

import "context"

// Deleter defines methods for deleting objects from S3
type Deleter interface {
	// DeleteObject deletes a single object from S3
	DeleteObject(ctx context.Context, bucket, key string) error

	// DeleteObjects deletes multiple objects from S3
	DeleteObjects(ctx context.Context, bucket string, keys []string) error

	// DeleteObjectsWithPrefix deletes all objects with the given prefix
	DeleteObjectsWithPrefix(ctx context.Context, bucket, prefix string) error
}
