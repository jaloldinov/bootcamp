package memory

import (
	"backend_bootcamp_17_07_2023/lesson_14/models"
	"errors"
)

type cardRepo struct {
	cards []models.Card
}

func NewCardRepo() *cardRepo {
	return &cardRepo{cards: make([]models.Card, 0)}
}

func (c *cardRepo) CreateCard(req models.CreateCard) (int, error) {
	var id int
	if len(c.cards) == 0 {
		id = 1
	} else {
		id = c.cards[len(c.cards)-1].Id + 1
	}

	c.cards = append(c.cards, models.Card{
		Id:         id,
		Name:       req.Name,
		Quantity:   req.Quantity,
		Product_id: req.Product_id,
		Created_at: req.Created_at,
	})
	return id, nil
}

func (c *cardRepo) UpdateCard(req models.Card) (string, error) {
	for i, v := range c.cards {
		if v.Id == req.Id {
			c.cards[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not card found with ID")
}

func (c *cardRepo) GetCard(req models.IdRequest) (resp models.Card, err error) {
	for _, v := range c.cards {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Card{}, errors.New("not found")
}

func (c *cardRepo) GetAllCard(req models.GetAllCardRequest) (resp models.GetAllCardResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(c.cards) {
		resp.Cardes = []models.Card{}
		resp.Count = len(c.cards)
		return
	} else if end > len(c.cards) {
		return models.GetAllCardResponse{
			Cardes: c.cards[start:],
			Count:  len(c.cards),
		}, nil
	}

	return models.GetAllCardResponse{
		Cardes: c.cards[start:end],
		Count:  len(c.cards),
	}, nil
}

func (c *cardRepo) DeleteCard(req models.IdRequest) (resp string, err error) {
	for i, v := range c.cards {
		if v.Id == req.Id {
			if i == len(c.cards)-1 {
				c.cards = c.cards[:i]
			} else {
				c.cards = append(c.cards[:i], c.cards[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
