package postgres

import (
	"auth/models"
	"auth/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (b *userRepo) CreateUser(c context.Context, req *models.CreateUser) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO users(
			"id", 
			"login", 
			"password", 
			"name", 
			"age", 
			"phone_number",
			"created_at")
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Login,
		req.Password,
		req.Name,
		req.Age,
		req.PhoneNumber,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (b *userRepo) GetUser(c context.Context, req *models.IdRequest) (resp *models.User, err error) {
	var (
		updatedAt sql.NullString
		createdAt sql.NullString
	)
	query := `
			SELECT 
				"id", 
				"login", 
				"password", 
				"name", 
				"age", 
				"phone_number",
				"created_at",
				"updated_at" 
			FROM users 
				WHERE "id"=$1`

	user := models.User{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Name,
		&user.Age,
		&user.PhoneNumber,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	user.CreatedAt = createdAt.String
	user.UpdatedAt = updatedAt.String
	return &user, nil
}

func (b *userRepo) UpdateUser(c context.Context, req *models.User) (string, error) {

	query := `
			UPDATE users 
				SET 
				"login" = $1,
				"password" = $2,
				"name" = $3, 
				"age" = $4,  
				"phone_number" = $5,
				"updated_at" = NOW() 
				WHERE "id" = $6 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Login,
		req.Password,
		req.Name,
		req.Age,
		req.PhoneNumber,
		req.ID,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("user with ID %s not found", req.ID)
	}

	return req.ID, nil
}

func (b *userRepo) GetAllUser(c context.Context, req *models.GetAllUserRequest) (*models.GetAllUser, error) {
	params := make(map[string]interface{})
	var resp = &models.GetAllUser{}

	resp.Users = make([]models.User, 0)

	filter := " WHERE true "
	query := `
			SELECT
				COUNT(*) OVER(),
				"id", 
				"login", 
				"password", 
				"name", 
				"age", 
				"phone_number",
				"created_at",
				"updated_at" 
			FROM users
		`
	if req.Name != "" {
		filter += ` AND "name" ILIKE '%' || :name || '%' `
		params["name"] = req.Name
	}

	offset := (req.Page - 1) * req.Limit
	params["limit"] = req.Limit
	params["offset"] = offset

	query = query + filter + " ORDER BY created_at DESC OFFSET :offset LIMIT :limit "
	rquery, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := b.db.Query(context.Background(), rquery, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			createdAt sql.NullString
			updatedAt sql.NullString
		)
		user := models.User{}

		err := rows.Scan(
			&resp.Count,
			&user.ID,
			&user.Login,
			&user.Password,
			&user.Name,
			&user.Age,
			&user.PhoneNumber,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		user.CreatedAt = createdAt.String
		user.UpdatedAt = updatedAt.String

		resp.Users = append(resp.Users, user)
	}
	return resp, nil

}

func (b *userRepo) DeleteUser(c context.Context, req *models.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM users 
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
		return "", fmt.Errorf("user with ID %s not found", req.Id)

	}

	return req.Id, nil
}

func (b *userRepo) GetByLogin(c context.Context, req *models.LoginRequest) (resp *models.LoginDataRespond, err error) {

	query := `
			SELECT 
				"login", 
				"password", 
				"phone_number"
			FROM users 
				WHERE "login"=$1`

	user := models.LoginDataRespond{}
	err = b.db.QueryRow(context.Background(), query, req.Login).Scan(
		&user.Login,
		&user.Password,
		&user.PhoneNumber,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
