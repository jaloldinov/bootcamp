package postgres

import (
	"context"
	"database/sql"
	"example-grpc-server/models"
	"example-grpc-server/pkg/helper"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{db: db}
}

func (s *staffRepo) CreateStaff(ctx context.Context, req *models.CreateStaff) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "staffs" (
				"id", 
				"branch_id", 
				"tariff_id", 
				"staff_type", 
				"name", 
				"username", 
				"password", 
				"created_at", 
				"updated_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING "id"
	`

	result := s.db.QueryRow(context.Background(), query,
		id,
		req.BranchID,
		req.TariffID,
		req.Type,
		req.Name,
		req.Username,
		req.Password)

	var createdID string
	if err := result.Scan(&createdID); err != nil {
		return "", fmt.Errorf("failed to create staff: %w", err)
	}

	return createdID, nil
}

func (s *staffRepo) UpdateStaff(ctx context.Context, req *models.Staff) (string, error) {
	query := `
		UPDATE "staffs"
			SET "branch_id" = $1, 
			"tariff_id" = $2, 
			"staff_type" = $3, 
			"balance" = $4, 
			"name" = $5, 
			"updated_at" = NOW()
		WHERE "id" = $6
		RETURNING "id"
	`

	result, err := s.db.Exec(context.Background(), query,
		req.BranchID,
		req.TariffID,
		req.Type,
		req.Balance,
		req.Name,
		req.ID)

	if err != nil {
		return "", fmt.Errorf("failed to update staff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff with ID %s not found", req.ID)
	}

	return req.ID, nil
}
func (s *staffRepo) GetStaff(ctx context.Context, req *models.IdRequest) (*models.Staff, error) {
	query := `
		SELECT 
		"id", 
		"branch_id", 
		"tariff_id", 
		"staff_type", 
		"name", 
		"balance", 
		"username", 
		"password", 
		"created_at", 
		"updated_at"
		FROM "staffs"
		WHERE "id" = $1
	`

	staff := models.Staff{}
	var (
		username sql.NullString
		password sql.NullString
		balance  sql.NullFloat64
	)

	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staff.ID,
		&staff.BranchID,
		&staff.TariffID,
		&staff.Type,
		&staff.Name,
		&balance,
		&username,
		&password,
		&staff.CreatedAt,
		&staff.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &models.Staff{}, fmt.Errorf("staff not found")
		}
		return &models.Staff{}, fmt.Errorf("failed to get staff: %w", err)
	}
	staff.Balance = balance.Float64
	staff.Username = username.String
	staff.Password = password.String

	return &staff, nil
}

func (u *staffRepo) GetByUsername(ctx context.Context, req *models.RequestByUsername) (*models.Staff, error) {
	query := `
		SELECT 
		"id", 
		"branch_id", 
		"tariff_id", 
		"staff_type", 
		"name", 
		"balance", 
		"username", 
		"password", 
		"created_at", 
		"updated_at"
		FROM "staffs"
		WHERE "username" = $1
	`

	staff := models.Staff{}
	var (
		username sql.NullString
		password sql.NullString
		balance  sql.NullFloat64
	)
	err := u.db.QueryRow(context.Background(), query, req.Username).Scan(
		&staff.ID,
		&staff.BranchID,
		&staff.TariffID,
		&staff.Type,
		&staff.Name,
		&balance,
		&username,
		&password,
		&staff.CreatedAt,
		&staff.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &models.Staff{}, fmt.Errorf("staff not found")
		}
		return &models.Staff{}, fmt.Errorf("failed to get staff: %w", err)
	}
	staff.Balance = balance.Float64
	staff.Username = username.String
	staff.Password = password.String
	return &staff, nil
}

func (s *staffRepo) GetAllStaff(ctx context.Context, req *models.GetAllStaffRequest) (*models.GetAllStaff, error) {
	params := make(map[string]interface{})
	filter := ""

	query := `
		SELECT 
			"id", 
			"branch_id", 
			"tariff_id", 
			"staff_type", 
			"name", 
			"balance", 
			"username", 
			"password", 
			"created_at", 
			"updated_at"
		FROM "staffs"
	`
	if req.Name != "" {
		filter += ` WHERE "name" ILIKE '%' || :search || '%' `
		params["search"] = req.Name
	}

	offset := (req.Page - 1) * req.Limit
	params["limit"] = req.Limit
	params["offset"] = offset

	query = query + filter + " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := s.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp := &models.GetAllStaff{}
	resp.Staffs = make([]models.Staff, 0)
	count := 0
	for rows.Next() {
		var (
			username sql.NullString
			password sql.NullString
			balance  sql.NullFloat64
		)
		var staff models.Staff
		count++
		err := rows.Scan(
			&staff.ID,
			&staff.BranchID,
			&staff.TariffID,
			&staff.Type,
			&staff.Name,
			&balance,
			&username,
			&password,
			&staff.CreatedAt,
			&staff.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		staff.Username = username.String
		staff.Password = password.String
		staff.Balance = balance.Float64
		resp.Staffs = append(resp.Staffs, staff)
	}

	resp.Count = count
	return resp, nil
}

func (s *staffRepo) DeleteStaff(ctx context.Context, req *models.IdRequest) (string, error) {
	query := `
		DELETE FROM "staffs"
		WHERE "id" = $1
		RETURNING "id"
	`

	var deletedID string

	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(&deletedID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("staff not found")
		}
		return "", fmt.Errorf("failed to delete staff: %w", err)
	}

	return deletedID, nil
}

func (s *staffRepo) UpdateBalance(ctx context.Context, req *models.UpdateBalanceRequest) (res string, err error) {

	// begin transaction
	tr, err := s.db.Begin(context.Background())
	// rollback and commit function
	defer func() {
		if err != nil {
			tr.Rollback(context.Background())
		} else {
			tr.Commit(context.Background())
		}
	}()

	// query for cashier
	cashierBQuery := `
			UPADTE "staffs" 
				SET "balance" = "balance" + $2
			WHERE "id" = $1`

	// checks transaction type if it is("withraw"), balance should be negative
	if req.TransactionType == "withdraw" {
		req.Cashier.Amount = -req.Cashier.Amount
		req.ShopAssisstant.Amount = -req.ShopAssisstant.Amount
	}

	// Execting for cashier
	_, err = tr.Exec(context.Background(), cashierBQuery, req.Cashier.StaffId, req.Cashier.Amount)
	if err != nil {
		return "", err
	}

	cashierTIQuery := `
				INSERT INTO "transactions" (
									"id", 
									"type", 
									"amount", 
									"source_type", 
									"text", 
									"sale_id", 
									"staff_id", 
									"created_at")
				VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`

	_, err = tr.Exec(context.Background(), cashierTIQuery,
		uuid.NewString(),
		req.TransactionType,
		req.Cashier.Amount,
		req.SourceType,
		req.Text,
		req.SaleId,
		req.Cashier.StaffId,
	)
	if err != nil {
		return "", err
	}

	// IF SHOP ASSSISTANT HAS ID RUN THIS
	if req.ShopAssisstant.StaffId != "" {
		// query for shop_assistant
		assistBQuery := `
	UPADTE "staffs" 
		SET "balance" = "balance" + $2
	WHERE "id" = $1`

		// Execting for shop_assistant
		_, err = tr.Exec(context.Background(), assistBQuery, req.ShopAssisstant.StaffId, req.ShopAssisstant.Amount)
		if err != nil {
			return "", err
		}

		assistTIQuery := `
		INSERT INTO "transactions" (
							"id", 
							"type", 
							"amount", 
							"source_type", 
							"text", 
							"sale_id", 
							"staff_id", 
							"created_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`

		_, err = tr.Exec(context.Background(), assistTIQuery,
			uuid.NewString(),
			req.TransactionType,
			req.ShopAssisstant.Amount,
			req.SourceType,
			req.Text,
			req.SaleId,
			req.ShopAssisstant.StaffId,
		)
		if err != nil {
			return "", err
		}
	}

	return "balance updated", nil
}

// func (u *staffRepo) Exists(ctx context.Context,req models.ExistsReq) bool {
// 	staffes := []models.Staff{}
// 	for _, s := range staffes {
// 		if req.Phone == s.Phone {
// 			return true
// 		}
// 	}
// 	return false
// }

func (s *staffRepo) ChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (res string, err error) {
	query := `
	UPDATE "staffs"
		SET "password" = $1, 
		"updated_at" = NOW()
	WHERE "id" = $2`

	result, err := s.db.Exec(context.Background(), query,
		req.NewPassword,
		req.Id)

	if err != nil {
		return "", fmt.Errorf("failed to update staff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff with ID %s not found", req.Id)
	}

	return req.Id, nil
}
