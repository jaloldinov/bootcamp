package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
)

type BranchRepo struct {
	fileName string
	file     *os.File
}

func NewBranchRepo(fileName string, file *os.File) *BranchRepo {
	return &BranchRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *BranchRepo) Create(req *models.CreateBranch) (*models.BranchPrimaryKey, error) {
	branches, err := u.read()
	if err != nil {
		return nil, err
	}
	var (
		id     = uuid.New().String()
		branch = models.Branch{
			Id:   id,
			Name: req.Name,
		}
	)
	branches[id] = branch
	err = u.write(branches)
	if err != nil {
		return nil, err
	}
	return &models.BranchPrimaryKey{Id: id}, nil
}

func (u *BranchRepo) GetById(req *models.BranchPrimaryKey) (*models.Branch, error) {

	Branches, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Branches[req.Id]; !ok {
		return nil, errors.New("Branch not found")
	}
	branch := Branches[req.Id]

	return &branch, nil
}

func (u *BranchRepo) GetList(req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {
	var resp = &models.BranchGetListResponse{}
	resp.Branches = []*models.Branch{}

	BranchesMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(BranchesMap)
	for _, val := range BranchesMap {
		Branches := val
		resp.Branches = append(resp.Branches, &Branches)
	}

	return resp, nil
}

func (u *BranchRepo) Update(req *models.Branch) (*models.Branch, error) {
	Branches, err := u.read()
	if err != nil {
		return nil, err
	}
	if _, ok := Branches[req.Id]; !ok {
		return nil, errors.New("Branch not found")
	}
	Branches[req.Id] = models.Branch{
		Id:   req.Id,
		Name: req.Name,
	}
	err = u.write(Branches)
	if err != nil {
		return nil, err
	}
	Branch := Branches[req.Id]
	return &Branch, nil
}

func (u *BranchRepo) Delete(req *models.BranchPrimaryKey) (string, error) {
	Branches, err := u.read()
	if err != nil {
		return "", err
	}
	delete(Branches, req.Id)
	err = u.write(Branches)
	if err != nil {
		return "", err
	}
	return "deleted", nil
}

func (u *BranchRepo) read() (map[string]models.Branch, error) {
	var (
		branches   []*models.Branch
		productMap = make(map[string]models.Branch)
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, product := range branches {
		productMap[product.Id] = *product
	}

	return productMap, nil
}

func (u *BranchRepo) write(productMap map[string]models.Branch) error {

	var branches []models.Branch

	for _, val := range productMap {
		branches = append(branches, val)
	}

	body, err := json.MarshalIndent(branches, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
