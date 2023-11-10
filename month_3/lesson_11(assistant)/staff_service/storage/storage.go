package storage

import (
	"context"
	pb "staff_service/genproto"
	"time"
)

type StorageI interface {
	Staff() StaffI
	Tariff() TariffI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type StaffI interface {
	CreateStaff(context.Context, *pb.CreateStaffRequest) (string, error)
	GetStaff(context.Context, *pb.IdRequest) (*pb.Staff, error)
	GetAllStaff(context.Context, *pb.ListStaffRequest) (*pb.ListStaffResponse, error)
	UpdateStaff(context.Context, *pb.UpdateStaffRequest) (string, error)
	DeleteStaff(context.Context, *pb.IdRequest) (string, error)
}

type TariffI interface {
	CreateTariff(context.Context, *pb.CreateTariffRequest) (string, error)
	GetTariff(context.Context, *pb.IdRequest) (*pb.Tariff, error)
	GetAllTariff(context.Context, *pb.ListTariffRequest) (*pb.ListTariffResponse, error)
	UpdateTariff(context.Context, *pb.UpdateTariffRequest) (string, error)
	DeleteTariff(context.Context, *pb.IdRequest) (string, error)
}
