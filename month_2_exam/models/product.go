package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    int     `json:"barcode"`
	CateogryId string  `json:"cateogry_id"`
}

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    int     `json:"barcode"`
	CateogryId string  `json:"cateogry_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type UpdateProduct struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    int     `json:"barcode"`
	CateogryId string  `json:"cateogry_id"`
}

type ProductGetListRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ProductGetListResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
