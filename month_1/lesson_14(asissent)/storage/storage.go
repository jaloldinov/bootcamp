package storage

import "backend_bootcamp_17_07_2023/lesson_14/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffesI
	Product() ProductI
	Client() ClientI
	Card() CardI
	Size() SizeI
	Sales() SalesI
	Transaction() TransactionI
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

type ProductI interface {
	CreateProduct(models.CreateProduct) (int, error)
	UpdateProduct(models.Product) (string, error)
	GetProduct(models.IdRequest) (models.Product, error)
	GetAllProduct(models.GetAllProductRequest) (models.GetAllProductResponse, error)
	DeleteProduct(models.IdRequest) (string, error)
}

type ClientI interface {
	CreateClient(models.CreateClient) (int, error)
	UpdateClient(models.Client) (string, error)
	GetClient(models.IdRequest) (models.Client, error)
	GetAllClient(models.GetAllClientRequest) (models.GetAllClientResponse, error)
	DeleteClient(models.IdRequest) (string, error)
}

type CardI interface {
	CreateCard(models.CreateCard) (int, error)
	UpdateCard(models.Card) (string, error)
	GetCard(models.IdRequest) (models.Card, error)
	GetAllCard(models.GetAllCardRequest) (models.GetAllCardResponse, error)
	DeleteCard(models.IdRequest) (string, error)
}

type SizeI interface {
	CreateSize(models.CreateSize) (int, error)
	UpdateSize(models.Size) (string, error)
	GetSize(models.IdRequest) (models.Size, error)
	GetAllSize(models.GetAllSizeRequest) (models.GetAllSizeResponse, error)
	DeleteSize(models.IdRequest) (string, error)
}

type TransactionI interface {
	CreateTransaction(models.CreateTransaction) (int, error)
	UpdateTransaction(models.Transaction) (string, error)
	GetTransaction(models.IdRequest) (models.Transaction, error)
	GetAllTransaction(models.GetAllTransactionRequest) (models.GetAllTransaction, error)
	DeleteTransaction(models.IdRequest) (string, error)
}

type SalesI interface {
	CreateSales(models.CreateSales) (int, error)
	UpdateSales(models.Sales) (string, error)
	GetSales(models.IdRequest) (models.Sales, error)
	GetAllSales(models.GetAllSalesRequest) (models.GetAllSalesResponse, error)
	DeleteSales(models.IdRequest) (string, error)
}
