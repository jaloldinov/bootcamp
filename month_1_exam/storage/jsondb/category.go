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

type CategoryRepo struct {
	fileName string
	file     *os.File
}

func NewCategoryRepo(fileName string, file *os.File) *CategoryRepo {
	return &CategoryRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *CategoryRepo) Create(req *models.CreateCategory) (*models.Category, error) {

	Categorys, err := u.read()
	if err != nil {
		return nil, err
	}

	var (
		id       = uuid.New().String()
		Category = models.Category{
			Id:   id,
			Name: req.Name,
		}
	)
	Categorys[id] = Category

	err = u.write(Categorys)
	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func (u *CategoryRepo) GetById(req *models.CategoryPrimaryKey) (*models.Category, error) {

	Categorys, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Categorys[req.Id]; !ok {
		return nil, errors.New("Category not found")
	}
	Category := Categorys[req.Id]

	return &Category, nil
}

func (u *CategoryRepo) GetList(req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error) {

	var resp = &models.CategoryGetListResponse{}
	resp.Categorys = []*models.Category{}

	CategoryMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(CategoryMap)
	for _, val := range CategoryMap {
		Categorys := val
		resp.Categorys = append(resp.Categorys, &Categorys)
	}

	return resp, nil
}

func (u *CategoryRepo) Update(req *models.UpdateCategory) (*models.Category, error) {

	Categorys, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Categorys[req.Id]; !ok {
		return nil, errors.New("Category not found")
	}

	Categorys[req.Id] = models.Category{
		Id:   req.Id,
		Name: req.Name,
	}

	err = u.write(Categorys)
	if err != nil {
		return nil, err
	}
	Category := Categorys[req.Id]

	return &Category, nil
}

func (u *CategoryRepo) Delete(req *models.CategoryPrimaryKey) error {

	Categorys, err := u.read()
	if err != nil {
		return err
	}

	delete(Categorys, req.Id)

	err = u.write(Categorys)
	if err != nil {
		return err
	}

	return nil
}

func (u *CategoryRepo) read() (map[string]models.Category, error) {
	var (
		Categorys   []*models.Category
		CategoryMap = make(map[string]models.Category)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &Categorys)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, Category := range Categorys {
		CategoryMap[Category.Id] = *Category
	}

	return CategoryMap, nil
}

func (u *CategoryRepo) write(CategoryMap map[string]models.Category) error {

	var Categorys []models.Category

	for _, val := range CategoryMap {
		Categorys = append(Categorys, val)
	}

	body, err := json.MarshalIndent(Categorys, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
