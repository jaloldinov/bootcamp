package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Category_id string `json:"category_id"`
}

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Category_id string `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateProduct struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Category_id string `json:"category_id"`
}

type ProductGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ProductGetListResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
