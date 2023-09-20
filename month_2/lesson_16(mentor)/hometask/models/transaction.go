package models

type CreateTransaction struct {
	Type        string //(withdraw,topup)
	Amount      int
	Source_type string //(sales,bonus)
	Text        string
	Sale_id     string
	Staff_id    string
	Created_at  string
}

type Transaction struct {
	Id          string
	Type        string //(withdraw,topup)
	Amount      int
	Source_type string //(sales,bonus)
	Text        string
	Sale_id     string
	Staff_id    string
	Created_at  string
}

type GetAllTransactionRequest struct {
	Page  int
	Limit int
	Text  string
}

type GetAllTransactionResponse struct {
	Transactions []Transaction
	Count        int
}

type TopWorkerRequest struct {
	Type     string
	FromDate string
	ToDate   string
}

type TopWorkerRespond struct {
	Staff []TopWorker
}

type TopWorker struct {
	BranchName string
	StaffName  string
	StaffType  string
	EarnedSum  int
}
