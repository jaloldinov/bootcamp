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

type userRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (b *userRepo) Create(c context.Context, req *user_service.CreateUsersRequest) (*user_service.Response, error) {
	hashedPass, err := helper.GeneratePasswordHash(req.Password)
	if err != nil {
		fmt.Println("error while generating password")
		return nil, err
	}
	query := `
	  INSERT INTO users (
		first_name,
		last_name,
		branch_id,  
		phone,
		login,
		password,
		created_at
	  ) VALUES (
		$1, $2, $3, $4, $5, $6, now()
	  ) RETURNING id`

	var id int
	err = b.db.QueryRow(c, query,
		req.Firstname,
		req.Lastname,
		req.BranchId,
		req.Phone,
		req.Login,
		hashedPass,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user_service.Response{Message: fmt.Sprintf("%d", id)}, nil
}

func (b *userRepo) Get(c context.Context, req *user_service.IdRequest) (resp *user_service.Users, err error) {
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
		created_at,
		updated_at 
    FROM users 
    WHERE id = $1 AND "active" AND "deleted_at" IS NULL;`

	user := user_service.Users{}

	err = b.db.QueryRow(c, query, req.Id).Scan(
		&user.Id,
		&user.Firstname,
		&user.Lastname,
		&user.BranchId,
		&user.Phone,
		&user.Login,
		&user.Active,
		&user.Password,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if createdAt.Valid {
		user.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.String
	}

	return &user, nil
}

func (b *userRepo) GetList(c context.Context, req *user_service.ListUsersRequest) (*user_service.ListUsersResponse, error) {

	var (
		updatedAt sql.NullString
		createdAt sql.NullString
		resp      user_service.ListUsersResponse
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

	countQuery := `SELECT count(1) FROM "users" WHERE "deleted_at" IS NULL AND "active"` + filter
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
	        branch_id,  
	        phone,
	        login,
			active,
	        password,
	    	created_at::text,
		    updated_at 
	FROM users  where "active" and "deleted_at" is null` + filter

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
		var user user_service.Users
		err = rows.Scan(
			&user.Id,
			&user.Firstname,
			&user.Lastname,
			&user.BranchId,
			&user.Phone,
			&user.Login,
			&user.Active,
			&user.Password,
			&user.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning users err: %w", err)
		}

		if createdAt.Valid {
			user.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.String
		}
		resp.Users = append(resp.Users, &user)
	}

	return &resp, nil
}

func (b *userRepo) Update(c context.Context, req *user_service.UpdateUsersRequest) (string, error) {
	hashedPass, err := helper.GeneratePasswordHash(req.Password)
	if err != nil {
		fmt.Println("error while generating password")
		return "", err
	}
	query := `
				UPDATE "users" 
				SET 
				"first_name" = $1,
				"last_name" = $2,
				"branch_id" = $3,  
				"phone" = $4,
				"login"=$5,
				"password"=$6,
				"updated_at" = NOW() 
				WHERE id = $7`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Firstname,
		req.Lastname,
		req.BranchId,
		req.Phone,
		req.Login,
		hashedPass,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update users: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("users with ID %d not found", req.Id)
	}

	return fmt.Sprintf("users with ID %d updated", req.Id), nil
}

func (b *userRepo) Delete(c context.Context, req *user_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "users" 
				SET 
				"active" = false,
				"deleted_at" = NOW() 
				WHERE id = $1  and "active"`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to delete users: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("users with ID %s not found", req.Id)
	}

	return "deleted", nil
}

func (b *userRepo) GetByLogin(c context.Context, req *user_service.IdRequest) (resp *user_service.Users, err error) {
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
		created_at,
		updated_at 
    FROM users 
    WHERE login = $1 AND "active" AND "deleted_at" IS NULL;`

	user := user_service.Users{}

	err = b.db.QueryRow(c, query, req.Id).Scan(
		&user.Id,
		&user.Firstname,
		&user.Lastname,
		&user.BranchId,
		&user.Phone,
		&user.Login,
		&user.Active,
		&user.Password,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if createdAt.Valid {
		user.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.String
	}

	return &user, nil
}
