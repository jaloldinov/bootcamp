package memory

import (
	"errors"
	"fmt"
	"lesson_15/models"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type branchRepo struct {
	branches []models.Branch
}

func NewBranchRepo() *branchRepo {
	return &branchRepo{branches: make([]models.Branch, 0)}
}

func (b *branchRepo) CreateBranch(req models.CreateBranch) (string, error) {
	id := uuid.New()
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	b.branches = append(b.branches, models.Branch{
		Id:        id.String(),
		Name:      req.Name,
		Adress:    req.Adress,
		FoundedAt: req.FoundedAt,
		CreatedAt: createdAt,
	})
	return id.String(), nil
}

func (b *branchRepo) UpdateBranch(req models.Branch) (string, error) {
	for i, v := range b.branches {
		if v.Id == req.Id {
			b.branches[i] = req
			fmt.Println(b.branches)
			return "updated", nil
		}
	}
	return "", errors.New("not branch found with ID")
}

func (b *branchRepo) GetBranch(req models.IdRequest) (resp models.Branch, err error) {
	for _, v := range b.branches {
		if v.Id == req.Id {
			foundedAt, _ := strconv.Atoi(v.FoundedAt[:4])
			v.Year = time.Now().Year() - foundedAt
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
			b.branches = append(b.branches[:i], b.branches[i+1:]...)
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}
