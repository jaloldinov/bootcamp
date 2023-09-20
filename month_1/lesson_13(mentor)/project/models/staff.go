package models

type CreateStaff struct {
	BranchId int
	TariffId int
	TypeId   int
	Name     string
	Balance  float64
}

type Staff struct {
	Id       int
	BranchId int
	TariffId int
	TypeId   int
	Name     string
	Balance  float64
}

type GetAllStaffRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllStaff struct {
	Staffes []Staff
	Count   int
}
