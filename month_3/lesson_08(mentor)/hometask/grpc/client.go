package grpc_client

import (
	"fmt"

	"market/config"
	person "market/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	PersonService() person.PersonServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connPerson, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.PersonServiceHost, cfg.PersonServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("courier service dial host: %s port:%s err: %s",
			cfg.PersonServiceHost, cfg.PersonServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"person_service": person.NewPersonServiceClient(connPerson),
		},
	}, nil
}

func (g *GrpcClient) PersonService() person.PersonServiceClient {
	return g.connections["courier_service"].(person.PersonServiceClient)
}
