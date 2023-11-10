package postgres

import (
	"catalog_service/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	catalog_service "catalog_service/genproto"

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

func (b *categoryRepo) Create(c context.Context, req *catalog_service.CreateCategoryRequest) (string, error) {
	if req.ParentId != 0 {
		query := `
		INSERT INTO "categories"(
			"title", 
			"image", 
			"parent_id",    
			"order_number", 
			"created_at")
		VALUES ($1, $2, $3, $4, NOW()) RETURNING "id"
	`

		var id int
		err := b.db.QueryRow(c, query,
			req.Title,
			req.Image,
			req.ParentId,
			req.OrderNumber,
		).Scan(&id)
		if err != nil {
			return "", fmt.Errorf("failed to create parent category: %w", err)
		}

		return fmt.Sprintf("created parent category with ID: %d", id), nil
		// 		CREATE PARENT CATEGORY
	} else {
		query := `
		INSERT INTO "categories"(
			"title", 
			"image",   
			"order_number", 
			"created_at")
		VALUES ($1, $2, $3, NOW()) RETURNING "id"
	`

		var id int
		err := b.db.QueryRow(c, query,
			req.Title,
			req.Image,
			req.OrderNumber,
		).Scan(&id)
		if err != nil {
			return "", fmt.Errorf("failed to create category: %w", err)
		}

		return fmt.Sprintf("created category with ID: %d", id), nil
	}

}

func (b *categoryRepo) Get(c context.Context, req *catalog_service.IdRequest) (resp *catalog_service.Category, err error) {
	query := `
		SELECT 
			"id",
			"title", 
			"image", 
			"active", 
			"parent_id",  
			"order_number",
			"created_at",
			"updated_at" 
		FROM "categories" 
		WHERE id=$1`

	var (
		createdAt sql.NullString
		updatedAt sql.NullString
		parentId  sql.NullInt32
	)

	category := catalog_service.Category{}
	err = b.db.QueryRow(c, query, req.Id).Scan(
		&category.Id,
		&category.Title,
		&category.Image,
		&category.Active,
		&parentId,
		&category.OrderNumber,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("category not found")
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	if parentId.Valid {
		category.ParentId = parentId.Int32
	}
	if createdAt.Valid {
		category.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		category.UpdatedAt = updatedAt.String
	}

	return &category, nil
}

func (b *categoryRepo) GetList(c context.Context, req *catalog_service.ListCategoryRequest) (*catalog_service.ListCategoryResponse, error) {
	var (
		resp   catalog_service.ListCategoryResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND title ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM "categories" WHERE "active" ` + filter

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
			"title", 
			"image", 
			"active", 
			"parent_id",  
			"order_number",
			"created_at",
			"updated_at" 
			FROM "categories" 
		    WHERE "active" ` + filter

	query += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"

	params["limit"] = 10
	params["offset"] = 0

	if req.Limit > 0 {
		params["limit"] = req.Limit
	}
	if req.Page >= 0 {
		params["offset"] = (req.Page - 1) * req.Limit
	}

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(c, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	defer rows.Close()

	var createdAt sql.NullString
	var updatedAt sql.NullString
	var parentId sql.NullInt32

	for rows.Next() {
		var category catalog_service.Category

		err = rows.Scan(
			&category.Id,
			&category.Title,
			&category.Image,
			&category.Active,
			&parentId,
			&category.OrderNumber,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning category err: %w", err)
		}
		if parentId.Valid {
			category.ParentId = parentId.Int32
		}
		if createdAt.Valid {
			category.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			category.UpdatedAt = createdAt.String
		}
		resp.Categories = append(resp.Categories, &category)
	}

	return &resp, nil
}

func (b *categoryRepo) Update(c context.Context, req *catalog_service.UpdateCategoryRequest) (string, error) {

	query := `
				UPDATE "categories" 
				SET 
				"title" = $1,
				"image" = $2,
				"parent_id" = $3,  
				"order_number" = $4,
				"updated_at" = NOW() 
				WHERE id = $5 AND "active" `

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Title,
		req.Image,
		req.ParentId,
		req.OrderNumber,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("category with ID %d not found", req.Id)
	}

	return fmt.Sprintf("category with ID %d updated", req.Id), nil
}

func (b *categoryRepo) Delete(c context.Context, req *catalog_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "categories" 
				SET 
				"active" = false,
				"deleted_at" = NOW() 
				WHERE id = $1 AND "active"`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to delete category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("category with ID %s not found", req.Id)
	}

	return "deleted", nil
}
