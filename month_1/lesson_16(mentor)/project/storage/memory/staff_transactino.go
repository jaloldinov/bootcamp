package memory

import (
	"encoding/json"
	"errors"
	"lesson_15/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type transactionRepo struct {
	fileName string
}

func NewTransactionRepo(fileName string) *transactionRepo {
	return &transactionRepo{fileName: fileName}
}

func (t *transactionRepo) CreateTransaction(req models.CreateTransaction) (string, error) {
	id := uuid.NewString()
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	transactions = append(transactions, models.Transaction{
		Id:          id,
		Type:        req.Type,
		Amount:      req.Amount,
		Source_type: req.Source_type,
		Text:        req.Text,
		Sale_id:     req.Sale_id,
		Staff_id:    req.Staff_id,
		Created_at:  time.Now().Format("2006-01-02 15:04:05"),
	})
	err = t.write(transactions)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t *transactionRepo) UpdateTransaction(req models.Transaction) (string, error) {
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	for i, v := range transactions {
		if v.Id == req.Id {
			transactions[i] = req
			transactions[i].Created_at = time.Now().Format("2006-01-02 15:04:05")
			err = t.write(transactions)
			if err != nil {
				return "", err
			}
			return "updated", nil
		}
	}
	return "", errors.New("not transaction found with ID")
}

func (t *transactionRepo) GetTransaction(req models.IdRequest) (models.Transaction, error) {
	transactions, err := t.read()
	if err != nil {
		return models.Transaction{}, err
	}
	for _, v := range transactions {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Transaction{}, errors.New("not transaction found with ID")
}

func (t *transactionRepo) GetAllTransaction(req models.GetAllTransactionRequest) (resp models.GetAllTransactionResponse, err error) {
	transactions, err := t.read()
	if err != nil {
		return models.GetAllTransactionResponse{}, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	var filtered []models.Transaction

	for _, v := range transactions {
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

func (t *transactionRepo) DeleteTransaction(req models.IdRequest) (resp string, err error) {
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	for i, v := range transactions {
		if v.Id == req.Id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			err = t.write(transactions)
			if err != nil {
				return "", err
			}
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}

func (u *transactionRepo) read() ([]models.Transaction, error) {
	var (
		transactions []models.Transaction
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return transactions, nil
}

func (u *transactionRepo) write(transasctions []models.Transaction) error {
	body, err := json.Marshal(transasctions)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
