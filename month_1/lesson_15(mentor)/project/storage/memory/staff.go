package memory

import (
	"errors"
	"lesson_15/models"
	"strings"

	"github.com/google/uuid"
)

type staffRepo struct {
	staffes []models.Staff
}

func NewStaffRepo() *staffRepo {
	return &staffRepo{staffes: make([]models.Staff, 0)}
}

func (s *staffRepo) CreateStaff(req models.CreateStaff) (string, error) {
	id := uuid.New()

	s.staffes = append(s.staffes, models.Staff{
		Id:       id.String(),
		BranchId: req.BranchId,
		TariffId: req.TariffId,
		TypeId:   req.TypeId,
		Name:     req.Name,
		Balance:  req.Balance,
	})
	return id.String(), nil
}

func (s *staffRepo) UpdateStaff(req models.Staff) (string, error) {
	for i, v := range s.staffes {
		if v.Id == req.Id {
			s.staffes[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not staff found with ID")
}

func (s *staffRepo) GetStaff(req models.IdRequest) (models.Staff, error) {
	for _, v := range s.staffes {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Staff{}, errors.New("not staff found with ID")
}

func (s *staffRepo) GetAllStaff(req models.GetAllStaffRequest) (resp models.GetAllStaff, err error) {
	var filtered []models.Staff
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	for _, v := range s.staffes {
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
	for i, v := range s.staffes {
		if v.Id == req.Id {
			s.staffes = append(s.staffes[:i], s.staffes[i+1:]...)
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}
