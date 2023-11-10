package postgres

import (
	"context"
	"fmt"
	"product_service/pkg/helper"
	"time"

	product_service "product_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProduct(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (b *productRepo) CreateProduct(c context.Context, req *product_service.CreateProductRequest) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "products"(
			"id",
			"category_id", 
			"name", 
			"description", 
			"price",  
			"created_at")
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.CategoryId,
		req.Name,
		req.Description,
		req.Price,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create product: %w", err)
	}

	return id, nil
}

func (b *productRepo) GetProduct(c context.Context, req *product_service.IdRequest) (resp *product_service.Product, err error) {
	query := `
			SELECT 
				"id",
				"category_id", 
				"name", 
				"description", 
				"price",  
				"created_at",
				"updated_at" 
			FROM "products" 
			WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	product := product_service.Product{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&product.Id,
		&product.CategoryId,
		&product.Name,
		&product.Description,
		&product.Price,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	product.CreatedAt = createdAt.Format(time.RFC3339)
	product.UpdatedAt = createdAt.Format(time.RFC3339)

	return &product, nil
}

func (b *productRepo) UpdateProduct(c context.Context, req *product_service.UpdateProductRequest) (string, error) {

	query := `
				UPDATE "products" 
				SET 
					"category_id" = $1, 
					"name" = $2, 
					"description" = $3, 
					"price" = $4, 
					"updated_at" = NOW() 
				WHERE id = $5 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.CategoryId,
		req.Name,
		req.Description,
		req.Price,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("product with ID %s not found", req.Id)
	}

	return fmt.Sprintf("product with ID %s updated", req.Id), nil
}

func (b *productRepo) GetAllProduct(c context.Context, req *product_service.ListProductRequest) (*product_service.ListProductResponse, error) {
	var (
		resp   product_service.ListProductResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM "products" WHERE true ` + filter

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
				"category_id", 
				"name", 
				"description", 
				"price",  
				"created_at",
				"updated_at" 
			FROM "products" 
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
		var product product_service.Product

		err = rows.Scan(
			&product.Id,
			&product.CategoryId,
			&product.Name,
			&product.Description,
			&product.Price,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning product err: %w", err)
		}

		product.CreatedAt = createdAt.Format(time.RFC3339)
		product.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.Productes = append(resp.Productes, &product)
	}

	return &resp, nil
}

func (b *productRepo) DeleteProduct(c context.Context, req *product_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM "products" 
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
		return "", fmt.Errorf("product with ID %s not found", req.Id)

	}

	return fmt.Sprintf("product with ID %s deleted", req.Id), nil
}
