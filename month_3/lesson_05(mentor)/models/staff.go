package models

import "time"

type CreateStaff struct {
	BranchID string `json:"branch_id"`
	TariffID string `json:"tariff_id"`
	Name     string `json:"name"`
	Type     string `json:"staff_type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Staff struct {
	ID        string    `json:"id"`
	BranchID  string    `json:"branch_id"`
	TariffID  string    `json:"tariff_id"`
	Type      string    `json:"staff_type"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExistsReq struct {
	Phone string `json:"phone"`
}

type StaffTop struct {
	BranchID string `json:"branch_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Money    int    `json:"money"`
}

type ChangeBalance struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type GetAllStaffRequest struct {
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	BalanceFrom float64 `json:"balance_from"`
	BalanceTo   float64 `json:"balance_to"`
}

type GetAllStaff struct {
	Staffs []Staff `json:"staffs"`
	Count  int     `json:"count"`
}

type UpdateBalanceRequest struct {
	SaleId          string
	TransactionType string
	SourceType      string
	Cashier         StaffIdAmount
	ShopAssisstant  StaffIdAmount
	Text            string
}

type StaffIdAmount struct {
	StaffId string
	Amount  float32
}

type GetPasswordById struct {
	OldPassword string `json:"old_password"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordRequest struct {
	Id          string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
