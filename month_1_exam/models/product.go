package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type Product struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type UpdateProduct struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type ProductGetListRequest struct {
	Offset int
	Limit  int
}

type ProductGetListResponse struct {
	Count    int
	Products []*Product
}

type NameCount struct {
	Name  string
	Count int
}
