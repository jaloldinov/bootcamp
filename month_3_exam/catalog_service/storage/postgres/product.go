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

type productRepo struct {
	db *pgxpool.Pool
}

func NewProduct(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (b *productRepo) Create(c context.Context, req *catalog_service.CreateProductRequest) (string, error) {

	query := `
		INSERT INTO "products"(
			"title", 
			"description", 
			"photo",    
			"order_number", 
			"type",
			"price",
			"category_id",
			"created_at"
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW()) RETURNING "id"
	`

	var id int
	err := b.db.QueryRow(c, query,
		req.Title,
		req.Description,
		req.Photo,
		req.OrderNumber,
		req.ProductType,
		req.Price,
		req.CategoryId,
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to create product: %w", err)
	}

	return fmt.Sprintf("created with ID: %d", id), nil

}

func (b *productRepo) Get(c context.Context, req *catalog_service.IdRequest) (resp *catalog_service.Product, err error) {
	query := `
		SELECT 
			"id",
			"title", 
			"description", 
			"photo",    
			"order_number", 
			"active",
			"type",
			"price",
			"category_id",
			"created_at",
			"updated_at" 
		FROM "products" 
		WHERE id=$1 AND "active" AND "deleted_at" IS NULL `

	var (
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	product := catalog_service.Product{}
	err = b.db.QueryRow(c, query, req.Id).Scan(
		&product.Id,
		&product.Title,
		&product.Description,
		&product.Photo,
		&product.OrderNumber,
		&product.Active,
		&product.ProductType,
		&product.Price,
		&product.CategoryId,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	if createdAt.Valid {
		product.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		product.UpdatedAt = updatedAt.String
	}

	return &product, nil
}

func (b *productRepo) GetList(c context.Context, req *catalog_service.ListProductRequest) (*catalog_service.ListProductResponse, error) {
	var (
		resp   catalog_service.ListProductResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND title ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	if req.Type != "" {
		filter += " AND type = :type"
		params["type"] = req.Type
	}

	if req.Category != 0 {
		filter += " AND category_id = :category_id"
		params["category_id"] = req.Category
	}

	countQuery := `SELECT count(1) FROM "products" WHERE "deleted_at" IS NULL AND "active"` + filter

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
			"description", 
			"photo",    
			"order_number", 
			"active",
			"type",
			"price",
			"category_id",
			"created_at",
			"updated_at" 
			FROM "products" 
		    WHERE "active" AND "deleted_at" IS NULL ` + filter

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

	for rows.Next() {
		var product catalog_service.Product

		err = rows.Scan(
			&product.Id,
			&product.Title,
			&product.Description,
			&product.Photo,
			&product.OrderNumber,
			&product.Active,
			&product.ProductType,
			&product.Price,
			&product.CategoryId,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning product err: %w", err)
		}

		if createdAt.Valid {
			product.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			product.UpdatedAt = createdAt.String
		}
		resp.Products = append(resp.Products, &product)
	}

	return &resp, nil
}

func (b *productRepo) Update(c context.Context, req *catalog_service.UpdateProductRequest) (string, error) {

	query := `
				UPDATE "products" 
				SET 
				"title" = $1,
				"description" = $2,
				"photo" = $3, 
				"order_number" = $4,
				"type" = $5,
				"price" = $6,
				"category_id" = $7,
				"updated_at" = NOW()
				WHERE id = $8 AND "active" AND "deleted_at" IS NULL`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Title,
		req.Description,
		req.Photo,
		req.OrderNumber,
		req.ProductType,
		req.Price,
		req.CategoryId,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("product with ID %d not found", req.Id)
	}

	return fmt.Sprintf("product with ID %d updated", req.Id), nil
}

func (b *productRepo) Delete(c context.Context, req *catalog_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "products" 
				SET 
				"active" = false,
				"deleted_at" = NOW() 
				WHERE id = $1 AND "active" AND "deleted_at" IS NULL`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to delete product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("product with ID %s not found", req.Id)
	}

	return "deleted", nil
}
