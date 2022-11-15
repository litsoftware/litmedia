package grpc

import (
	"google.golang.org/grpc"
	"log"
)

var (
	conn *grpc.ClientConn
	err  error
)

func initClient(address string) error {
	conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}

	return nil
}

func Close() {
	if conn != nil {
		_ = conn.Close()
	}
}

// NewGrpcClient 创建一个到指定服务器的 grpc 连接
func NewGrpcClient(address string) (*grpc.ClientConn, error) {
	if err := initClient(address); err != nil {
		return nil, err
	}
	return conn, nil
}
