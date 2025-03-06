package s3

import (
	"context"

	"github.com/blueberry-guy/go-kit/pkg/s3/internal"
)

// Client combines all S3 operations
type Client interface {
	Reader
	Writer
	Deleter
	Presigner
}

// NewClient creates a new S3 client with all operations
func NewClient(ctx context.Context, region string) (Client, error) {
	return internal.NewS3Client(ctx, region)
}

// NewReader creates a new S3 reader
func NewReader(ctx context.Context, region string) (Reader, error) {
	return internal.NewS3Client(ctx, region)
}

// NewWriter creates a new S3 writer
func NewWriter(ctx context.Context, region string) (Writer, error) {
	return internal.NewS3Client(ctx, region)
}

// NewDeleter creates a new S3 deleter
func NewDeleter(ctx context.Context, region string) (Deleter, error) {
	return internal.NewS3Client(ctx, region)
}

// NewPresigner creates a new S3 presigner
func NewPresigner(ctx context.Context, region string) (Presigner, error) {
	return internal.NewS3Client(ctx, region)
}
