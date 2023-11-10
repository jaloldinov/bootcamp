package memory

import (
	"app/models"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type staffRepo struct {
	fileName string
}

func NewStaffRepo(fileName string) *staffRepo {
	return &staffRepo{fileName: fileName}
}

func (s *staffRepo) CreateStaff(req models.CreateStaff) (string, error) {
	staffes, err := s.read()
	if err != nil {
		return "", err
	}
	id := uuid.NewString()
	staffes = append(staffes, models.Staff{
		Id:       id,
		BranchId: req.BranchId,
		TariffId: req.TariffId,
		TypeId:   req.TypeId,
		Name:     req.Name,
		Balance:  req.Balance,
	})
	err = s.write(staffes)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *staffRepo) UpdateStaff(req models.Staff) (string, error) {
	staffes, err := s.read()
	if err != nil {
		return "", err
	}
	for i, v := range staffes {
		if v.Id == req.Id {
			staffes[i] = req
			err = s.write(staffes)
			if err != nil {
				return "", err
			}
			return "updated", nil
		}
	}
	return "", errors.New("not staff found with ID")
}

func (s *staffRepo) GetStaff(req models.IdRequest) (models.Staff, error) {
	staffes, err := s.read()
	if err != nil {
		return models.Staff{}, err
	}
	for _, v := range staffes {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Staff{}, errors.New("not staff found with ID")
}

func (u *staffRepo) GetByLogin(req models.LoginRequest) (models.Staff, error) {
	staffes := []models.Staff{}
	for _, s := range staffes {
		if req.Login == s.Login {
			return s, nil
		}
	}
	return models.Staff{}, nil
}

func (s *staffRepo) GetAllStaff(req models.GetAllStaffRequest) (resp models.GetAllStaff, err error) {
	staffes, err := s.read()
	if err != nil {
		return models.GetAllStaff{}, err
	}
	var filtered []models.Staff
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	for _, v := range staffes {
		if strings.Contains(v.Name, req.Name) || v.TypeId == req.Type && req.BalanceFrom >= v.Balance && req.BalanceTo <= v.Balance {
			filtered = append(filtered, v)
		}
	}

	if start > len(filtered) {
		resp.Staffes = []models.Staff{}
		resp.Count = len(filtered)
		return
	} else if end > len(filtered) {
		return models.GetAllStaff{
			Staffes: filtered[start:],
			Count:   len(filtered),
		}, nil
	}

	return models.GetAllStaff{
		Staffes: filtered[start:end],
		Count:   len(filtered),
	}, nil

}

func (s *staffRepo) DeleteStaff(req models.IdRequest) (resp string, err error) {
	staffes, err := s.read()
	if err != nil {
		return "", err
	}

	for i, v := range staffes {
		if v.Id == req.Id {
			staffes = append(staffes[:i], staffes[i+1:]...)
			err = s.write(staffes)
			if err != nil {
				return "", err
			}
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}

func (u *staffRepo) ChangeBalance(req models.ChangeBalance) (string, error) {
	staffes, err := u.read()
	if err != nil {
		return "", err
	}
	for i, v := range staffes {
		if v.Id == req.Id {
			staffes[i].Balance = req.Balance
			err = u.write(staffes)
			if err != nil {
				return "", err
			}
			return "updated balance", nil
		}
	}
	return "", errors.New("not staff balance found with ID")
}

func (u *staffRepo) Exists(req models.ExistsReq) bool {
	staffes := []models.Staff{}
	for _, s := range staffes {
		if req.Phone == s.Phone {
			return true
		}
	}
	return false
}

func (u *staffRepo) read() ([]models.Staff, error) {
	var staffes []models.Staff

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &staffes)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return staffes, nil
}

func (u *staffRepo) write(staffes []models.Staff) error {

	body, err := json.Marshal(staffes)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
