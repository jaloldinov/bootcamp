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

type courierRepo struct {
	db *pgxpool.Pool
}

func NewCourier(db *pgxpool.Pool) *courierRepo {
	return &courierRepo{
		db: db,
	}
}

func (b *courierRepo) Create(c context.Context, req *user_service.CreateCouriersRequest) (*user_service.Response, error) {

	hashedPass, err := helper.GeneratePasswordHash(req.Password)
	if err != nil {
		fmt.Println("error while generating password")
		return nil, err
	}

	query := `
	  INSERT INTO couriers (
		first_name,
		last_name,
		branch_id,  
		phone,
		login,
		password,
		max_order_count,
		created_at
	  ) VALUES (
		$1, $2, $3, $4, $5, $6,$7, now()
	  ) RETURNING id`

	var id int
	err = b.db.QueryRow(c, query,
		req.Firstname,
		req.Lastname,
		req.BranchId,
		req.Phone,
		req.Login,
		hashedPass,
		req.MaxOrderCount,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to create courier: %w", err)
	}

	return &user_service.Response{Message: fmt.Sprintf("%d", id)}, nil
}

func (b *courierRepo) Get(c context.Context, req *user_service.IdRequest) (resp *user_service.Couriers, err error) {
	var (
		createdAt sql.NullString
		updatedAt sql.NullString
	)
	query := `
    SELECT 
		id,
		first_name,
		last_name,
		branch_id,  
		phone,
		login,
		active,
		password,
		max_order_count,
		created_at,
		updated_at 
    FROM couriers 
    WHERE id = $1 AND "active" AND "deleted_at" IS NULL;`

	courier := user_service.Couriers{}

	err = b.db.QueryRow(c, query, req.Id).Scan(
		&courier.Id,
		&courier.Firstname,
		&courier.Lastname,
		&courier.BranchId,
		&courier.Phone,
		&courier.Login,
		&courier.Active,
		&courier.Password,
		&courier.MaxOrderCount,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("courier not found")
		}
		return nil, fmt.Errorf("failed to get courier: %w", err)
	}
	if createdAt.Valid {
		courier.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		courier.UpdatedAt = updatedAt.String
	}

	return &courier, nil
}

func (b *courierRepo) GetList(c context.Context, req *user_service.ListCouriersRequest) (*user_service.ListCouriersResponse, error) {

	var (
		updatedAt sql.NullString
		createdAt sql.NullString
		resp      user_service.ListCouriersResponse
		err       error
		filter    string
		params    = make(map[string]interface{})
	)

	if req.Firstname != "" {
		filter += " AND (first_name ILIKE '%' || :firstname || '%') "
		params["firstname"] = req.Firstname
	}

	if req.Lastname != "" {
		filter += " AND (last_name ILIKE '%' || :lastname || '%') "
		params["lastname"] = req.Lastname
	}

	if req.Phone != "" {
		filter += " AND (phone = :phone) "
		params["phone"] = req.Phone
	}

	if req.CreatedAtFrom != "" {
		filter += " AND created_at >= :created_at_from"
		params["created_at_from"] = req.CreatedAtFrom
	}

	if req.CreatedAtTo != "" {
		filter += " AND created_at <= :created_at_to"
		params["created_at_to"] = req.CreatedAtTo
	}

	countQuery := `SELECT count(1) FROM "couriers" WHERE "deleted_at" IS NULL AND "active"` + filter
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
	        first_name,
	        last_name,
	        phone,
			active,
			branch_id,
	        login,
	        password,
			max_order_count,
	    	created_at::text,
		    updated_at 
	FROM couriers  where "active" and "deleted_at" is null` + filter

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
		var courier = user_service.Couriers{}
		err = rows.Scan(
			&courier.Id,
			&courier.Firstname,
			&courier.Lastname,
			&courier.Phone,
			&courier.Active,
			&courier.BranchId,
			&courier.Login,
			&courier.Password,
			&courier.MaxOrderCount,
			&courier.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning courier err: %w", err)
		}

		if createdAt.Valid {
			courier.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			courier.UpdatedAt = updatedAt.String
		}
		resp.Couriers = append(resp.Couriers, &courier)
	}

	return &resp, nil
}

func (b *courierRepo) Update(c context.Context, req *user_service.UpdateCouriersRequest) (string, error) {

	query := `
				UPDATE "couriers" 
				SET 
				"first_name" = $1,
				"last_name" = $2,
				"branch_id" = $3,  
				"phone" = $4,
				"login"=$5,
				"password"=$6,
				"max_order_count"=$7,
				"updated_at" = NOW() 
				WHERE id = $8`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Firstname,
		req.Lastname,
		req.BranchId,
		req.Phone,
		req.Login,
		req.Password,
		req.MaxOrderCount,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update courier: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("courier with ID %d not found", req.Id)
	}

	return fmt.Sprintf("courier with ID %d updated", req.Id), nil
}

func (b *courierRepo) Delete(c context.Context, req *user_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "couriers" 
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
		return "", fmt.Errorf("failed to delete courier: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("courier with ID %s not found", req.Id)
	}

	return "deleted", nil
}

func (b *courierRepo) GetByLogin(c context.Context, req *user_service.IdRequest) (resp *user_service.Couriers, err error) {
	var (
		createdAt sql.NullString
		updatedAt sql.NullString
	)
	query := `
    SELECT 
		id,
		first_name,
		last_name,
		branch_id,  
		phone,
		login,
		active,
		password,
		max_order_count,
		created_at,
		updated_at 
    FROM couriers 
    WHERE login = $1 AND "active" AND "deleted_at" IS NULL;`

	courier := user_service.Couriers{}

	err = b.db.QueryRow(c, query, req.Id).Scan(
		&courier.Id,
		&courier.Firstname,
		&courier.Lastname,
		&courier.BranchId,
		&courier.Phone,
		&courier.Login,
		&courier.Active,
		&courier.Password,
		&courier.MaxOrderCount,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("courier not found")
		}
		return nil, fmt.Errorf("failed to get courier: %w", err)
	}
	if createdAt.Valid {
		courier.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		courier.UpdatedAt = updatedAt.String
	}

	return &courier, nil
}
