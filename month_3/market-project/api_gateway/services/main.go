package services

import (
	"api_gateway/config"
	branch_service "api_gateway/genproto"

	"fmt"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	BranchService() branch_service.BranchServiceClient
	BranchProductService() branch_service.BranchProductServiceClient

	// AttributeService() position_service.AttributeServiceClient
	// CompanyService() company_service.CompanyServiceClient
	// PositionService() position_service.PositionServiceClient
}

type grpcClients struct {
	branchService        branch_service.BranchServiceClient
	branchProductService branch_service.BranchProductServiceClient

	// attributeService  position_service.AttributeServiceClient
	// companyService    company_service.CompanyServiceClient
	// positionService   position_service.PositionServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connBranchService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BranchServiceHost, conf.BranchServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connBranchProductService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BranchServiceHost, conf.BranchServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		branchService:        branch_service.NewBranchServiceClient(connBranchService),
		branchProductService: branch_service.NewBranchProductServiceClient(connBranchProductService),
	}, nil
}

func (g *grpcClients) BranchService() branch_service.BranchServiceClient {
	return g.branchService
}

func (g *grpcClients) BranchProductService() branch_service.BranchProductServiceClient {
	return g.branchProductService
}
