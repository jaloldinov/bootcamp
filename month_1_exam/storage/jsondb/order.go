package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"app/models"
	"app/pkg/convert"
	"app/pkg/file"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type OrderRepo struct {
	fileName string
	file     *os.File
}

func NewOrderRepo(fileName string, file *os.File) *OrderRepo {
	return &OrderRepo{
		fileName: fileName,
		file:     file,
	}
}
func (o *OrderRepo) Create(ord *models.CreateOrder) (*models.Order, error) {
	orders, err := file.Read(o.fileName)
	if err != nil {
		return nil, err
	}
	var (
		id    = uuid.New().String()
		order = models.Order{
			Id:       id,
			UserId:   ord.UserId,
			Sum:      ord.Sum,
			SumCount: ord.SumCount,
			DateTime: time.Now().Format("2006-01-02 15:04:05"),
			Status:   ord.Status,
		}
	)
	orders[id] = order
	err = file.Write(o.fileName, orders)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepo) GetById(ord *models.OrderPrimaryKey) (*models.Order, error) {

	var resp models.Order
	orders, err := file.Read(o.fileName)
	if err != nil {
		return nil, err
	}

	if _, ok := orders[ord.Id]; !ok {
		return nil, errors.New("not found")
	}

	err = convert.ConvertStructToStruct(orders[ord.Id], &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (o *OrderRepo) GetList(ord *models.OrderGetListRequest) (*models.OrderGetList, error) {
	var resp = &models.OrderGetList{}
	resp.Orders = []*models.Order{}

	orderMap, err := o.read()
	if err != nil {
		return nil, err
	}
	resp.Count = len(orderMap)
	for _, val := range orderMap {

		order := val
		resp.Orders = append(resp.Orders, &order)

	}
	return resp, nil
}

func (o *OrderRepo) Update(ord *models.UpdateOrder) (*models.Order, error) {

	var resp models.Order
	orders, err := file.Read(o.fileName)
	if err != nil {
		return nil, err
	}

	if _, ok := orders[ord.Id]; !ok {
		return nil, errors.New("Order not found!")
	}

	orders[ord.Id] = models.Order{
		Id:         ord.Id,
		UserId:     ord.UserId,
		Sum:        ord.Sum,
		SumCount:   ord.SumCount,
		DateTime:   ord.DateTime,
		Status:     ord.Status,
		OrderItems: ord.OrderItems,
	}

	err = file.Write(o.fileName, orders)
	if err != nil {
		return nil, err
	}

	err = convert.ConvertStructToStruct(orders[ord.Id], &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (o *OrderRepo) Delete(ord *models.OrderPrimaryKey) error {
	orders, err := file.Read(o.fileName)
	if err != nil {
		return err
	}
	delete(orders, ord.Id)
	err = file.Write(o.fileName, orders)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepo) AddOrderItem(req *models.CreateOrderItem) error {

	orders, err := file.Read(o.fileName)
	if err != nil {
		return err
	}

	var (
		order      = cast.ToStringMap(orders[req.OrderId])
		orderItems = cast.ToSlice(order["order_items"])
	)

	req.Id = uuid.New().String()
	orderItems = append(orderItems, req)

	order["order_items"] = orderItems
	orders[req.OrderId] = order

	err = file.Write(o.fileName, orders)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepo) RemoveOrderItem(req *models.RemoveOrderItemPrimaryKey) error {

	orders, err := file.Read(o.fileName)
	if err != nil {
		return err
	}

	var (
		order      = cast.ToStringMap(orders[req.OrderId])
		orderItems = cast.ToSlice(order["order_items"])
		removeInd  = 0
	)

	for index, orderItem := range orderItems {
		if cast.ToStringMap(orderItem)["id"] == req.Id {
			removeInd = index
			break
		}
	}

	orderItems = append(orderItems[:removeInd], orderItems[removeInd+1:]...)

	order["order_items"] = orderItems
	orders[req.OrderId] = order

	err = file.Write(o.fileName, orders)
	if err != nil {
		return err
	}

	return nil
}

func (u *OrderRepo) read() (map[string]models.Order, error) {
	var (
		orders   []*models.Order
		orderMap = make(map[string]models.Order)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &orders)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, order := range orders {
		orderMap[order.Id] = *order
	}

	return orderMap, nil
}
