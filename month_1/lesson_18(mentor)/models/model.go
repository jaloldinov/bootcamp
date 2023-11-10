package models

type BranchTransactionCount struct {
	BranchID   int
	BranchName string
	Count      int
}
type BranchTransactionCategoryCount struct {
	BranchID     int
	BranchName   string
	CategoryName string
	Count        int
}
type Transaction struct {
	ID        int    `json:"id"`
	BranchID  int    `json:"branch_id"`
	ProductID int    `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}
type Branch struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Adress string `json:"adress"`
}

type Products struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
}

type ProductTop struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int
}

type BranchProductPrice struct {
	BranchID   int
	BranchName string
	Sum        int
}

type PlusMinus struct {
	TranType  string
	Quantity  int
	TranPlus  int
	TranMinus int
	SumPlus   int
	SumMinus  int
}

type ProductIncome struct {
	Day   string
	Count int
}

type Product7task struct {
	BranchId  int `json:"branch_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ModelFor9 struct {
	Name string
	Sum  int
}
