package models

type CreateSales struct {
	Name             string
	Price            float64
	Payment_Type     int
	Status           int
	Client_id        int
	Branch_id        int
	Shop_asissent_id int
	Cashier_id       int
	Created_at       string
}

type Sales struct {
	Id               int
	Name             string
	Price            float64
	Payment_Type     int
	Status           int
	Client_id        int
	Branch_id        int
	Shop_asissent_id int
	Cashier_id       int
	Created_at       string
}

type GetAllSalesRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllSalesResponse struct {
	Saleses []Sales
	Count   int
}
