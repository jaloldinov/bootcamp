package models

type CreateSize struct {
	Name       string
	Price      float64
	Created_at string
}

type Size struct {
	Id         int
	Name       string
	Price      float64
	Created_at string
}

type GetAllSizeRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllSizeResponse struct {
	Sizes []Size
	Count int
}
