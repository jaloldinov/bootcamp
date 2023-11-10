package memory

import (
	"backend_bootcamp_17_07_2023/lesson_8/project/models"
	"errors"
)

type branchRepo struct {
	branches []models.Branch
}

func NewBranchRepo() *branchRepo {
	return &branchRepo{branches: make([]models.Branch, 0)}
}

func (b *branchRepo) CreateBranch(req models.CreateBranch) (int, error) {
	var id int
	if len(b.branches) == 0 {
		id = 1
	} else {
		id = b.branches[len(b.branches)-1].Id + 1
	}

	b.branches = append(b.branches, models.Branch{
		Id:     id,
		Name:   req.Name,
		Adress: req.Adress,
	})
	return id, nil
}

func (b *branchRepo) UpdateBranch(req models.Branch) (string, error) {
	for i, v := range b.branches {
		if v.Id == req.Id {
			b.branches[i] = req
			return "updated", nil
		}
	}
	return "", errors.New("not branch found with ID")
}

func (b *branchRepo) GetBranch(req models.IdRequest) (resp models.Branch, err error) {
	for _, v := range b.branches {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Branch{}, errors.New("not found")
}

func (b *branchRepo) GetAllBranch(req models.GetAllBranchRequest) (resp models.GetAllBranch, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	if start > len(b.branches) {
		resp.Branches = []models.Branch{}
		resp.Count = len(b.branches)
		return
	} else if end > len(b.branches) {
		return models.GetAllBranch{
			Branches: b.branches[start:],
			Count:    len(b.branches),
		}, nil
	}

	return models.GetAllBranch{
		Branches: b.branches[start:end],
		Count:    len(b.branches),
	}, nil
}

func (b *branchRepo) DeleteBranch(req models.IdRequest) (resp string, err error) {
	for i, v := range b.branches {
		if v.Id == req.Id {
			if i == len(b.branches)-1 {
				b.branches = b.branches[:i]
			} else {
				b.branches = append(b.branches[:i], b.branches[i+1:]...)
				return "deleted", nil
			}
		}
	}
	return "", errors.New("not found")
}
