package models

type S3BucketSummary struct {
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
	ObjectCount int    `json:"objectCount"`
	SizeBytes   int64  `json:"sizeBytes"`
	SizeLabel   string `json:"sizeLabel"`
}

type S3ObjectSummary struct {
	Key          string `json:"key"`
	SizeBytes    int64  `json:"sizeBytes"`
	SizeLabel    string `json:"sizeLabel"`
	LastModified string `json:"lastModified"`
	ETag         string `json:"etag"`
	StorageClass string `json:"storageClass"`
}

type S3ObjectMoveRequest struct {
	SourceBucket string `json:"sourceBucket" validate:"required"`
	SourceKey    string `json:"sourceKey" validate:"required"`
	TargetBucket string `json:"targetBucket" validate:"required"`
	TargetKey    string `json:"targetKey" validate:"required"`
	Mode         string `json:"mode"`
}
