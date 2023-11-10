package memory

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"errors"
)

type productRepo struct {
	products []models.Product
}

func NewProductRepo() *productRepo {
	return &productRepo{products: make([]models.Product, 0)}
}

func (p *productRepo) CreateProduct(req models.CreateProduct) (int, error) {
	var id int
	if len(p.products) == 0 {
		id = 1
	} else {
		id = p.products[len(p.products)-1].Id + 1
	}

	p.products = append(p.products, models.Product{
		Id:         id,
		Card_Id:    req.Card_Id,
		Size_Id:    req.Size_Id,
		Created_at: req.Created_at,
	})
	return id, nil
}

func (p *productRepo) UpdateProduct(req models.Product) (string, error) {
	for i, v := range p.products {
		if v.Id == req.Id {
			p.products[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not product found with ID")
}

func (p *productRepo) GetProduct(req models.IdRequest) (models.Product, error) {
	for _, v := range p.products {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Product{}, errors.New("not product found with ID")
}

func (p *productRepo) GetAllProduct(req models.GetAllProductRequest) (resp models.GetAllProductResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(p.products) {
		resp.Productes = []models.Product{}
		resp.Count = len(p.products)
		return
	} else if end > len(p.products) {
		return models.GetAllProductResponse{
			Productes: p.products[start:],
			Count:     len(p.products),
		}, nil
	}

	return models.GetAllProductResponse{
		Productes: p.products[start:end],
		Count:     len(p.products),
	}, nil
}

func (p *productRepo) DeleteProduct(req models.IdRequest) (resp string, err error) {
	for i, v := range p.products {
		if v.Id == req.Id {
			if i == len(p.products)-1 {
				p.products = p.products[:i]
			} else {
				p.products = append(p.products[:i], p.products[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
