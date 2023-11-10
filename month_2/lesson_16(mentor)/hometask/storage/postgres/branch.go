package postgres

import (
	"app/models"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) CreateBranch(req *models.CreateBranch) (string, error) {
	id := uuid.NewString()
	yearNow := time.Now().Year()
	year := yearNow - req.FoundedAt

	query := `
		INSERT INTO branch(id, name, address, year, founded_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		year,
		req.FoundedAt,
		time.Now(),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *branchRepo) GetBranch(req *models.IdRequest) (resp *models.Branch, err error) {
	// branches, err := b.read()
	// if err != nil {
	// 	return &models.Branch{}, err
	// }
	// for _, v := range branches {
	// 	if v.Id == req.Id {
	// 		foundedAt, _ := strconv.Atoi(v.FoundedAt[:4])
	// 		v.Year = time.Now().Year() - foundedAt
	// 		return v, nil
	// 	}
	// }

	return &models.Branch{}, errors.New("not found")
}

/*
func (b *branchRepo) UpdateBranch(req models.Branch) (string, error) {
	branches, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range branches {
		if v.Id == req.Id {
			branches[i] = req
			branches[i].CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			err = b.write(branches)
			if err != nil {
				return "", err
			}
			return v.Id, nil
		}
	}

	return "", errors.New("not branch found with ID")
}

func (b *branchRepo) GetAllBranch(req models.GetAllBranchRequest) (resp models.GetAllBranch, err error) {
	branches, err := b.read()
	if err != nil {
		return models.GetAllBranch{}, err
	}

	var filtered []models.Branch

	for _, v := range branches {
		if req.Name == "" || strings.Contains(v.Name, req.Name) || strings.Contains(v.Adress, req.Name) {
			filtered = append(filtered, v)
		}
	}

	if req.Limit == 0 {
		req.Limit = len(filtered)
	}

	start := req.Page * req.Limit
	end := (req.Page + 1) * req.Limit

	if start > len(filtered) {
		resp.Branches = []models.Branch{}
		resp.Count = len(filtered)
		return
	} else if end > len(filtered) {
		return models.GetAllBranch{
			Branches: filtered[start:],
			Count:    len(filtered),
		}, nil
	}

	return models.GetAllBranch{
		Branches: filtered[start:end],
		Count:    len(filtered),
	}, nil
}

func (b *branchRepo) DeleteBranch(req models.IdRequest) (resp string, err error) {
	branches, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range branches {
		if v.Id == req.Id {
			branches = append(branches[:i], branches[i+1:]...)
			err = b.write(branches)
			if err != nil {
				return "", err
			}
			return "deleted", nil
		}
	}

	return "", errors.New("not found")
}

func (u *branchRepo) read() ([]models.Branch, error) {
	var (
		branches []models.Branch
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

	return branches, nil
}

func (u *branchRepo) write(branches []models.Branch) error {

	body, err := json.Marshal(branches)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
*/
