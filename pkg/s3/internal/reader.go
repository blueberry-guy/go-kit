package internal

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetObject downloads an object from S3 and returns its content
func (c *S3Client) GetObject(ctx context.Context, bucket, key string) ([]byte, error) {
	result, err := c.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	defer result.Body.Close()

	return ioutil.ReadAll(result.Body)
}

// GetObjectStream downloads an object and returns it as a stream
func (c *S3Client) GetObjectStream(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	result, err := c.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get object stream: %w", err)
	}
	return result.Body, nil
}

// ListObjects lists objects in a bucket with the given prefix
func (c *S3Client) ListObjects(ctx context.Context, bucket, prefix string) ([]string, error) {
	var keys []string
	paginator := s3.NewListObjectsV2Paginator(c.client, &s3.ListObjectsV2Input{
		Bucket: &bucket,
		Prefix: &prefix,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %w", err)
		}

		for _, obj := range page.Contents {
			keys = append(keys, *obj.Key)
		}
	}

	return keys, nil
}

// ObjectExists checks if an object exists in the bucket
func (c *S3Client) ObjectExists(ctx context.Context, bucket, key string) (bool, error) {
	_, err := c.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return false, nil
	}
	return true, nil
}

// GetObjectMetadata retrieves metadata of an object
func (c *S3Client) GetObjectMetadata(ctx context.Context, bucket, key string) (map[string]string, error) {
	result, err := c.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get object metadata: %w", err)
	}
	return result.Metadata, nil
}
