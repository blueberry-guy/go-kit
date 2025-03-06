package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// DeleteObject deletes a single object from S3
func (c *S3Client) DeleteObject(ctx context.Context, bucket, key string) error {
	_, err := c.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}
	return nil
}

// DeleteObjects deletes multiple objects from S3
func (c *S3Client) DeleteObjects(ctx context.Context, bucket string, keys []string) error {
	var objects []types.ObjectIdentifier
	for _, key := range keys {
		objects = append(objects, types.ObjectIdentifier{Key: &key})
	}
	quiet := true
	_, err := c.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: &bucket,
		Delete: &types.Delete{
			Objects: objects,
			Quiet:   &quiet,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to delete objects: %w", err)
	}
	return nil
}

// DeleteObjectsWithPrefix deletes all objects with the given prefix
func (c *S3Client) DeleteObjectsWithPrefix(ctx context.Context, bucket, prefix string) error {
	keys, err := c.ListObjects(ctx, bucket, prefix)
	if err != nil {
		return fmt.Errorf("failed to list objects for deletion: %w", err)
	}

	if len(keys) == 0 {
		return nil
	}

	return c.DeleteObjects(ctx, bucket, keys)
}
