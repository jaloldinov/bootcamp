package storage

import "backend_bootcamp_17_07_2023/lesson_8/project/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffesI
}

type BranchesI interface {
	CreateBranch(models.CreateBranch) (int, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}

type StaffesI interface {
	CreateStaff(models.CreateStaff) (int, error)
	UpdateStaff(models.Staff) (string, error)
	GetStaff(models.IdRequest) (models.Staff, error)
	GetAllStaff(models.GetAllStaffRequest) (models.GetAllStaff, error)
	DeleteStaff(models.IdRequest) (string, error)
}
