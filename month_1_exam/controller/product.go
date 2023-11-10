package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) ProductCreate(req *models.CreateProduct) (*models.Product, error) {

	log.Printf("User create req: %+v\n", req)

	resp, err := c.Strg.Product().Create(req)
	if err != nil {
		log.Printf("error while user Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}
	return resp, nil
}

func (c *Controller) GetByIdPoduct(req *models.ProductPrimaryKey) (*models.Product, error) {

	resp, err := c.Strg.Product().GetById(req)
	if err != nil {
		log.Printf("error while user GetById: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ProductGetList(req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {

	resp, err := c.Strg.Product().GetList(req)
	if err != nil {
		log.Printf("error while user GetList: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ProductUpdate(req *models.UpdateProduct) (*models.Product, error) {

	resp, err := c.Strg.Product().Update(req)
	if err != nil {
		log.Printf("error while user Update: %+v\n", err)
		return nil, err
	}
	return resp, nil
}

func (c *Controller) ProductDelete(req *models.ProductPrimaryKey) error {

	err := c.Strg.Product().Delete(req)
	if err != nil {
		log.Printf("error while user Delete: %+v\n", err)
		return err
	}

	return nil
}
