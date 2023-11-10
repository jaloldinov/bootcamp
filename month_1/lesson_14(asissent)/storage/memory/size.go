package memory

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"errors"
)

type sizeRepo struct {
	sizes []models.Size
}

func NewSizeRepo() *sizeRepo {
	return &sizeRepo{sizes: make([]models.Size, 0)}
}

func (c *sizeRepo) CreateSize(req models.CreateSize) (int, error) {
	var id int
	if len(c.sizes) == 0 {
		id = 1
	} else {
		id = c.sizes[len(c.sizes)-1].Id + 1
	}

	c.sizes = append(c.sizes, models.Size{
		Id:         id,
		Name:       req.Name,
		Price:      req.Price,
		Created_at: req.Created_at,
	})

	return id, nil
}

func (p *sizeRepo) UpdateSize(req models.Size) (string, error) {
	for i, v := range p.sizes {
		if v.Id == req.Id {
			p.sizes[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not size found with ID")
}

func (p *sizeRepo) GetSize(req models.IdRequest) (models.Size, error) {
	for _, v := range p.sizes {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Size{}, errors.New("not size found with ID")
}

func (p *sizeRepo) GetAllSize(req models.GetAllSizeRequest) (resp models.GetAllSizeResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(p.sizes) {
		resp.Sizes = []models.Size{}
		resp.Count = len(p.sizes)
		return
	} else if end > len(p.sizes) {
		return models.GetAllSizeResponse{
			Sizes: p.sizes[start:],
			Count: len(p.sizes),
		}, nil
	}

	return models.GetAllSizeResponse{
		Sizes: p.sizes[start:end],
		Count: len(p.sizes),
	}, nil
}

func (p *sizeRepo) DeleteSize(req models.IdRequest) (resp string, err error) {
	for i, v := range p.sizes {
		if v.Id == req.Id {
			if i == len(p.sizes)-1 {
				p.sizes = p.sizes[:i]
			} else {
				p.sizes = append(p.sizes[:i], p.sizes[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
