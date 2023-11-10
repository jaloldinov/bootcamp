package postgres

import (
	"context"
	"fmt"
	"staff_service/pkg/helper"
	"time"

	staff_service "staff_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaff(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (b *staffRepo) CreateStaff(c context.Context, req *staff_service.CreateStaffRequest) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO "staffs"(
			"id",
			"name", 
			"branch_id", 
			"tariff_id", 
			"staff_type",    
			"username",  
			"password",
			"created_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.BranchId,
		req.TariffId,
		req.StaffType,
		req.Username,
		req.Password,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create staff: %w", err)
	}

	return id, nil
}

func (b *staffRepo) GetStaff(c context.Context, req *staff_service.IdRequest) (resp *staff_service.Staff, err error) {
	query := `
			SELECT 
			"id",
			"name", 
			"branch_id", 
			"tariff_id", 
			"staff_type",  
			"balance",
			"username",  
			"password",
			"created_at",
			"updated_at" 
			FROM "staffs" 
			WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	staff := staff_service.Staff{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staff.Id,
		&staff.Name,
		&staff.BranchId,
		&staff.TariffId,
		&staff.StaffType,
		&staff.Balance,
		&staff.Username,
		&staff.Password,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("staff not found")
		}
		return nil, fmt.Errorf("failed to get staff: %w", err)
	}

	staff.CreatedAt = createdAt.Format(time.RFC3339)
	staff.UpdatedAt = createdAt.Format(time.RFC3339)

	return &staff, nil
}

func (b *staffRepo) UpdateStaff(c context.Context, req *staff_service.UpdateStaffRequest) (string, error) {

	query := `
				UPDATE "staffs" 
				SET 
				"name" = $1,
				"branch_id" = $2,
				"tariff_id" = $3,
				"staff_type" = $4,  
				"balance" = $5,
				"username" = $6,  
				"password" = $7,
				"updated_at" = NOW() 
				WHERE id = $8 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.BranchId,
		req.TariffId,
		req.StaffType,
		req.Balance,
		req.Username,
		req.Password,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update staff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff with ID %s not found", req.Id)
	}

	return fmt.Sprintf("staff with ID %s updated", req.Id), nil
}

func (b *staffRepo) GetAllStaff(c context.Context, req *staff_service.ListStaffRequest) (*staff_service.ListStaffResponse, error) {
	var (
		resp   staff_service.ListStaffResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM "staffs" WHERE true ` + filter

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
			"branch_id", 
			"tariff_id", 
			"staff_type",  
			"balance",
			"username",  
			"password",
			"created_at",
			"updated_at" 
			FROM "staffs" 
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
		var staff staff_service.Staff

		err = rows.Scan(
			&staff.Id,
			&staff.Name,
			&staff.BranchId,
			&staff.TariffId,
			&staff.StaffType,
			&staff.Balance,
			&staff.Username,
			&staff.Password,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning staff err: %w", err)
		}

		staff.CreatedAt = createdAt.Format(time.RFC3339)
		staff.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.Staffs = append(resp.Staffs, &staff)
	}

	return &resp, nil
}

func (b *staffRepo) DeleteStaff(c context.Context, req *staff_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM "staffs" 
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
		return "", fmt.Errorf("staff with ID %s not found", req.Id)

	}

	return fmt.Sprintf("staff with ID %s deleted", req.Id), nil
}
