package unssupported

import (
	"context"
	"github.com/litsoftware/litmedia/internal/gateway/errors"
	"github.com/minio/minio-go/v7"
	"io"
)

type GatewayUnsupported struct{}

func (g *GatewayUnsupported) Name() string {
	return "Unsupported"
}

func (g *GatewayUnsupported) PutObject(ctx context.Context, objectKey string, reader io.Reader, size int64, options minio.PutObjectOptions) (minio.UploadInfo, error) {
	return minio.UploadInfo{}, &errors.NotImplemented{}
}

func (g *GatewayUnsupported) GetObject(ctx context.Context, objectKey string, options minio.GetObjectOptions) (interface{}, error) {
	return nil, &errors.NotImplemented{}
}

func (g *GatewayUnsupported) CopyObject(ctx context.Context, dest minio.CopyDestOptions, src minio.CopySrcOptions) (minio.UploadInfo, error) {
	return minio.UploadInfo{}, &errors.NotImplemented{}
}

func (g *GatewayUnsupported) AppendObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayUnsupported) DeleteObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayUnsupported) DeleteMultipleObjects() error {
	return &errors.NotImplemented{}
}

func (g *GatewayUnsupported) HeadObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayUnsupported) GetObjectMeta() error {
	return &errors.NotImplemented{}
}

func (g *GatewayUnsupported) RestoreObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayUnsupported) PostObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayUnsupported) SelectObject() error {
	return &errors.NotImplemented{}
}
