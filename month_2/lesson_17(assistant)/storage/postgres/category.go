package postgres

import (
	"app/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(req *models.CreateCategory) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO category(id, title, parent_id, updated_at)
		VALUES ($1, $2, $3, NOW())
	`

	_, err := r.db.Exec(context.Background(), query,
		id,
		req.Title,
		helper.NewNullString(req.ParentID),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *categoryRepo) GetByID(req *models.CategoryPrimaryKey) (*models.Category, error) {

	var (
		query string

		id        sql.NullString
		title     sql.NullString
		parentId  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			title,
			parent_id,
			created_at,
			updated_at
		FROM category
		WHERE id = $1
	`

	err := r.db.QueryRow(context.Background(), query, req.Id).Scan(
		&id,
		&title,
		&parentId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:        id.String,
		Title:     title.String,
		ParentID:  parentId.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *categoryRepo) GetList(req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error) {

	var (
		resp   = &models.CategoryGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			title,
			parent_id,
			created_at,
			updated_at
		FROM category
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	fmt.Println(query)

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			title     sql.NullString
			parentId  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&title,
			&parentId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &models.Category{
			Id:        id.String,
			Title:     title.String,
			ParentID:  parentId.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *categoryRepo) Update(req *models.UpdateCategory) (string, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			category
		SET
			title = :title,
			parent_id = :parent_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":        req.Id,
		"title":     req.Title,
		"parent_id": helper.NewNullString(req.ParentID),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		return "", err
	}

	return req.Id, nil
}

func (r *categoryRepo) Delete(req *models.CategoryPrimaryKey) error {

	_, err := r.db.Exec(context.Background(), "DELETE FROM category WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
