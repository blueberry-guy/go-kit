package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// PutObject uploads data to S3
func (c *S3Client) PutObject(ctx context.Context, bucket, key string, data []byte) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return fmt.Errorf("failed to put object: %w", err)
	}
	return nil
}

// PutObjectStream uploads data from a reader to S3
func (c *S3Client) PutObjectStream(ctx context.Context, bucket, key string, reader io.Reader) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   reader,
	})
	if err != nil {
		return fmt.Errorf("failed to put object stream: %w", err)
	}
	return nil
}

// PutObjectWithMetadata uploads data with custom metadata
func (c *S3Client) PutObjectWithMetadata(ctx context.Context, bucket, key string, data []byte, metadata map[string]string) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:   &bucket,
		Key:      &key,
		Body:     bytes.NewReader(data),
		Metadata: metadata,
	})
	if err != nil {
		return fmt.Errorf("failed to put object with metadata: %w", err)
	}
	return nil
}

// CopyObject copies an object within S3
func (c *S3Client) CopyObject(ctx context.Context, sourceBucket, sourceKey, destBucket, destKey string) error {
	source := fmt.Sprintf("%s/%s", sourceBucket, sourceKey)
	_, err := c.client.CopyObject(ctx, &s3.CopyObjectInput{
		CopySource: &source,
		Bucket:     &destBucket,
		Key:        &destKey,
	})
	if err != nil {
		return fmt.Errorf("failed to copy object: %w", err)
	}
	return nil
}
