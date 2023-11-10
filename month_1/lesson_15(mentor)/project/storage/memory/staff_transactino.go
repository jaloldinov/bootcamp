package memory

import (
	"errors"
	"lesson_15/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

type transactionRepo struct {
	transactions []models.Transaction
}

func NewTransactionRepo() *transactionRepo {
	return &transactionRepo{transactions: make([]models.Transaction, 0)}
}

func (t *transactionRepo) CreateTransaction(req models.CreateTransaction) (string, error) {
	id := uuid.New()

	t.transactions = append(t.transactions, models.Transaction{
		Id:          id.String(),
		Amount:      req.Amount,
		Source_type: req.Source_type,
		Text:        req.Text,
		Sale_id:     req.Sale_id,
		Staff_id:    req.Staff_id,
		Created_at:  time.Now().Format("2006-01-02 15:04:05"),
	})

	return id.String(), nil
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

func (p *transactionRepo) GetAllTransaction(req models.GetAllTransactionRequest) (resp models.GetAllTransactionResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	var filtered []models.Transaction

	for _, v := range p.transactions {
		if strings.Contains(v.Text, req.Text) {
			filtered = append(filtered, v)
		}
	}

	if start > len(filtered) {
		resp.Transactions = []models.Transaction{}
		resp.Count = len(filtered)
		return
	} else if end > len(filtered) {
		return models.GetAllTransactionResponse{
			Transactions: filtered[start:],
			Count:        len(filtered),
		}, nil
	}

	return models.GetAllTransactionResponse{
		Transactions: filtered[start:end],
		Count:        len(filtered),
	}, nil
}

func (p *transactionRepo) DeleteTransaction(req models.IdRequest) (resp string, err error) {
	for i, v := range p.transactions {
		if v.Id == req.Id {
			p.transactions = append(p.transactions[:i], p.transactions[i+1:]...)
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}
