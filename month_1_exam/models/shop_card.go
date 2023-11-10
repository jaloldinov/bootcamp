package models

type ShopCard struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}
type Task8Model struct {
	Name  string
	Time  string
	Count int
}
