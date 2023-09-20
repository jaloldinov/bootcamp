package models

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type Order struct {
	Id         string             `json:"id"`
	UserId     string             `json:"user_id"`
	Sum        int                `json:"sum"`
	SumCount   int                `json:"sum_count"`
	DateTime   string             `json:"date_time"`
	Status     string             `json:"status"`
	OrderItems []*CreateOrderItem `json:"order_items"`
}
type Task2request struct {
	FromDate string
	ToDate   string
}

type OrderGetList struct {
	Count  int
	Orders []*Order
}
type CreateOrder struct {
	UserId   string `json:"user_id"`
	Sum      int    `json:"sum"`
	SumCount int    `json:"sum_count"`
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
}

type OrderGetListRequest struct {
	Offset int
	Limit  int
}

type UpdateOrder struct {
	Id         string             `json:"id"`
	UserId     string             `json:"user_id"`
	Sum        int                `json:"sum"`
	SumCount   int                `json:"sum_count"`
	DateTime   string             `json:"date_time"`
	Status     string             `json:"status"`
	OrderItems []*CreateOrderItem `json:"order_items"`
}

type CreateOrderItem struct {
	Id         string `json:"id"`
	ProductId  string `json:"product_id"`
	OrderId    string `json:"order_id"`
	Count      int    `json:"count"`
	TotalPrice int    `json:"total_price"`
}

type RemoveOrderItemPrimaryKey struct {
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
}

type OrderPayment struct {
	OrderId string `json:"order_id"`
}
