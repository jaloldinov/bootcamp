package models

type CreateStaffTarif struct {
	Name          string
	Type          string // (fixed, percent)
	AmountForCash int
	AmountForCard int
}

type StaffTarif struct {
	Id            string
	Name          string
	Type          string // (fixed, percent)
	AmountForCash int
	AmountForCard int
	CreatedAt     string
}

type GetAllStaffTarifRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllStaffTarif struct {
	StaffTarifs []StaffTarif
	Count       int
}
