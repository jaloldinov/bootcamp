package models

type CreateClient struct {
	Name       string
	Card_Id    int
	Created_at string
}

type Client struct {
	Id         int
	Name       string
	Card_Id    int
	Created_at string
}

type GetAllClientRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllClientResponse struct {
	Clientes []Client
	Count    int
}
