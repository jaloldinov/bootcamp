package service

import (
	"context"
	"staff_service/config"
	staff_service "staff_service/genproto"
	"staff_service/pkg/logger"
	"staff_service/storage"
)

type StaffService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	staff_service.UnimplementedStaffServiceServer
}

func NewStaffService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *StaffService {
	return &StaffService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *StaffService) Create(ctx context.Context, req *staff_service.CreateStaffRequest) (*staff_service.CreateStaffResponse, error) {
	id, err := b.storage.Staff().CreateStaff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.CreateStaffResponse{Id: id}, nil
}

func (b *StaffService) Get(ctx context.Context, req *staff_service.IdRequest) (*staff_service.GetStaffResponse, error) {
	branch, err := b.storage.Staff().GetStaff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.GetStaffResponse{Staff: branch}, nil
}

func (b *StaffService) List(ctx context.Context, req *staff_service.ListStaffRequest) (*staff_service.ListStaffResponse, error) {
	staffs, err := b.storage.Staff().GetAllStaff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.ListStaffResponse{Staffs: staffs.Staffs,
		Count: staffs.Count}, nil
}

func (s *StaffService) Update(ctx context.Context, req *staff_service.UpdateStaffRequest) (*staff_service.Response, error) {
	resp, err := s.storage.Staff().UpdateStaff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.Response{Message: resp}, nil
}

func (s *StaffService) Delete(ctx context.Context, req *staff_service.IdRequest) (*staff_service.Response, error) {
	resp, err := s.storage.Staff().DeleteStaff(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.Response{Message: resp}, nil
}
