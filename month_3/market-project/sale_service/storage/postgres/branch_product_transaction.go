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

type branchPrTranRepo struct {
	db *pgxpool.Pool
}

func NewBranchPrTranRepo(db *pgxpool.Pool) *branchPrTranRepo {
	return &branchPrTranRepo{db: db}
}

func (c *branchPrTranRepo) CreateBranchPrTran(ctx context.Context, req *pb.CreateBranchPrTransactionRequest) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "branch_product_transactions"(
			"id", 
			"branch_id", 
			"staff_id", 
			"product_id", 
			"price", 
			"type",
			"quantity",
			"created_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`
	_, err := c.db.Exec(context.Background(), query,
		id,
		req.BranchId,
		req.StaffId,
		req.ProductId,
		req.Price,
		req.Type,
		req.Quantity,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create staff_transaction: %w", err)
	}

	return id, nil
}

func (c *branchPrTranRepo) GetBranchPrTran(ctx context.Context, req *pb.IdRequest) (resp *pb.BranchPrTransaction, err error) {
	var created_at sql.NullString
	var updated_at sql.NullString
	query := `
    SELECT 
			"id", 
			"branch_id", 
			"staff_id", 
			"product_id", 
			"price", 
			"type",
			"quantity",
			"created_at",
			"updated_at"
    FROM "branch_product_transactions" WHERE "deleted_at" IS NULL AND id = $1
    `

	branchPrTran := pb.BranchPrTransaction{}
	err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&branchPrTran.Id,
		&branchPrTran.BranchId,
		&branchPrTran.StaffId,
		&branchPrTran.ProductId,
		&branchPrTran.Price,
		&branchPrTran.Type,
		&branchPrTran.Quantity,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("branchPrTran not found")
		}
		return nil, fmt.Errorf("failed to get branchPrTran: %w", err)
	}

	branchPrTran.CreatedAt = created_at.String

	if updated_at.Valid {
		branchPrTran.UpdatedAt = updated_at.String
	}

	return &branchPrTran, nil
}

func (c *branchPrTranRepo) GetAllBranchPrTran(ctx context.Context, req *pb.ListBranchPrTransactionRequest) (*pb.ListBranchPrTransactionResponse, error) {
	params := make(map[string]interface{})
	filter := ` WHERE "deleted_at" IS NULL `

	var created_at sql.NullString
	var updated_at sql.NullString

	selectQuery := `
		SELECT 
			"id", 
			"branch_id", 
			"staff_id", 
			"product_id", 
			"price", 
			"type",
			"quantity",
			"created_at",
			"updated_at"
		FROM "branch_product_transactions"
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

	resp := &pb.ListBranchPrTransactionResponse{}
	resp.Branches = make([]*pb.BranchPrTransaction, 0)
	count := 0
	for rows.Next() {
		var branchPrTran pb.BranchPrTransaction
		count++
		err := rows.Scan(
			&branchPrTran.Id,
			&branchPrTran.BranchId,
			&branchPrTran.StaffId,
			&branchPrTran.ProductId,
			&branchPrTran.Price,
			&branchPrTran.Type,
			&branchPrTran.Quantity,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		branchPrTran.CreatedAt = created_at.String
		if updated_at.Valid {
			branchPrTran.UpdatedAt = updated_at.String
		}

		resp.Branches = append(resp.Branches, &branchPrTran)
	}

	resp.Count = int32(count)
	return resp, nil
}

func (c *branchPrTranRepo) UpdateBranchPrTran(ctx context.Context, req *pb.UpdateBranchPrTransactionRequest) (string, error) {

	query := `
				UPDATE branch_product_transactions 
				SET  
				"branch_id" = $1, 
				"staff_id" = $2, 
				"product_id" = $3, 
				"price" = $4, 
				"type" = $5,
				"quantity" = $6,
				"updated_at" = NOW() 
				WHERE id = $7 RETURNING id`

	result, err := c.db.Exec(
		context.Background(),
		query,
		&req.BranchId,
		&req.StaffId,
		&req.ProductId,
		&req.Price,
		&req.Type,
		&req.Quantity,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update staff_transaction: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff_transaction with ID %s not found", req.Id)
	}

	return "updated", nil
}

func (c *branchPrTranRepo) DeleteBranchPrTran(ctx context.Context, req *pb.IdRequest) (resp string, err error) {
	query := `UPDATE  "branch_product_transactions"  
				SET "deleted_at" = NOW() 
			WHERE "deleted_at" IS  NULL AND "id" = $1 `

	resul, err := c.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete branchPrTran: %w", err)
	}

	if resul.RowsAffected() == 0 {
		return "", fmt.Errorf("branchPrTran with ID %s not found", req.Id)
	}

	return "deleted", nil
}
