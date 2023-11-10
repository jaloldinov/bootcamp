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

type staffTransactionRepo struct {
	db *pgxpool.Pool
}

func NewStaffTransactionRepo(db *pgxpool.Pool) *staffTransactionRepo {
	return &staffTransactionRepo{db: db}
}

func (c *staffTransactionRepo) CreateStaffTransaction(ctx context.Context, req *pb.CreateStaffTransactionRequest) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "staff_transactions"(
			"id", 
			"type", 
			"amount", 
			"source_type", 
			"sale_id", 
			"staff_id",
			"text",
			"created_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`
	_, err := c.db.Exec(context.Background(), query,
		id,
		req.TrType,
		req.Amount,
		req.SourceType,
		req.SaleId,
		req.StaffId,
		req.AboutText,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create staff_transaction: %w", err)
	}

	return id, nil
}

func (c *staffTransactionRepo) GetStaffTransaction(ctx context.Context, req *pb.IdRequest) (resp *pb.StaffTransaction, err error) {
	var created_at sql.NullString
	var updated_at sql.NullString
	query := `
    SELECT 
			"id", 
			"type", 
			"amount", 
			"source_type", 
			"sale_id", 
			"staff_id",
			"text",
			"created_at",
			"updated_at"
    FROM "staff_transactions" WHERE "deleted_at" IS NULL AND id = $1
    `

	staffTransaction := pb.StaffTransaction{}
	err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staffTransaction.Id,
		&staffTransaction.TrType,
		&staffTransaction.Amount,
		&staffTransaction.SourceType,
		&staffTransaction.SaleId,
		&staffTransaction.StaffId,
		&staffTransaction.AboutText,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("staffTransaction not found")
		}
		return nil, fmt.Errorf("failed to get staffTransaction: %w", err)
	}

	staffTransaction.CreatedAt = created_at.String

	if updated_at.Valid {
		staffTransaction.UpdatedAt = updated_at.String
	}

	return &staffTransaction, nil
}

func (c *staffTransactionRepo) GetAllStaffTransaction(ctx context.Context, req *pb.ListStaffTransactionRequest) (*pb.ListStaffTransactionResponse, error) {
	params := make(map[string]interface{})
	filter := ` WHERE "deleted_at" IS NULL `

	var created_at sql.NullString
	var updated_at sql.NullString

	selectQuery := `
		SELECT 
			"id", 
			"type", 
			"amount", 
			"source_type", 
			"sale_id", 
			"staff_id",
			"text",
			"created_at",
			"updated_at"
		FROM "staff_transactions"
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

	resp := &pb.ListStaffTransactionResponse{}
	resp.StaffTransactions = make([]*pb.StaffTransaction, 0)
	count := 0
	for rows.Next() {
		var staffTransaction pb.StaffTransaction
		count++
		err := rows.Scan(
			&staffTransaction.Id,
			&staffTransaction.TrType,
			&staffTransaction.Amount,
			&staffTransaction.SourceType,
			&staffTransaction.SaleId,
			&staffTransaction.StaffId,
			&staffTransaction.AboutText,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		staffTransaction.CreatedAt = created_at.String
		if updated_at.Valid {
			staffTransaction.UpdatedAt = updated_at.String
		}

		resp.StaffTransactions = append(resp.StaffTransactions, &staffTransaction)
	}

	resp.Count = int32(count)
	return resp, nil
}

func (c *staffTransactionRepo) UpdateStaffTransaction(ctx context.Context, req *pb.UpdateStaffTransactionRequest) (string, error) {

	query := `
				UPDATE staff_transactions 
				SET 
				"type" = $1,
				"amount" = $2,
				"source_type" = $3,
				"sale_id" = $4,
				"staff_id" = $5,
				"text" = $6,
				"updated_at" = NOW() 
				WHERE id = $7 RETURNING id`

	result, err := c.db.Exec(
		context.Background(),
		query,
		&req.TrType,
		&req.Amount,
		&req.SourceType,
		&req.SaleId,
		&req.StaffId,
		&req.AboutText,
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

func (c *staffTransactionRepo) DeleteStaffTransaction(ctx context.Context, req *pb.IdRequest) (resp string, err error) {
	query := `UPDATE  "staff_transactions"  
				SET "deleted_at" = NOW() 
			WHERE "deleted_at" IS  NULL AND "id" = $1 `

	resul, err := c.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete staffTransaction: %w", err)
	}

	if resul.RowsAffected() == 0 {
		return "", fmt.Errorf("staffTransaction with ID %s not found", req.Id)
	}

	return "deleted", nil
}
