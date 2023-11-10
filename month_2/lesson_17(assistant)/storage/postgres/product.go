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

type productRepo struct {
	db *pgxpool.Pool
}

func NewPorductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(req *models.CreateProduct) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO product(id, name, price, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`

	_, err := p.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		helper.NewNullString(req.Category_id),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *productRepo) GetByID(req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query string

		id          sql.NullString
		name        sql.NullString
		price       sql.NullInt32
		category_id sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			price,
			parent_id,
			created_at,
			updated_at
		FROM product
		WHERE id = $1
	`

	err := r.db.QueryRow(context.Background(), query, req.Id).Scan(
		&id,
		&name,
		&price,
		&category_id,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Product{
		Id:          id.String,
		Name:        name.String,
		Price:       int(price.Int32),
		Category_id: category_id.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *productRepo) GetList(req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {

	var (
		resp   = &models.ProductGetListResponse{}
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
			id          sql.NullString
			name        sql.NullString
			price       sql.NullInt32
			category_id sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&price,
			&category_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &models.Product{
			Id:          id.String,
			Name:        name.String,
			Price:       int(price.Int32),
			Category_id: category_id.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *productRepo) Update(req *models.UpdateProduct) (string, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			product
		SET
			name = :name,
			price = :price,
			category_id = :category_id,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"price":       req.Price,
		"category_id": helper.NewNullString(req.Category_id),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		return "", err
	}

	return req.Id, nil
}

func (r *productRepo) Delete(req *models.ProductPrimaryKey) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM product WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
