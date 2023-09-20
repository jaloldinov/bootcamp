package models

type BranchTransactionCount struct {
	BranchID   int
	BranchName string
	Count      int
}

type Task2 struct {
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
	UserID    int    `json:"user_id"`
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
	CategoryID int    `json:"category_id"`
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

type ProductIncome struct {
	Day   string
	Count int
}

type Category struct {
	Id   int
	Name string
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

type User struct {
	ID   int
	Name string
}
