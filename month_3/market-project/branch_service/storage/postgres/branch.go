package postgres

import (
	"branch_service/pkg/helper"
	"context"
	"fmt"
	"time"

	branch_service "branch_service/genproto"

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

func (b *branchRepo) CreateBranch(c context.Context, req *branch_service.CreateBranchRequest) (string, error) {
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

func (b *branchRepo) GetBranch(c context.Context, req *branch_service.IdRequest) (resp *branch_service.Branch, err error) {
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
		updatedAt time.Time
	)

	branch := branch_service.Branch{}
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
	branch.UpdatedAt = createdAt.Format(time.RFC3339)

	return &branch, nil
}

func (b *branchRepo) UpdateBranch(c context.Context, req *branch_service.UpdateBranchRequest) (string, error) {
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

func (b *branchRepo) GetAllBranch(c context.Context, req *branch_service.ListBranchRequest) (*branch_service.ListBranchResponse, error) {
	var (
		resp   branch_service.ListBranchResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM branches WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

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
		WHERE true` + filter

	query += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(c, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	defer rows.Close()

	createdAt := time.Time{}
	updatedAt := time.Time{}
	for rows.Next() {
		var branch branch_service.Branch

		err = rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.Year,
			&branch.FoundedAt,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning branch err: %w", err)
		}

		branch.CreatedAt = createdAt.Format(time.RFC3339)
		branch.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.Branches = append(resp.Branches, &branch)
	}

	return &resp, nil
}

func (b *branchRepo) DeleteBranch(c context.Context, req *branch_service.IdRequest) (resp string, err error) {
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
