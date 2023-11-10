package models

type CreateCard struct {
	Name       string
	Quantity   int
	Product_id int
	Created_at string
}

type Card struct {
	Id         int
	Name       string
	Quantity   int
	Product_id int
	Created_at string
}

type GetAllCardRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllCardResponse struct {
	Cardes []Card
	Count  int
}
