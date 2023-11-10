package postgres

import (
	"context"
	"database/sql"
	"fmt"

	user_service "user_service/genproto"
	"user_service/pkg/helper"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranch(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) GetListActive(c context.Context, req *user_service.ListBranchActiveRequest) (*user_service.ListBranchResponse, error) {

	var (
		updatedAt sql.NullString
		createdAt sql.NullString
		resp      user_service.ListBranchResponse
		err       error
		filter    string
		params    = make(map[string]interface{})
	)

	if req.Time != "" {
		filter += " AND :time BETWEEN work_hour_start AND work_hout_end "
		params["time"] = req.Time
	}

	countQuery := `SELECT count(1) FROM "branches" WHERE "deleted_at" IS NULL AND "active"` + filter
	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `
	SELECT 
		id,
		name,
		phone,
		photo,
		delivery_tarif_id,
		work_hour_start::text,
		work_hout_end::text,
		address,
		active,
		destination,
		created_at,
		updated_at 
	FROM branches  where "active" and "deleted_at" is null` + filter

	query += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	params["limit"] = 10
	params["offset"] = 0

	if req.Limit > 0 {
		params["limit"] = req.Limit
	}
	if req.Page > 0 {
		params["offset"] = (req.Page - 1) * req.Limit
	}

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(c, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var branch user_service.Branch
		err = rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Phone,
			&branch.Photo,
			&branch.DeliveryTarifId,
			&branch.WorkHourStart,
			&branch.WorkHourEnd,
			&branch.Address,
			&branch.Active,
			&branch.Destination,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning branch err: %w", err)
		}

		if createdAt.Valid {
			branch.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			branch.UpdatedAt = updatedAt.String
		}
		resp.Branches = append(resp.Branches, &branch)
	}

	return &resp, nil
}

func (b *branchRepo) Create(c context.Context, req *user_service.CreateBranchRequest) (*user_service.Response, error) {

	query := `
	  INSERT INTO branches (
		name,  
		phone,
		photo,
		delivery_tarif_id,
		work_hour_start,
		work_hout_end,
		address,
		destination,
		created_at
	  ) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, NOW()
	  ) RETURNING id`

	var id int
	err := b.db.QueryRow(c, query,
		req.Name,
		req.Phone,
		req.Photo,
		req.DeliveryTarifId,
		req.WorkHourStart,
		req.WorkHourEnd,
		req.Address,
		req.Destination,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to create branch: %w", err)
	}

	return &user_service.Response{Message: fmt.Sprintf("%d", id)}, nil
}

func (b *branchRepo) Get(c context.Context, req *user_service.IdRequest) (resp *user_service.Branch, err error) {
	var (
		updatedAt sql.NullString
	)
	query := `
    SELECT 
		id,
		name,
		phone,
		photo,
		delivery_tarif_id,
		work_hour_start::TEXT,
		work_hout_end::TEXT,
		address,
		active,
		destination,
		created_at::TEXT,
		updated_at 
    FROM branches 
    WHERE id = $1 AND "active" AND "deleted_at" IS NULL;`

	branch := user_service.Branch{}

	err = b.db.QueryRow(c, query, req.Id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Phone,
		&branch.Photo,
		&branch.DeliveryTarifId,
		&branch.WorkHourStart,
		&branch.WorkHourEnd,
		&branch.Address,
		&branch.Active,
		&branch.Destination,
		&branch.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("branch not found")
		}
		return nil, fmt.Errorf("failed to get branch: %w", err)
	}

	if updatedAt.Valid {
		branch.UpdatedAt = updatedAt.String
	}

	return &branch, nil
}

func (b *branchRepo) GetList(c context.Context, req *user_service.ListBranchRequest) (*user_service.ListBranchResponse, error) {

	var (
		updatedAt sql.NullString
		createdAt sql.NullString
		resp      user_service.ListBranchResponse
		err       error
		filter    string
		params    = make(map[string]interface{})
	)

	if req.Name != "" {
		filter += " AND name ILIKE '%' || :name || '%' "
		params["name"] = req.Name
	}

	if req.CreatedAtFrom != "" {
		filter += " AND created_at >= :created_at_from"
		params["created_at_from"] = req.CreatedAtFrom
	}

	if req.CreatedAtTo != "" {
		filter += " AND created_at <= :created_at_to"
		params["created_at_to"] = req.CreatedAtTo
	}
	countQuery := `SELECT count(1) FROM "branches" WHERE "deleted_at" IS NULL AND "active"` + filter
	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `
	SELECT 
		id,
		name,
		phone,
		photo,
		delivery_tarif_id,
		work_hour_start::text,
		work_hout_end::text,
		address,
		active,
		destination,
		created_at,
		updated_at 
	FROM branches  where "active" and "deleted_at" is null` + filter

	query += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	params["limit"] = 10
	params["offset"] = 0

	if req.Limit > 0 {
		params["limit"] = req.Limit
	}
	if req.Page > 0 {
		params["offset"] = (req.Page - 1) * req.Limit
	}

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(c, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var branch user_service.Branch
		err = rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Phone,
			&branch.Photo,
			&branch.DeliveryTarifId,
			&branch.WorkHourStart,
			&branch.WorkHourEnd,
			&branch.Address,
			&branch.Active,
			&branch.Destination,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning branch err: %w", err)
		}

		if createdAt.Valid {
			branch.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			branch.UpdatedAt = updatedAt.String
		}
		resp.Branches = append(resp.Branches, &branch)
	}

	return &resp, nil
}

func (b *branchRepo) Update(c context.Context, req *user_service.UpdateBranchRequest) (string, error) {

	query := `
				UPDATE "branches" 
				SET 
				"name" = $1,
				"phone" = $2,
				"photo" = $3,  
				"delivery_tarif_id" = $4,
				"work_hour_start"=$5,
				"work_hout_end"=$6,
				"address"=$7,
				"destination"=$8,
				"updated_at" = NOW() 
				WHERE id = $9`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.Phone,
		req.Photo,
		req.DeliveryTarifId,
		req.WorkHourStart,
		req.WorkHourEnd,
		req.Address,
		req.Destination,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update branch: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch with ID %d not found", req.Id)
	}

	return fmt.Sprintf("branch with ID %d updated", req.Id), nil
}

func (b *branchRepo) Delete(c context.Context, req *user_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "branches" 
				SET 
				"active" = false,
				"deleted_at" = NOW() 
				WHERE id = $1 AND "deleted_at" IS NULL and "active"`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to delete branch: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch with ID %s not found", req.Id)
	}

	return "deleted", nil
}
