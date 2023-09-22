package storage

import "market/models"

type StorageI interface {
	Close()
	Branch() BranchRepoI
}

type BranchRepoI interface {
	Create(*models.CreateBranch) (string, error)
	GetByID(*models.BranchPrimaryKey) (*models.Branch, error)
	GetList(*models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(*models.UpdateBranch) (string, error)
	Delete(*models.BranchPrimaryKey) error
}
