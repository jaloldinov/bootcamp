package storage

import (
	"lesson_20/models"
)

type StorageI interface {
	Branch() BranchesI
	Staff() StaffesI
	Sales() SalesI
	Transaction() TransactionI
	StaffTarif() StaffTarifsI
}

type BranchesI interface {
	CreateBranch(models.CreateBranch) (string, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}

type StaffesI interface {
	CreateStaff(models.CreateStaff) (string, error)
	UpdateStaff(models.Staff) (string, error)
	GetStaff(models.IdRequest) (models.Staff, error)
	GetByLogin(models.LoginRequest) (models.Staff, error)
	GetAllStaff(models.GetAllStaffRequest) (models.GetAllStaff, error)
	DeleteStaff(models.IdRequest) (string, error)
	ChangeBalance(models.ChangeBalance) (string, error)
	Exists(models.ExistsReq) bool
}

type TransactionI interface {
	CreateTransaction(models.CreateTransaction) (string, error)
	UpdateTransaction(models.Transaction) (string, error)
	GetTransaction(models.IdRequest) (models.Transaction, error)
	GetAllTransaction(models.GetAllTransactionRequest) (models.GetAllTransactionResponse, error)
	DeleteTransaction(models.IdRequest) (string, error)
	GetTopStaffs(models.TopWorkerRequest) (map[string]models.StaffTop, error)
}

type SalesI interface {
	CreateSale(models.CreateSales) (string, error)
	UpdateSale(models.Sales) (string, error)
	GetSale(models.IdRequest) (models.Sales, error)
	GetAllSale(models.GetAllSalesRequest) (models.GetAllSalesResponse, error)
	DeleteSale(models.IdRequest) (string, error)
	GetTopSaleBranch() (resp map[string]models.SaleTopBranch, err error)
	GetSaleCountBranch() (resp map[string]models.SaleCountSumBranch, err error)
}

type StaffTarifsI interface {
	CreateStaffTarif(models.CreateStaffTarif) (string, error)
	UpdateStaffTarif(req models.StaffTarif) (string, error)
	GetStaffTarif(req models.IdRequest) (resp models.StaffTarif, err error)
	GetAllStaffTarif(req models.GetAllStaffTarifRequest) (resp models.GetAllStaffTarif, err error)
	DeleteStaffTarif(req models.IdRequest) (string, error)
}
