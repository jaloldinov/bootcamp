package memory

import (
	"encoding/json"
	"errors"
	"lesson_20/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type staffTarifRepo struct {
	fileName string
}

func NewStaffTarifRepo(fileName string) *staffTarifRepo {
	return &staffTarifRepo{fileName: fileName}
}

func (s *staffTarifRepo) CreateStaffTarif(req models.CreateStaffTarif) (string, error) {
	staffTarifs, err := s.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	staffTarifs = append(staffTarifs, models.StaffTarif{
		Id:            id,
		Name:          req.Name,
		Type:          req.Type,
		AmountForCash: req.AmountForCash,
		AmountForCard: req.AmountForCard,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	})

	err = s.write(staffTarifs)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *staffTarifRepo) UpdateStaffTarif(req models.StaffTarif) (string, error) {
	staffTarifs, err := s.read()
	if err != nil {
		return "", err
	}

	for i, v := range staffTarifs {
		if v.Id == req.Id {
			staffTarifs[i].Name = req.Name
			staffTarifs[i].Type = req.Type
			staffTarifs[i].AmountForCash = req.AmountForCash
			staffTarifs[i].AmountForCard = req.AmountForCard
			err = s.write(staffTarifs)
			if err != nil {
				return "", err
			}
			return "updated successfully", nil
		}
	}
	return "", errors.New("not found")
}

func (s *staffTarifRepo) GetStaffTarif(req models.IdRequest) (resp models.StaffTarif, err error) {
	staffTarifs, err := s.read()
	if err != nil {
		return resp, err
	}

	for _, v := range staffTarifs {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return resp, errors.New("not found")
}

func (s *staffTarifRepo) GetAllStaffTarif(req models.GetAllStaffTarifRequest) (resp models.GetAllStaffTarif, err error) {
	staffTarifs, err := s.read()
	if err != nil {
		return resp, err
	}

	filtered := []models.StaffTarif{}
	for _, v := range staffTarifs {
		if strings.Contains(v.Name, req.Name) {
			filtered = append(filtered, v)
		}
	}

	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(filtered) {
		resp.StaffTarifs = []models.StaffTarif{}
		resp.Count = len(filtered)
		return resp, nil
	} else if end > len(filtered) {
		return models.GetAllStaffTarif{
			StaffTarifs: filtered[start:],
			Count:       len(filtered),
		}, nil
	}

	return models.GetAllStaffTarif{
		StaffTarifs: filtered[start:end],
		Count:       len(filtered)}, nil
}

func (s *staffTarifRepo) DeleteStaffTarif(req models.IdRequest) (string, error) {
	staffTarifs, err := s.read()
	if err != nil {
		return "", err
	}
	for i, v := range staffTarifs {
		if v.Id == req.Id {
			staffTarifs = append(staffTarifs[:i], staffTarifs[i+1:]...)
			err = s.write(staffTarifs)
			if err != nil {
				return "", err
			}
			return "deleted", nil
		}
	}
	return "", errors.New("not found id")
}

func (u *staffTarifRepo) read() ([]models.StaffTarif, error) {
	var staffTarifs []models.StaffTarif

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &staffTarifs)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return staffTarifs, nil
}

func (u *staffTarifRepo) write(staffTarifs []models.StaffTarif) error {
	body, err := json.Marshal(staffTarifs)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
