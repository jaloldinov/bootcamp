package models

type CreateTransaction struct {
	Type        int
	Amount      int
	Source_type string
	Text        string
	Sale_id     float64
	Staff_id    int
	Created_at  string
}

type Transaction struct {
	Id          int
	Amount      int
	Source_type string
	Text        string
	Sale_id     float64
	Staff_id    int
	Created_at  string
}

type GetAllTransactionRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllTransaction struct {
	Transactions []Transaction
	Count        int
}
