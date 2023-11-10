package memory

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"errors"
)

type clientRepo struct {
	clients []models.Client
}

func NewClientRepo() *clientRepo {
	return &clientRepo{clients: make([]models.Client, 0)}
}

func (c *clientRepo) CreateClient(req models.CreateClient) (int, error) {
	var id int
	if len(c.clients) == 0 {
		id = 1
	} else {
		id = c.clients[len(c.clients)-1].Id + 1
	}

	c.clients = append(c.clients, models.Client{
		Id:         id,
		Name:       req.Name,
		Card_Id:    req.Card_Id,
		Created_at: req.Created_at,
	})

	return id, nil
}

func (p *clientRepo) UpdateClient(req models.Client) (string, error) {
	for i, v := range p.clients {
		if v.Id == req.Id {
			p.clients[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not client found with ID")
}

func (p *clientRepo) GetClient(req models.IdRequest) (models.Client, error) {
	for _, v := range p.clients {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Client{}, errors.New("not client found with ID")
}

func (p *clientRepo) GetAllClient(req models.GetAllClientRequest) (resp models.GetAllClientResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(p.clients) {
		resp.Clientes = []models.Client{}
		resp.Count = len(p.clients)
		return
	} else if end > len(p.clients) {
		return models.GetAllClientResponse{
			Clientes: p.clients[start:],
			Count:    len(p.clients),
		}, nil
	}

	return models.GetAllClientResponse{
		Clientes: p.clients[start:end],
		Count:    len(p.clients),
	}, nil
}

func (p *clientRepo) DeleteClient(req models.IdRequest) (resp string, err error) {
	for i, v := range p.clients {
		if v.Id == req.Id {
			if i == len(p.clients)-1 {
				p.clients = p.clients[:i]
			} else {
				p.clients = append(p.clients[:i], p.clients[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
