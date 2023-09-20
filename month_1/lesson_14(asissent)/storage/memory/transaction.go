package memory

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"errors"
)

type transactionRepo struct {
	transactions []models.Transaction
}

func NewTransactionRepo() *transactionRepo {
	return &transactionRepo{transactions: make([]models.Transaction, 0)}
}

func (t *transactionRepo) CreateTransaction(req models.CreateTransaction) (int, error) {
	var id int
	if len(t.transactions) == 0 {
		id = 1
	} else {
		id = t.transactions[len(t.transactions)-1].Id + 1
	}

	t.transactions = append(t.transactions, models.Transaction{
		Id:          id,
		Amount:      req.Amount,
		Source_type: req.Source_type,
		Text:        req.Text,
		Sale_id:     req.Sale_id,
		Staff_id:    req.Staff_id,
		Created_at:  req.Created_at,
	})

	return id, nil
}

func (p *transactionRepo) UpdateTransaction(req models.Transaction) (string, error) {
	for i, v := range p.transactions {
		if v.Id == req.Id {
			p.transactions[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not transaction found with ID")
}

func (p *transactionRepo) GetTransaction(req models.IdRequest) (models.Transaction, error) {
	for _, v := range p.transactions {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Transaction{}, errors.New("not transaction found with ID")
}

func (p *transactionRepo) GetAllTransaction(req models.GetAllTransactionRequest) (resp models.GetAllTransaction, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(p.transactions) {
		resp.Transactions = []models.Transaction{}
		resp.Count = len(p.transactions)
		return
	} else if end > len(p.transactions) {
		return models.GetAllTransaction{
			Transactions: p.transactions[start:],
			Count:        len(p.transactions),
		}, nil
	}

	return models.GetAllTransaction{
		Transactions: p.transactions[start:end],
		Count:        len(p.transactions),
	}, nil
}

func (p *transactionRepo) DeleteTransaction(req models.IdRequest) (resp string, err error) {
	for i, v := range p.transactions {
		if v.Id == req.Id {
			if i == len(p.transactions)-1 {
				p.transactions = p.transactions[:i]
			} else {
				p.transactions = append(p.transactions[:i], p.transactions[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
