package core

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

type Gateway interface {
	Name() string
	NewGatewayLayer(creds Credentials) (ObjectLayer, error)
}

type ObjectLayer interface {
	PutObject(ctx context.Context, objectKey string, reader io.Reader, size int64, options minio.PutObjectOptions) (minio.UploadInfo, error)
	GetObject(ctx context.Context, objectKey string, options minio.GetObjectOptions) (interface{}, error)
	CopyObject(ctx context.Context, dest minio.CopyDestOptions, src minio.CopySrcOptions) (minio.UploadInfo, error)
}
