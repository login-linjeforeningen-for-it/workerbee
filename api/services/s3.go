package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"sort"
	"strings"
	"workerbee/models"
	"workerbee/repositories"
)

type StorageService struct {
	repo repositories.StorageRepository
}

func NewStorageService(repo repositories.StorageRepository) *StorageService {
	return &StorageService{repo: repo}
}

func (ss *StorageService) ListBuckets(ctx context.Context) ([]models.S3BucketSummary, error) {
	output, err := ss.repo.ListBuckets(ctx)
	if err != nil {
		return nil, err
	}

	buckets := make([]models.S3BucketSummary, 0, len(output.Buckets))
	for _, bucket := range output.Buckets {
		name := bucket.Name
		if name == nil || *name == "" {
			continue
		}

		objects, err := ss.repo.ListObjects(ctx, *name, "")
		if err != nil {
			return nil, err
		}

		var sizeBytes int64
		for _, object := range objects {
			sizeBytes += objectInt64(object.Size)
		}

		createdAt := ""
		if bucket.CreationDate != nil {
			createdAt = bucket.CreationDate.Format("2006-01-02T15:04:05Z07:00")
		}

		buckets = append(buckets, models.S3BucketSummary{
			Name:        *name,
			CreatedAt:   createdAt,
			ObjectCount: len(objects),
			SizeBytes:   sizeBytes,
			SizeLabel:   formatBytes(sizeBytes),
		})
	}

	sort.SliceStable(buckets, func(i, j int) bool {
		if buckets[i].SizeBytes == buckets[j].SizeBytes {
			return buckets[i].Name < buckets[j].Name
		}
		return buckets[i].SizeBytes > buckets[j].SizeBytes
	})

	return buckets, nil
}

func (ss *StorageService) ListObjects(ctx context.Context, bucket string, prefix string) ([]models.S3ObjectSummary, error) {
	objects, err := ss.repo.ListObjects(ctx, bucket, prefix)
	if err != nil {
		return nil, err
	}

	summaries := make([]models.S3ObjectSummary, 0, len(objects))
	for _, object := range objects {
		if object.Key == nil || strings.HasSuffix(*object.Key, "/") {
			continue
		}

		lastModified := ""
		if object.LastModified != nil {
			lastModified = object.LastModified.Format("2006-01-02T15:04:05Z07:00")
		}

		sizeBytes := objectInt64(object.Size)
		summaries = append(summaries, models.S3ObjectSummary{
			Key:          *object.Key,
			SizeBytes:    sizeBytes,
			SizeLabel:    formatBytes(sizeBytes),
			LastModified: lastModified,
			ETag:         strings.Trim(objectString(object.ETag), "\""),
			StorageClass: string(object.StorageClass),
		})
	}

	sort.SliceStable(summaries, func(i, j int) bool {
		return summaries[i].Key < summaries[j].Key
	})

	return summaries, nil
}

func (ss *StorageService) CreateBucket(ctx context.Context, bucket string) error {
	return ss.repo.CreateBucket(ctx, bucket)
}

func (ss *StorageService) DeleteBucket(ctx context.Context, bucket string) error {
	return ss.repo.DeleteBucket(ctx, bucket)
}

func (ss *StorageService) PutObject(ctx context.Context, bucket string, key string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return ss.repo.PutObject(ctx, bucket, key, contentType, src)
}

func (ss *StorageService) MoveObject(ctx context.Context, body models.S3ObjectMoveRequest) error {
	mode := body.Mode
	if mode == "" {
		mode = "move"
	}

	if err := ss.repo.CopyObject(ctx, body.SourceBucket, body.SourceKey, body.TargetBucket, body.TargetKey); err != nil {
		return err
	}

	if mode != "copy" {
		return ss.repo.DeleteObject(ctx, body.SourceBucket, body.SourceKey)
	}

	return nil
}

func (ss *StorageService) DeleteObject(ctx context.Context, bucket string, key string) error {
	return ss.repo.DeleteObject(ctx, bucket, key)
}

func (ss *StorageService) GetObject(ctx context.Context, bucket string, key string) (io.ReadCloser, string, int64, error) {
	output, err := ss.repo.GetObject(ctx, bucket, key)
	if err != nil {
		return nil, "", 0, err
	}

	return output.Body, objectString(output.ContentType), objectInt64(output.ContentLength), nil
}

func objectString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func objectInt64(value *int64) int64 {
	if value == nil {
		return 0
	}
	return *value
}

func formatBytes(bytes int64) string {
	if bytes <= 0 {
		return "0 B"
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	value := float64(bytes)
	index := 0
	for value >= 1024 && index < len(units)-1 {
		value /= 1024
		index++
	}

	if index == 0 {
		return strings.TrimSuffix(strings.TrimSuffix(formatFloat(value), ".0"), ".") + " " + units[index]
	}
	return formatFloat(value) + " " + units[index]
}

func formatFloat(value float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.1f", value), "0"), ".")
}
