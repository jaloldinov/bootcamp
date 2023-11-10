package postgres

import (
	"context"
	"database/sql"
	"fmt"
	pb "sale_service/genproto"
	"sale_service/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type saleProductRepo struct {
	db *pgxpool.Pool
}

func NewSaleProductRepo(db *pgxpool.Pool) *saleProductRepo {
	return &saleProductRepo{db: db}
}

func (c *saleProductRepo) CreateSaleProduct(ctx context.Context, req *pb.CreateSaleProductRequest) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "sale_products"(
			"id", 
			"sale_id", 
			"product_id", 
			"quantity", 
			"price", 
			"created_at")
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := c.db.Exec(context.Background(), query,
		id,
		req.SaleId,
		req.ProductId,
		req.Quantity,
		req.Price,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create sale_product: %w", err)
	}

	return id, nil
}

func (c *saleProductRepo) GetSaleProduct(ctx context.Context, req *pb.IdRequest) (resp *pb.SaleProduct, err error) {
	var created_at sql.NullString
	var updated_at sql.NullString
	query := `
    SELECT 
			"id", 
			"sale_id", 
			"product_id", 
			"quantity", 
			"price", 
			"created_at", 
			"updated_at"
    FROM "sale_products" WHERE "deleted_at" IS NULL AND id = $1
    `

	saleProduct := pb.SaleProduct{}
	err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&saleProduct.Id,
		&saleProduct.SaleId,
		&saleProduct.ProductId,
		&saleProduct.Quantity,
		&saleProduct.Price,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("saleProduct not found")
		}
		return nil, fmt.Errorf("failed to get saleProduct: %w", err)
	}

	saleProduct.CreatedAt = created_at.String

	if updated_at.Valid {
		saleProduct.UpdatedAt = updated_at.String
	}

	return &saleProduct, nil
}

func (c *saleProductRepo) GetAllSaleProduct(ctx context.Context, req *pb.ListSaleProductRequest) (*pb.ListSaleProductResponse, error) {
	params := make(map[string]interface{})
	filter := ` WHERE "deleted_at" IS NULL `

	var created_at sql.NullString
	var updated_at sql.NullString

	selectQuery := `
		SELECT 
			"id", 
			"sale_id", 
			"product_id", 
			"quantity", 
			"price", 
			"created_at", 
			"updated_at"
		FROM "sale_products"
	`

	offset := (req.Page - 1) * req.Limit

	params["limit"] = req.Limit
	params["offset"] = offset

	query := selectQuery + filter + " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := c.db.Query(ctx, q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp := &pb.ListSaleProductResponse{}
	resp.SaleProducts = make([]*pb.SaleProduct, 0)
	count := 0
	for rows.Next() {
		var saleProduct pb.SaleProduct
		count++
		err := rows.Scan(
			&saleProduct.Id,
			&saleProduct.SaleId,
			&saleProduct.ProductId,
			&saleProduct.Quantity,
			&saleProduct.Price,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		saleProduct.CreatedAt = created_at.String
		if updated_at.Valid {
			saleProduct.UpdatedAt = updated_at.String
		}

		resp.SaleProducts = append(resp.SaleProducts, &saleProduct)
	}

	resp.Count = int32(count)
	return resp, nil
}

func (c *saleProductRepo) UpdateSaleProduct(ctx context.Context, req *pb.UpdateSaleProductRequest) (string, error) {

	query := `
				UPDATE sale_products 
				SET 
					sale_id = $1, 
					product_id = $2, 
					quantity = $3, 
					price = $4, 
					updated_at = NOW() 
				WHERE id = $5 RETURNING id`

	result, err := c.db.Exec(
		context.Background(),
		query,
		req.SaleId,
		req.ProductId,
		req.Quantity,
		req.Price,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update sale_product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("sale_product with ID %s not found", req.Id)
	}

	return "updated", nil
}

func (c *saleProductRepo) DeleteSaleProduct(ctx context.Context, req *pb.IdRequest) (resp string, err error) {
	query := `UPDATE  "sale_products"  
				SET "deleted_at" = NOW() 
			WHERE "deleted_at" IS  NULL AND "id" = $1 `

	resul, err := c.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete saleProduct: %w", err)
	}

	if resul.RowsAffected() == 0 {
		return "", fmt.Errorf("saleProduct with ID %s not found", req.Id)
	}

	return "deleted", nil
}
