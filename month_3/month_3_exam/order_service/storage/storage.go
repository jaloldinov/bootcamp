package storage

import (
	"context"
	pb "order_service/genproto"
)

type StorageI interface {
	Order() OrderI
	DeliveryTariff() DeliveryTariffI
}

type OrderI interface {
	Create(context.Context, *pb.CreateOrderRequest) (string, error)
	Get(context.Context, *pb.IdRequest) (*pb.Order, error)
	GetList(context.Context, *pb.ListOrderRequest) (*pb.ListOrderResponse, error)
	Update(context.Context, *pb.UpdateOrderRequest) (string, error)
	Delete(context.Context, *pb.IdRequest) (string, error)

	UpdateStatus(context.Context, *pb.UpdateOrderStatusRequest) (string, error)
	GetListByCourierId(context.Context, *pb.IdRequest) (*pb.ListOrderResponse, error)
}

type DeliveryTariffI interface {
	Create(context.Context, *pb.CreateDeliveryTariffRequest) (string, error)
	Get(context.Context, *pb.IdRequest) (*pb.DeliveryTariff, error)
	GetList(context.Context, *pb.ListDeliveryTariffRequest) (*pb.ListDeliveryTariffResponse, error)
	Update(context.Context, *pb.UpdateDeliveryTariffRequest) (string, error)
	Delete(context.Context, *pb.IdRequest) (string, error)
}
