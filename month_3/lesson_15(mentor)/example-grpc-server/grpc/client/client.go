package grpc_client

import (
	"example-grpc-server/config"
	sale_service "example-grpc-server/genproto"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	StreamService() sale_service.StreamServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connSream, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.PersonServiceHost, cfg.PersonServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("courier service dial host: %s port:%s err: %s",
			cfg.PersonServiceHost, cfg.PersonServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"stream_service": sale_service.NewStreamServiceClient(connSream),
		},
	}, nil
}

func (g *GrpcClient) StreamService() sale_service.StreamServiceClient {
	return g.connections["stream_service"].(sale_service.StreamServiceClient)
}
