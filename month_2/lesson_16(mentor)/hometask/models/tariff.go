package models

type CreateStaffTarif struct {
	Name          string
	Type          int // (1-fixed, 2-percent)
	AmountForCash float64
	AmountForCard float64
}

type StaffTarif struct {
	Id            string
	Name          string
	Type          int // (1-fixed, 2-percent)
	AmountForCash float64
	AmountForCard float64
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
