package postgres

import (
	"context"
	"fmt"
	"product_service/pkg/helper"
	"time"

	pb "product_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategory(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (b *categoryRepo) CreateCategory(c context.Context, req *pb.CreateCategoryRequest) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "category"(
			"id", 
			"name",  
			"created_at")
		VALUES ($1, $2, NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create category: %w", err)
	}

	return id, nil
}

func (b *categoryRepo) GetCategory(c context.Context, req *pb.IdRequest) (resp *pb.Category, err error) {
	query := `
			SELECT 
				"id", 
				"name", 
				"created_at",
				"updated_at" 
			FROM "category" 
			WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	category := pb.Category{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("category not found")
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	category.CreatedAt = createdAt.Format(time.RFC3339)
	category.UpdatedAt = createdAt.Format(time.RFC3339)

	return &category, nil
}

func (b *categoryRepo) GetAllCategory(c context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {
	var (
		resp   pb.ListCategoryResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM "category" WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `
			SELECT 
				"id", 
				"name",
				"created_at",
				"updated_at"
			FROM "category" 
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
		var category pb.Category

		err = rows.Scan(
			&category.Id,
			&category.Name,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning category err: %w", err)
		}

		category.CreatedAt = createdAt.Format(time.RFC3339)
		category.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.Categoryes = append(resp.Categoryes, &category)
	}

	return &resp, nil
}

func (b *categoryRepo) UpdateCategory(c context.Context, req *pb.UpdateCategoryRequest) (string, error) {

	query := `
				UPDATE "category" 
				SET 
				"name" = $1, 
				"updated_at" = NOW() 
				WHERE id = $2 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("category with ID %s not found", req.Id)
	}

	return fmt.Sprintf("category with ID %s updated", req.Id), nil
}

func (b *categoryRepo) DeleteCategory(c context.Context, req *pb.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM "category" 
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
		return "", fmt.Errorf("category with ID %s not found", req.Id)

	}

	return fmt.Sprintf("category with ID %s deleted", req.Id), nil
}
