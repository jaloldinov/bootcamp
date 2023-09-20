package controller

import (
	"errors"
	"fmt"
	"log"
	"math"

	"app/config"
	"app/models"
	"app/pkg/convert"
)

func (c *Controller) OrderCreate(req *models.CreateOrder) (*models.Order, error) {

	log.Printf("User create req: %+v\n", req)

	req.Status = config.OrderStatus["0"]
	resp, err := c.Strg.Order().Create(req)
	if err != nil {
		log.Printf("error while order Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) GetByIdOrder(req *models.OrderPrimaryKey) (*models.Order, error) {

	resp, err := c.Strg.Order().GetById(req)
	if err != nil {
		log.Printf("error while order GetById: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) OrderGetList(req *models.OrderGetListRequest) (*models.OrderGetList, error) {

	resp, err := c.Strg.Order().GetList(req)
	if err != nil {
		log.Printf("error while order GetList: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) OrderUpdate(req *models.UpdateOrder) (*models.Order, error) {

	resp, err := c.Strg.Order().Update(req)
	if err != nil {
		log.Printf("error while order Update: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) OrderDelete(req *models.OrderPrimaryKey) error {

	err := c.Strg.Order().Delete(req)
	if err != nil {
		log.Printf("error while order Delete: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) AddOrderItem(req *models.CreateOrderItem) error {

	product, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		log.Printf("error while Product => GetById: %+v\n", err)
		return err
	}

	productPrice := product.Price
	if product.DiscountType == config.PercentDiscountType {
		productPrice = int(float64(product.Price) * ((100 - float64(product.Discount)) / 100))
	} else if product.DiscountType == config.FixDiscountType {
		productPrice = product.Price - product.Discount
	}

	req.TotalPrice = req.Count * productPrice
	err = c.Strg.Order().AddOrderItem(req)
	if err != nil {
		log.Printf("error while order => AddOrderItem: %+v\n", err)
		return err
	}

	order, err := c.Strg.Order().GetById(&models.OrderPrimaryKey{Id: req.OrderId})
	if err != nil {
		log.Printf("error while Order => GetById: %+v\n", err)
		return err
	}

	order.Sum += req.TotalPrice
	order.SumCount += req.Count

	var updateOrder models.UpdateOrder
	err = convert.ConvertStructToStruct(order, &updateOrder)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.Order().Update(&updateOrder)
	if err != nil {
		log.Printf("error while order => Update: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) RemoveOrderItem(req *models.RemoveOrderItemPrimaryKey) error {

	err := c.Strg.Order().RemoveOrderItem(req)
	if err != nil {
		log.Printf("error while order => RemoveOrderItem: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) OrderPayment(req *models.OrderPayment) error {

	order, err := c.Strg.Order().GetById(&models.OrderPrimaryKey{Id: req.OrderId})
	if err != nil {
		log.Printf("error while Order => GetById: %+v\n", err)
		return err
	}

	user, err := c.Strg.User().GetById(&models.UserPrimaryKey{Id: order.UserId})
	if err != nil {
		log.Printf("error while User => GetById: %+v\n", err)
		return err
	}

	if order.Sum > user.Balance {
		return errors.New("Not enough balance " + user.Name + " " + user.Surname)
	}
	// =========================================
	if order.SumCount > 9 {
		minPrice := math.MaxInt64
		for _, item := range order.OrderItems {
			if (item.TotalPrice / item.Count) < minPrice {
				minPrice = item.TotalPrice / item.Count
			}
		}
		order.Sum -= minPrice
	}
	// =========================================

	order.Status = config.OrderStatus["1"]
	user.Balance -= order.Sum
	fmt.Println(order.Status)
	var updateOrder models.UpdateOrder
	err = convert.ConvertStructToStruct(order, &updateOrder)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.Order().Update(&updateOrder)
	if err != nil {
		log.Printf("error while order => Update: %+v\n", err)
		return err
	}

	var updateUser models.UpdateUser
	err = convert.ConvertStructToStruct(user, &updateUser)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.User().Update(&updateUser)
	if err != nil {
		log.Printf("error while User => Update: %+v\n", err)
		return err
	}

	return nil
}
