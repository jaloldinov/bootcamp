package postgres

import (
	"context"
	"database/sql"
	"example-grpc-server/pkg/helper"
	"fmt"
	"time"

	sale_service "example-grpc-server/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranch(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) CreateBranch(c context.Context, req *sale_service.CreateBranchRequest) (string, error) {
	id := uuid.NewString()
	yearNow := time.Now().Year()
	year := yearNow - int(req.FoundedAt)

	query := `
		INSERT INTO "branches"(
			"id", 
			"name", 
			"adress", 
			"year", 
			"founded_at", 
			"created_at")
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
		return "", fmt.Errorf("failed to create branch: %w", err)
	}

	return id, nil
}

func (b *branchRepo) GetBranch(c context.Context, req *sale_service.IdRequest) (resp *sale_service.Branch, err error) {
	query := `
				SELECT 
					id, 
					name, 
					adress, 
					year, 
					founded_at, 
					created_at, 
					updated_at 
				FROM branches 
				WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	branch := sale_service.Branch{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.Year,
		&branch.FoundedAt,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("branch not found")
		}
		return nil, fmt.Errorf("failed to get branch: %w", err)
	}

	branch.CreatedAt = createdAt.Format(time.RFC3339)
	if updatedAt.Valid {
		branch.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &branch, nil
}

func (b *branchRepo) UpdateBranch(c context.Context, req *sale_service.UpdateBranchRequest) (string, error) {
	yearNow := time.Now().Year()
	year := yearNow - int(req.FoundedAt)

	query := `
				UPDATE branches 
				SET 
					name = $1, 
					adress = $2, 
					year = $3, 
					founded_at = $4, 
					updated_at = NOW() 
				WHERE id = $5 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.Address,
		year,
		req.FoundedAt,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update branch: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch with ID %s not found", req.Id)
	}

	return fmt.Sprintf("branch with ID %s updated", req.Id), nil
}

func (b *branchRepo) GetAllBranch(c context.Context, req *sale_service.ListBranchRequest) (*sale_service.ListBranchResponse, error) {
	params := make(map[string]interface{})
	filter := ""
	offset := 0

	if req.Page != 0 {
		offset = int((req.Page - 1) * req.Limit)
	}

	createdAt := time.Time{}
	updatedAt := sql.NullTime{}

	s := `
		SELECT 
			id, 
			name, 
			adress, 
			year, 
			founded_at, 
			created_at, 
			updated_at 
		FROM branches`

	if req.Search != "" {
		filter += ` WHERE name ILIKE '%' || :search || '%' `
		params["search"] = req.Search
	}

	limit := fmt.Sprintf(" LIMIT %d", req.Limit)
	offsetQ := fmt.Sprintf(" OFFSET %d", offset)
	query := s + filter + " ORDER BY created_at DESC" + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := &sale_service.ListBranchResponse{}
	resp.Branches = make([]*sale_service.Branch, 0)
	count := 0
	for rows.Next() {
		var branch sale_service.Branch
		count++
		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.Year,
			&branch.FoundedAt,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		branch.CreatedAt = createdAt.Format(time.RFC3339)
		if updatedAt.Valid {
			branch.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
		}
		resp.Branches = append(resp.Branches, &branch)
	}

	resp.Count = int32(count)
	return resp, nil
}

func (b *branchRepo) DeleteBranch(c context.Context, req *sale_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM branches 
			WHERE id = $1 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)
	if err != nil {
		return "", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch with ID %s not found", req.Id)

	}

	return fmt.Sprintf("branch with ID %s deleted", req.Id), nil
}
