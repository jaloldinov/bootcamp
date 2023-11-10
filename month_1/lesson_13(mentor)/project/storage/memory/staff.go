package memory

import (
	"backend_bootcamp_17_07_2023/lesson_8/project/models"
	"errors"
)

type staffRepo struct {
	staffes []models.Staff
}

func NewStaffRepo() *staffRepo {
	return &staffRepo{staffes: make([]models.Staff, 0)}
}

func (s *staffRepo) CreateStaff(req models.CreateStaff) (int, error) {
	var id int
	if len(s.staffes) == 0 {
		id = 1
	} else {
		id = s.staffes[len(s.staffes)-1].Id + 1
	}

	s.staffes = append(s.staffes, models.Staff{
		Id:       id,
		BranchId: req.BranchId,
		TariffId: req.TariffId,
		TypeId:   req.TypeId,
		Name:     req.Name,
		Balance:  req.Balance,
	})
	return id, nil
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
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(s.staffes) {
		resp.Staffes = []models.Staff{}
		resp.Count = len(s.staffes)
		return
	} else if end > len(s.staffes) {
		return models.GetAllStaff{
			Staffes: s.staffes[start:],
			Count:   len(s.staffes),
		}, nil
	}

	return models.GetAllStaff{
		Staffes: s.staffes[start:end],
		Count:   len(s.staffes),
	}, nil
}

func (s *staffRepo) DeleteStaff(req models.IdRequest) (resp string, err error) {
	for i, v := range s.staffes {
		if v.Id == req.Id {
			if i == len(s.staffes)-1 {
				s.staffes = s.staffes[:i]
			} else {
				s.staffes = append(s.staffes[:i], s.staffes[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
