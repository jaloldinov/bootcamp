package postgres

import (
	"branch_service/pkg/helper"
	"context"
	"fmt"
	"time"

	pb "branch_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchProductRepo struct {
	db *pgxpool.Pool
}

func NewBranchProduct(db *pgxpool.Pool) *branchProductRepo {
	return &branchProductRepo{
		db: db,
	}
}

func (b *branchProductRepo) CreateBranchProduct(c context.Context, req *pb.CreateBranchProductRequest) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "branch_products"(
			"id", 
			"product_id",
			"branch_id", 
			"quantity",  
			"created_at")
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.ProductId,
		req.BranchId,
		req.Quantity,
		time.Now(),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create branch_product: %w", err)
	}

	return id, nil
}

func (b *branchProductRepo) GetBranchProduct(c context.Context, req *pb.IdRequest) (resp *pb.BranchProduct, err error) {
	query := `
			SELECT 
				"id", 
				"product_id",
				"branch_id", 
				"quantity",  
				"created_at",
				"updated_at" 
			FROM "branch_products" 
			WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	branch := pb.BranchProduct{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&branch.Id,
		&branch.ProductId,
		&branch.BranchId,
		&branch.Quantity,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("branch product not found")
		}
		return nil, fmt.Errorf("failed to get branch product: %w", err)
	}

	branch.CreatedAt = createdAt.Format(time.RFC3339)
	branch.UpdatedAt = createdAt.Format(time.RFC3339)

	return &branch, nil
}

func (b *branchProductRepo) GetAllBranchProduct(c context.Context, req *pb.ListBranchProductRequest) (*pb.ListBranchProductResponse, error) {
	var (
		resp   pb.ListBranchProductResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.ProductId != "" {
		filter += " AND product_id = :product_id"
		params["product_id"] = req.ProductId
	}

	if req.BranchId != "" {
		filter += " AND branch_id = :branch_id"
		params["branch_id"] = req.BranchId
	}

	countQuery := `SELECT count(1) FROM "branch_products" WHERE true ` + filter

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
				"product_id",
				"branch_id", 
				"quantity",  
				"created_at",
				"updated_at" 
			FROM "branch_products" 
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
		var branch pb.BranchProduct

		err = rows.Scan(
			&branch.Id,
			&branch.ProductId,
			&branch.BranchId,
			&branch.Quantity,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning branch err: %w", err)
		}

		branch.CreatedAt = createdAt.Format(time.RFC3339)
		branch.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.BranchProductes = append(resp.BranchProductes, &branch)
	}

	return &resp, nil
}

func (b *branchProductRepo) UpdateBranchProduct(c context.Context, req *pb.UpdateBranchProductRequest) (string, error) {

	query := `
				UPDATE "branch_products" 
				SET 
				"product_id" = $1, 
				"branch_id" = $2, 
				"quantity" = $3, 
				"updated_at" = NOW() 
				WHERE id = $4 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.ProductId,
		req.BranchId,
		req.Quantity,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update branch product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch product with ID %s not found", req.Id)
	}

	return fmt.Sprintf("branch product with ID %s updated", req.Id), nil
}

func (b *branchProductRepo) DeleteBranchProduct(c context.Context, req *pb.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM "branch_products" 
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
