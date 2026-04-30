package repositories

import (
	"context"
	"io"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type StorageRepository interface {
	ListBuckets(ctx context.Context) (*s3.ListBucketsOutput, error)
	ListObjects(ctx context.Context, bucket string, prefix string) ([]types.Object, error)
	CreateBucket(ctx context.Context, bucket string) error
	DeleteBucket(ctx context.Context, bucket string) error
	PutObject(ctx context.Context, bucket string, key string, contentType string, body io.Reader) error
	CopyObject(ctx context.Context, sourceBucket string, sourceKey string, targetBucket string, targetKey string) error
	DeleteObject(ctx context.Context, bucket string, key string) error
	GetObject(ctx context.Context, bucket string, key string) (*s3.GetObjectOutput, error)
}

type storageRepository struct {
	objectStorage *s3.Client
}

func NewStorageRepository(objectStorage *s3.Client) StorageRepository {
	return &storageRepository{
		objectStorage: objectStorage,
	}
}

func (sr *storageRepository) ListBuckets(ctx context.Context) (*s3.ListBucketsOutput, error) {
	return sr.objectStorage.ListBuckets(ctx, &s3.ListBucketsInput{})
}

func (sr *storageRepository) ListObjects(ctx context.Context, bucket string, prefix string) ([]types.Object, error) {
	paginator := s3.NewListObjectsV2Paginator(sr.objectStorage, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})

	objects := make([]types.Object, 0)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		objects = append(objects, page.Contents...)
	}

	return objects, nil
}

func (sr *storageRepository) CreateBucket(ctx context.Context, bucket string) error {
	_, err := sr.objectStorage.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	return err
}

func (sr *storageRepository) DeleteBucket(ctx context.Context, bucket string) error {
	_, err := sr.objectStorage.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	return err
}

func (sr *storageRepository) PutObject(ctx context.Context, bucket string, key string, contentType string, body io.Reader) error {
	_, err := sr.objectStorage.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	return err
}

func (sr *storageRepository) CopyObject(ctx context.Context, sourceBucket string, sourceKey string, targetBucket string, targetKey string) error {
	_, err := sr.objectStorage.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(targetBucket),
		CopySource: aws.String(copySource(sourceBucket, sourceKey)),
		Key:        aws.String(targetKey),
		ACL:        types.ObjectCannedACLPublicRead,
	})
	return err
}

func (sr *storageRepository) DeleteObject(ctx context.Context, bucket string, key string) error {
	_, err := sr.objectStorage.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err
}

func (sr *storageRepository) GetObject(ctx context.Context, bucket string, key string) (*s3.GetObjectOutput, error) {
	return sr.objectStorage.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
}

func copySource(bucket string, key string) string {
	parts := strings.Split(key, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	return "/" + url.PathEscape(bucket) + "/" + strings.Join(parts, "/")
}
