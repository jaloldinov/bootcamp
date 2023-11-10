package storage

import (
	"app/models"
	"context"
)

type StorageI interface {
	Branch() BranchesI
	Tariff() TariffsI
	Staff() StaffesI
	Sales() SalesI
	Transaction() TransactionI
	BiznesLoggic() BiznesLogicI
	Close()
}

type BranchesI interface {
	CreateBranch(context.Context, *models.CreateBranch) (string, error)
	GetBranch(context.Context, *models.IdRequest) (*models.Branch, error)
	GetAllBranch(context.Context, *models.GetAllBranchRequest) (*models.GetAllBranch, error)
	UpdateBranch(context.Context, *models.Branch) (string, error)
	DeleteBranch(context.Context, *models.IdRequest) (string, error)
}

type TariffsI interface {
	CreateStaffTarif(context.Context, *models.CreateStaffTarif) (string, error)
	GetStaffTarif(context.Context, *models.IdRequest) (*models.StaffTarif, error)
	GetAllStaffTarif(context.Context, *models.GetAllStaffTarifRequest) (*models.GetAllStaffTarif, error)
	UpdateStaffTarif(context.Context, *models.StaffTarif) (string, error)
	DeleteStaffTarif(context.Context, *models.IdRequest) (string, error)
}

type StaffesI interface {
	CreateStaff(context.Context, *models.CreateStaff) (string, error)
	UpdateStaff(context.Context, *models.Staff) (string, error)
	GetStaff(context.Context, *models.IdRequest) (*models.Staff, error)
	// GetByLogin(context.Context, *models.LoginRequest) (*models.Staff, error)
	GetAllStaff(context.Context, *models.GetAllStaffRequest) (*models.GetAllStaff, error)
	DeleteStaff(context.Context, *models.IdRequest) (string, error)
	// ChangeBalance(models.ChangeBalance) (string, error)
	// Exists(models.ExistsReq) bool
	GetByUsername(context.Context, *models.RequestByUsername) (*models.Staff, error)

	ChangePassword(context.Context, *models.ChangePasswordRequest) (string, error)
}

type SalesI interface {
	CreateSale(context.Context, *models.CreateSales) (string, error)
	UpdateSale(context.Context, *models.Sales) (string, error)
	GetSale(context.Context, *models.IdRequest) (*models.Sales, error)
	GetAllSale(context.Context, *models.GetAllSalesRequest) (*models.GetAllSalesResponse, error)
	DeleteSale(context.Context, *models.IdRequest) (string, error)
	// GetTopSaleBranch(context.Context) (resp map[string]models.SaleTopBranch, err error)
	// GetSaleCountBranch(context.Context) (resp map[string]models.SaleCountSumBranch, err error)
	// CancelSale(context.Context,*models.IdRequest) (string, error)
}

type TransactionI interface {
	CreateTransaction(context.Context, *models.CreateTransaction) (string, error)
	UpdateTransaction(context.Context, *models.Transaction) (string, error)
	GetTransaction(context.Context, *models.IdRequest) (*models.Transaction, error)
	GetAllTransaction(context.Context, *models.GetAllTransactionRequest) (*models.GetAllTransactionResponse, error)
	DeleteTransaction(context.Context, *models.IdRequest) (string, error)
	// GetTopStaffs(models.TopWorkerRequest) (map[string]models.StaffTop, error)
}

type BiznesLogicI interface {
	GetTopStaff(context.Context, *models.TopStaffRequest) (*models.TopStaffResponse, error)
}
