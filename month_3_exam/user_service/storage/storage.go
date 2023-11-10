package storage

import (
	"context"
	pb "user_service/genproto"
)

type StorageI interface {
	Branch() BranchI
	Users() UsersI
	Couriers() CouriersI
	Clients() ClientsI
}

type BranchI interface {
	Create(context.Context, *pb.CreateBranchRequest) (*pb.Response, error)
	Get(context.Context, *pb.IdRequest) (*pb.Branch, error)
	GetList(context.Context, *pb.ListBranchRequest) (*pb.ListBranchResponse, error)
	Update(context.Context, *pb.UpdateBranchRequest) (string, error)
	Delete(context.Context, *pb.IdRequest) (string, error)

	GetListActive(context.Context, *pb.ListBranchActiveRequest) (*pb.ListBranchResponse, error)
}

type UsersI interface {
	Create(context.Context, *pb.CreateUsersRequest) (*pb.Response, error)
	Get(context.Context, *pb.IdRequest) (*pb.Users, error)
	GetList(context.Context, *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
	Update(context.Context, *pb.UpdateUsersRequest) (string, error)
	Delete(context.Context, *pb.IdRequest) (string, error)

	GetByLogin(context.Context, *pb.IdRequest) (*pb.Users, error)
}

type CouriersI interface {
	Create(context.Context, *pb.CreateCouriersRequest) (*pb.Response, error)
	Get(context.Context, *pb.IdRequest) (*pb.Couriers, error)
	GetList(context.Context, *pb.ListCouriersRequest) (*pb.ListCouriersResponse, error)
	Update(context.Context, *pb.UpdateCouriersRequest) (string, error)
	Delete(context.Context, *pb.IdRequest) (string, error)

	GetByLogin(context.Context, *pb.IdRequest) (*pb.Couriers, error)
}

type ClientsI interface {
	Create(context.Context, *pb.CreateClientsRequest) (*pb.Response, error)
	Get(context.Context, *pb.IdRequest) (*pb.Clients, error)
	GetList(context.Context, *pb.ListClientsRequest) (*pb.ListClientsResponse, error)
	Update(context.Context, *pb.UpdateClientsRequest) (string, error)
	Delete(context.Context, *pb.IdRequest) (string, error)

	UpdateOrder(context.Context, *pb.UpdateClientsOrderRequest) (string, error)
}
