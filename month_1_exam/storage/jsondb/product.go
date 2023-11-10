package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"

	"app/models"
)

type ProductRepo struct {
	fileName string
	file     *os.File
}

func NewProductRepo(fileName string, file *os.File) *ProductRepo {
	return &ProductRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *ProductRepo) Create(req *models.CreateProduct) (*models.Product, error) {

	products, err := u.read()
	if err != nil {
		return nil, err
	}

	var (
		id      = uuid.New().String()
		product = models.Product{
			Id:           id,
			Name:         req.Name,
			Price:        req.Price,
			Discount:     req.Discount,
			DiscountType: req.DiscountType,
			CategoryID:   req.CategoryID,
		}
	)
	products[id] = product
	err = u.write(products)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (u *ProductRepo) GetById(req *models.ProductPrimaryKey) (*models.Product, error) {

	products, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := products[req.Id]; !ok {
		return nil, errors.New("user not found")
	}
	product := products[req.Id]

	return &product, nil
}

func (u *ProductRepo) GetList(req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {

	var resp = &models.ProductGetListResponse{}
	resp.Products = []*models.Product{}

	productMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(productMap)
	for _, val := range productMap {
		users := val
		resp.Products = append(resp.Products, &users)
	}

	return resp, nil
}

func (u *ProductRepo) Update(req *models.UpdateProduct) (*models.Product, error) {

	products, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := products[req.Id]; !ok {
		return nil, errors.New("user not found")
	}

	products[req.Id] = models.Product{
		Id:           req.Id,
		Name:         req.Name,
		Price:        req.Price,
		Discount:     req.Discount,
		DiscountType: req.DiscountType,
		CategoryID:   req.CategoryID,
	}

	err = u.write(products)
	if err != nil {
		return nil, err
	}
	product := products[req.Id]

	return &product, nil
}

func (u *ProductRepo) Delete(req *models.ProductPrimaryKey) error {

	products, err := u.read()
	if err != nil {
		return err
	}

	delete(products, req.Id)

	err = u.write(products)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductRepo) read() (map[string]models.Product, error) {
	var (
		products   []*models.Product
		productMap = make(map[string]models.Product)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, product := range products {
		productMap[product.Id] = *product
	}

	return productMap, nil
}

func (u *ProductRepo) write(productMap map[string]models.Product) error {

	var products []models.Product

	for _, val := range productMap {
		products = append(products, val)
	}

	body, err := json.MarshalIndent(products, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
