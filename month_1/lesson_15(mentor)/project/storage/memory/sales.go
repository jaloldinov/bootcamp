package memory

import (
	"errors"
	"lesson_15/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

type saleRepo struct {
	sales []models.Sales
}

func NewSaleRepo() *saleRepo {
	return &saleRepo{sales: make([]models.Sales, 0)}
}

func (c *saleRepo) CreateSale(req models.CreateSales) (string, error) {
	id := uuid.New()

	c.sales = append(c.sales, models.Sales{
		Id:               id.String(),
		Client_name:      req.Client_name,
		Price:            req.Price,
		Payment_Type:     req.Payment_Type,
		Status:           req.Status,
		Branch_id:        req.Branch_id,
		Shop_asissent_id: req.Shop_asissent_id,
		Cashier_id:       req.Cashier_id,
		Created_at:       time.Now().Format("2006-01-02 15:04:05"),
	})
	return id.String(), nil
}

func (c *saleRepo) UpdateSale(req models.Sales) (string, error) {
	for i, v := range c.sales {
		if v.Id == req.Id {
			c.sales[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not sale found with ID")
}

func (c *saleRepo) GetSale(req models.IdRequest) (resp models.Sales, err error) {
	for _, v := range c.sales {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Sales{}, errors.New("not found")
}

func (c *saleRepo) GetAllSale(req models.GetAllSalesRequest) (resp models.GetAllSalesResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	var filtered []models.Sales

	for _, v := range c.sales {
		if strings.Contains(v.Client_name, req.Client_name) {
			filtered = append(filtered, v)
		}
	}

	if start > len(filtered) {
		resp.Sales = []models.Sales{}
		resp.Count = len(filtered)
		return
	} else if end > len(filtered) {
		return models.GetAllSalesResponse{
			Sales: filtered[start:],
			Count: len(filtered),
		}, nil
	}

	return models.GetAllSalesResponse{
		Sales: filtered[start:end],
		Count: len(filtered),
	}, nil

}

func (c *saleRepo) DeleteSale(req models.IdRequest) (resp string, err error) {
	for i, v := range c.sales {
		if v.Id == req.Id {
			c.sales = append(c.sales[:i], c.sales[i+1:]...)
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}
