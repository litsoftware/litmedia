package s3

import (
	"context"
	"github.com/litsoftware/litmedia/internal/gateway/core"
	"github.com/litsoftware/litmedia/internal/gateway/errors"
	"github.com/litsoftware/litmedia/internal/gateway/unssupported"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"strings"
)

type GatewayS3Layer struct {
	unssupported.GatewayUnsupported
	client *minio.Client
	creds  core.Credentials
}

func (g *GatewayS3Layer) NewGatewayLayer(creds core.Credentials) (core.ObjectLayer, error) {
	var err error

	if strings.HasPrefix(creds.Endpoint, "http://") {
		creds.Ssl = false
	}
	if strings.HasPrefix(creds.Endpoint, "https://") {
		creds.Ssl = true
	}

	creds.Endpoint = strings.TrimPrefix(creds.Endpoint, "https://")
	creds.Endpoint = strings.TrimPrefix(creds.Endpoint, "http://")

	g.creds = creds
	minioClient, err := minio.New(creds.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(creds.AccessKeyId, creds.AccessKeySecret, ""),
		Secure: creds.Ssl,
	})
	if err != nil {
		log.Fatalf("GatewayS3Layer error: %s", err)
	}

	g.client = minioClient
	return g, err
}

func (g *GatewayS3Layer) Name() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) PutObject(ctx context.Context, objectKey string, reader io.Reader, size int64, options minio.PutObjectOptions) (minio.UploadInfo, error) {
	return g.client.PutObject(ctx, g.creds.Bucket, objectKey, reader, size, options)
}

func (g *GatewayS3Layer) GetObject(ctx context.Context, objectKey string, options minio.GetObjectOptions) (interface{}, error) {
	return g.client.GetObject(ctx, g.creds.Bucket, objectKey, options)
}

func (g *GatewayS3Layer) CopyObject(ctx context.Context, dest minio.CopyDestOptions, src minio.CopySrcOptions) (minio.UploadInfo, error) {
	return g.client.CopyObject(ctx, dest, src)
}

func (g *GatewayS3Layer) AppendObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) DeleteObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) DeleteMultipleObjects() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) HeadObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) GetObjectMeta() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) RestoreObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) PostObject() error {
	return &errors.NotImplemented{}
}

func (g *GatewayS3Layer) SelectObject() error {
	return &errors.NotImplemented{}
}
