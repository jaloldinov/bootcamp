package grpc

import (
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/grpc/service"

	"sale_service/pkg/logger"
	"sale_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	sale_service.RegisterSaleServiceServer(grpcServer, service.NewSaleService(cfg, log, strg))
	sale_service.RegisterSaleProductServiceServer(grpcServer, service.NewSaleProductService(cfg, log, strg))
	sale_service.RegisterStaffTransactionServiceServer(grpcServer, service.NewStaffTransactionService(cfg, log, strg))
	sale_service.RegisterBranchPrTransactionServiceServer(grpcServer, service.NewBranchProductTransactionsService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
