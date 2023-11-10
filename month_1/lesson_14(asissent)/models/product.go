package models

type CreateProduct struct {
	Name       string
	Card_Id    int
	Size_Id    int
	Created_at string
}

type Product struct {
	Id         int
	Name       string
	Card_Id    int
	Size_Id    int
	Created_at string
}

type GetAllProductRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllProductResponse struct {
	Productes []Product
	Count     int
}
