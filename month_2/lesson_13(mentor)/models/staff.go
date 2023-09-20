package models

type CreateStaff struct {
	BranchId string
	TariffId string
	TypeId   StaffType
	Name     string
	Balance  float64
}

type Staff struct {
	Id       string
	BranchId string
	TariffId string
	TypeId   StaffType
	Name     string
	Balance  float64
	Login    string
	Password string
	Phone    string
}

type ExistsReq struct {
	Phone string
}
type StaffTop struct {
	BranchId string
	TypeId   StaffType
	Name     string
	Money    int
}
type ChangeBalance struct {
	Id      string
	Balance float64
}

type StaffType string

const (
	Cashier       StaffType = "cashier"
	ShopAssistant StaffType = "shop_assistant"
)

type GetAllStaffRequest struct {
	Page        int
	Limit       int
	Type        StaffType
	Name        string
	BalanceFrom float64
	BalanceTo   float64
}

type GetAllStaff struct {
	Staffes []Staff
	Count   int
}
