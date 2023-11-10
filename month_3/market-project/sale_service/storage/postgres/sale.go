package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "sale_service/genproto"
	"sale_service/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type saleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *saleRepo {
	return &saleRepo{db: db}
}

func (c *saleRepo) CreateSale(ctx context.Context, req *pb.CreateSaleRequest) (string, error) {
	id := uuid.NewString()

	query := `
        INSERT INTO "sale" (
			"id", 
			"client_name", 
			"branch_id", 
			"shop_assistant_id",
        	"cashier_id", 
			"price", 
			"payment_type", 
			"created_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
    `

	var err error
	var insertedID string

	if req.CashierId != "" && req.ShopAssistantId != "" {
		// Both cashier_id and shop_assistant_id have values
		_, err = c.db.Exec(context.Background(),
			query,
			id,
			req.ClientName,
			req.BranchId,
			req.ShopAssistantId,
			req.CashierId,
			req.Price,
			req.PaymentType,
		)
		insertedID = id
	} else if req.CashierId != "" {
		// Only cashier_id has a value
		_, err = c.db.Exec(context.Background(),
			query,
			id,
			req.CashierId,
			req.BranchId,
			nil,
			req.CashierId,
			req.Price,
			req.PaymentType,
		)
		insertedID = req.CashierId
	} else if req.ShopAssistantId != "" {
		// Only shop_assistant_id has a value
		_, err = c.db.Exec(context.Background(),
			query,
			id,
			req.ClientName,
			req.BranchId,
			req.ShopAssistantId,
			nil,
			req.Price,
			req.PaymentType,
		)
		insertedID = req.ShopAssistantId
	} else {
		return "", errors.New("either cashier_id or shop_assistant_id should be provided")
	}

	if err != nil {
		return "", fmt.Errorf("failed to create sale: %w", err)
	}

	return insertedID, nil
}

func (c *saleRepo) GetSale(ctx context.Context, req *pb.IdRequest) (resp *pb.Sale, err error) {
	var shopAssistantID sql.NullString
	var created_at sql.NullString
	var updated_at sql.NullString
	query := `
    SELECT 
			"id", 
			"client_name", 
			"branch_id", 
			"shop_assistant_id",
    		"cashier_id", 
			"price", 
			"payment_type", 
			"status", 
			"created_at", 
			"updated_at"
    FROM "sale" WHERE "deleted_at" IS NULL AND id = $1
    `

	sale := pb.Sale{}
	err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&sale.Id,
		&sale.ClientName,
		&sale.BranchId,
		&shopAssistantID,
		&sale.CashierId,
		&sale.Price,
		&sale.PaymentType,
		&sale.Status,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("sale not found")
		}
		return nil, fmt.Errorf("failed to get sale: %w", err)
	}

	// Check if the fields are null and set them to empty strings if necessary
	if shopAssistantID.Valid {
		sale.ShopAssistantId = shopAssistantID.String
	}

	sale.CreatedAt = created_at.String

	if updated_at.Valid {
		sale.UpdatedAt = updated_at.String
	}

	return &sale, nil
}

func (c *saleRepo) GetAllSale(ctx context.Context, req *pb.ListSaleRequest) (*pb.ListSaleResponse, error) {
	params := make(map[string]interface{})
	filter := ` WHERE "deleted_at" IS NULL `

	var shopAssistantID sql.NullString
	var created_at sql.NullString
	var updated_at sql.NullString

	selectQuery := `
		SELECT 
			"id", 
			"client_name", 
			"branch_id", 
			"shop_assistant_id",
			"cashier_id", 
			"price", 
			"payment_type", 
			"status", 
			"created_at", 
			"updated_at"
		FROM "sale"
	`

	if req.Search != "" {
		filter += ` AND "client_name" ILIKE '%' || :search || '%' `
		params["search"] = req.Search
	}

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

	resp := &pb.ListSaleResponse{}
	resp.Sales = make([]*pb.Sale, 0)
	count := 0
	for rows.Next() {
		var sale pb.Sale
		count++
		err := rows.Scan(
			&sale.Id,
			&sale.ClientName,
			&sale.BranchId,
			&shopAssistantID,
			&sale.CashierId,
			&sale.Price,
			&sale.PaymentType,
			&sale.Status,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Check if the fields are null and set them to empty strings if necessary
		if shopAssistantID.Valid {
			sale.ShopAssistantId = shopAssistantID.String
		}

		sale.CreatedAt = created_at.String

		if updated_at.Valid {
			sale.UpdatedAt = updated_at.String
		}

		resp.Sales = append(resp.Sales, &sale)
	}

	resp.Count = int32(count)
	return resp, nil
}

func (c *saleRepo) UpdateSale(ctx context.Context, req *pb.UpdateSaleRequest) (string, error) {
	query := `
	UPDATE "sale" SET
	"client_name" = $1,
	"branch_id" = $2,
	"shop_assistant_id" = $3,
	"cashier_id" = $4,
	"price" = $5,
	"payment_type" = $6,
	"updated_at" = NOW()
	WHERE "id" = $7
	RETURNING "id"
	`

	var updatedID string
	err := error(nil)

	if req.ShopAssistantId != "" && req.CashierId != "" {
		// update both shop_assistant_id and cashier_id,
		err = c.db.QueryRow(
			context.Background(),
			query,
			req.ClientName,
			req.BranchId,
			req.ShopAssistantId,
			req.CashierId,
			req.Price,
			req.PaymentType,
			req.Id,
		).Scan(&updatedID)
	} else if req.ShopAssistantId == "" {
		// shop_assistant_id is empty, update only cashier_id
		err = c.db.QueryRow(
			context.Background(),
			query,
			req.ClientName,
			req.BranchId,
			nil,
			req.CashierId,
			req.Price,
			req.PaymentType,
			req.Id,
		).Scan(&updatedID)
	} else if req.CashierId == "" {
		// cashier_id is empty, update only shop_assistant_id
		err = c.db.QueryRow(
			context.Background(),
			query,
			req.ClientName,
			req.BranchId,
			req.ShopAssistantId,
			nil,
			req.Price,
			req.PaymentType,
			req.Id,
		).Scan(&updatedID)
	}

	if err != nil {
		return "", fmt.Errorf("failed to update sale: %w", err)
	}

	return updatedID, nil
}

func (c *saleRepo) DeleteSale(ctx context.Context, req *pb.IdRequest) (resp string, err error) {
	query := `UPDATE  "sale"  
				SET "deleted_at" = NOW() 
			WHERE "deleted_at" IS  NULL AND "id" = $1 `

	resul, err := c.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete sale: %w", err)
	}

	if resul.RowsAffected() == 0 {
		return "", fmt.Errorf("sale with ID %s not found", req.Id)
	}

	return req.Id, nil
}
