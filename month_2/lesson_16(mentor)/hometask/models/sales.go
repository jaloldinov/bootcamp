package models

type CreateSales struct {
	Client_name      string
	Branch_id        string
	Shop_asissent_id string
	Cashier_id       string
	Price            float64
	Payment_Type     int // 1 for card, 2 for cash
	Status           int // 1 for success, 2 for cancel
	Created_at       string
}

type Sales struct {
	Id               string
	Client_name      string
	Branch_id        string
	Shop_asissent_id string
	Cashier_id       string
	Price            float64
	Payment_Type     int // 1 for card, 2 for cash
	Status           int // 1 for success, 2 for cancel
	Created_at       string
}

type GetAllSalesRequest struct {
	Page        int
	Limit       int
	Client_name string
}

type GetAllSalesResponse struct {
	Sales []Sales
	Count int
}

type SaleTopBranch struct {
	Day         string
	BranchId    string
	SalesAmount float64
}

type SaleCountSumBranch struct {
	BranchId    string
	Count       int
	SalesAmount float64
}
