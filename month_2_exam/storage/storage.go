package storage

import "market/models"

type StorageI interface {
	Close()
	Branch() BranchRepoI
	Category() CategoryRepoI
	Product() ProductRepoI
	ComingTable() ComingTableRepoI
}

type BranchRepoI interface {
	Create(*models.CreateBranch) (string, error)
	GetByID(*models.BranchPrimaryKey) (*models.Branch, error)
	GetList(*models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(*models.UpdateBranch) (string, error)
	Delete(*models.BranchPrimaryKey) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (*models.Category, error)
	GetList(*models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(*models.UpdateCategory) (string, error)
	Delete(*models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	GetByID(*models.ProductPrimaryKey) (*models.Product, error)
	GetList(*models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(*models.UpdateProduct) (string, error)
	Delete(*models.ProductPrimaryKey) error
}

type ComingTableRepoI interface {
	Create(*models.CreateComingTable) (string, error)
	GetByID(*models.ComingTableComingIdKey) (*models.ComingTable, error)
	GetList(*models.ComingTableGetListRequest) (*models.ComingTableGetListResponse, error)
	Update(*models.UpdateComingTable) (string, error)
	UpdateStatus(*models.ComingTablePrimaryKey) (string, error)
	Delete(*models.ComingTablePrimaryKey) error
}
