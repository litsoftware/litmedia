package oss

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/litsoftware/litmedia/internal/gateway/core"
	"github.com/litsoftware/litmedia/internal/gateway/unssupported"
	"github.com/litsoftware/litmedia/pkg/jsonh"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"net/http"
)

type OssLayer struct {
	unssupported.GatewayUnsupported
	client     *oss.Client
	bucket     *oss.Bucket
	httpClient *http.Client
}

func (g *OssLayer) NewGatewayLayer(creds core.Credentials) (core.ObjectLayer, error) {
	var err error

	g.client, err = oss.New(creds.Endpoint, creds.AccessKeyId, creds.AccessKeySecret)
	if err != nil {
		log.Fatalf("oss.New error: %s", err)
	}

	g.bucket, err = g.client.Bucket(creds.Bucket)
	if err != nil {
		log.Fatalf("oss.Bucket error: %s", err)
	}

	return g, err
}

func (g *OssLayer) Name() string {
	return "oss"
}

func (g *OssLayer) convertOpts(options ...core.Option) []oss.Option {
	var OssOpts []oss.Option
	if len(options) > 0 {
		jsonh.ConvertTo(options, &OssOpts)
	}
	return OssOpts
}

func (g *OssLayer) PutObject(ctx context.Context, objectKey string, reader io.Reader, size int64, options minio.PutObjectOptions) (minio.UploadInfo, error) {
	fmt.Println("unused param ", ctx, size, options)
	opts := g.convertOpts(nil)
	err := g.bucket.PutObject(objectKey, reader, opts...)
	if err != nil {
		return minio.UploadInfo{}, err
	}

	return minio.UploadInfo{}, err
}

func (g *OssLayer) GetObject(ctx context.Context, objectKey string, options minio.GetObjectOptions) (interface{}, error) {
	fmt.Println("unused param ", ctx, options)
	opts := g.convertOpts(nil)
	return g.bucket.GetObject(objectKey, opts...)
}

func (g *OssLayer) CopyObject(ctx context.Context, dest minio.CopyDestOptions, src minio.CopySrcOptions) (minio.UploadInfo, error) {
	OssOpts := g.convertOpts(nil)
	res, err := g.bucket.CopyObject(src.Object, dest.Object, OssOpts...)
	return minio.UploadInfo{
		ETag:         res.ETag,
		LastModified: res.LastModified,
	}, err
}

func (g *OssLayer) AppendObject() error {
	return nil
}

func (g *OssLayer) DeleteObject() error {
	return nil
}

func (g *OssLayer) DeleteMultipleObjects() error {
	return nil
}

func (g *OssLayer) HeadObject() error {
	return nil
}

func (g *OssLayer) GetObjectMeta() error {
	return nil
}

func (g *OssLayer) RestoreObject() error {
	return nil
}

func (g *OssLayer) PostObject() error {
	return nil
}

func (g *OssLayer) SelectObject() error {
	return nil
}
