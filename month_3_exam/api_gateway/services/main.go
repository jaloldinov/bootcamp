package services

import (
	"api_gateway/config"
	catalog_service "api_gateway/genproto/catalog_service"
	order_service "api_gateway/genproto/order_service"
	user_service "api_gateway/genproto/user_service"

	"fmt"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	// 	CATALOG-SERVICE
	Category() catalog_service.CategoryServiceClient
	Product() catalog_service.ProductServiceClient

	// 	ORDER-SERVICE
	Order() order_service.OrderServiceClient
	DeliveryTariff() order_service.DeliveryTariffServiceClient

	// 	USER-SERVICE
	User() user_service.UserServiceClient
	Client() user_service.ClientServiceClient
	Courier() user_service.CourierServiceClient
	Branch() user_service.BranchServiceClient
}

type grpcClients struct {
	// catalog service
	categoryService catalog_service.CategoryServiceClient
	productService  catalog_service.ProductServiceClient

	// order service
	orderService          order_service.OrderServiceClient
	deliveryTariftService order_service.DeliveryTariffServiceClient

	// user service
	userService    user_service.UserServiceClient
	clientService  user_service.ClientServiceClient
	courierService user_service.CourierServiceClient
	branchService  user_service.BranchServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connCatalogService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CatalogServiceHost, conf.CatalogServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connOrderService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.OrderServiceHost, conf.OrderServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		categoryService: catalog_service.NewCategoryServiceClient(connCatalogService),
		orderService:    order_service.NewOrderServiceClient(connOrderService),
		userService:     user_service.NewUserServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) Category() catalog_service.CategoryServiceClient {
	return g.categoryService
}

func (g *grpcClients) Product() catalog_service.ProductServiceClient {
	return g.productService
}

func (g *grpcClients) Order() order_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) DeliveryTariff() order_service.DeliveryTariffServiceClient {
	return g.deliveryTariftService
}

func (g *grpcClients) User() user_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) Client() user_service.ClientServiceClient {
	return g.clientService
}

func (g *grpcClients) Courier() user_service.CourierServiceClient {
	return g.courierService
}

func (g *grpcClients) Branch() user_service.BranchServiceClient {
	return g.branchService
}
