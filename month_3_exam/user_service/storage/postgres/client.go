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

type clientRepo struct {
	db *pgxpool.Pool
}

func NewClient(db *pgxpool.Pool) *clientRepo {
	return &clientRepo{
		db: db,
	}
}

func (b *clientRepo) Create(c context.Context, req *user_service.CreateClientsRequest) (*user_service.Response, error) {
	query := `
        INSERT INTO clients (
            first_name,
            last_name,
            photo,  
            phone,
            birth_date,
            discount_type,
            discount_amount,
            total_orders_count,
            total_orders_sum,
			last_ordered_date,
			created_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now()
        ) RETURNING id`

	var id int
	err := b.db.QueryRow(c, query,
		req.Firstname,
		req.Lastname,
		req.Photo,
		req.Phone,
		req.BirthDate,
		req.DiscountType,
		req.DiscountAmount,
		req.TotalOrdersCount,
		req.TotalOrdersSum,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to create Client: %w", err)
	}

	return &user_service.Response{Message: fmt.Sprintf("created with ID: %d", id)}, nil
}

func (b *clientRepo) Get(c context.Context, req *user_service.IdRequest) (resp *user_service.Clients, err error) {
	var (
		createdAt sql.NullString
		updatedAt sql.NullString
	)
	query := `
    SELECT 
		id,
		first_name,
		last_name,
		photo,  
		phone,
		birth_date::text,
		last_ordered_date::text,
		total_orders_count,
		total_orders_sum,
		discount_type,
		discount_amount,
		created_at,
		updated_at 
    FROM clients 
    WHERE id = $1 AND "deleted_at" IS NULL;`

	client := user_service.Clients{}

	err = b.db.QueryRow(c, query, req.Id).Scan(
		&client.Id,
		&client.Firstname,
		&client.Lastname,
		&client.Photo,
		&client.Phone,
		&client.BirthDate,
		&client.LastOrderedDate,
		&client.TotalOrdersCount,
		&client.TotalOrdersSum,
		&client.DiscountType,
		&client.DiscountAmount,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("client not found")
		}
		return nil, fmt.Errorf("failed to get client: %w", err)
	}
	if createdAt.Valid {
		client.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		client.UpdatedAt = updatedAt.String
	}

	return &client, nil
}

func (b *clientRepo) GetList(c context.Context, req *user_service.ListClientsRequest) (*user_service.ListClientsResponse, error) {

	var (
		resp   user_service.ListClientsResponse
		err    error
		filter string
		params = make(map[string]interface{})
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

	if req.DiscountType != "" {
		filter += " AND discount_type = :discount_type"
		params["discount_type"] = req.DiscountType
	}

	if req.LastOrderedDateFrom != "" {
		filter += " AND last_ordered_date >= :last_ordered_date_from"
		params["last_ordered_date_from"] = req.LastOrderedDateFrom
	}

	if req.LastOrderedDateTo != "" {
		filter += " AND last_ordered_date <= :last_ordered_date_to"
		params["last_ordered_date_to"] = req.LastOrderedDateTo
	}

	if req.TotalOrdersSumFrom != 0 {
		filter += " AND total_orders_sum >= :total_orders_sum_from"
		params["total_orders_sum_from"] = req.TotalOrdersSumFrom
	}

	if req.TotalOrdersSumTo != 0 {
		filter += " AND total_orders_sum <= :total_orders_sum_to"
		params["total_orders_sum_to"] = req.TotalOrdersSumTo
	}

	if req.TotalOrdersCountFrom != 0 {
		filter += " AND total_orders_count >= :total_orders_count_from"
		params["total_orders_count_from"] = req.TotalOrdersCountFrom
	}

	if req.TotalOrdersCountTo != 0 {
		filter += " AND total_orders_count <= :total_orders_count_to"
		params["total_orders_count_to"] = req.TotalOrdersCountTo
	}

	if req.DiscountAmountFrom != "" {
		filter += " AND discount_amount >= :discount_amount_from"
		params["discount_amount_from"] = req.DiscountAmountFrom
	}

	if req.DiscountAmountTo != "" {
		filter += " AND discount_amount <= :discount_amount_to"
		params["discount_amount_to"] = req.DiscountAmountTo
	}

	if req.CreatedAtFrom != "" {
		filter += " AND created_at >= :created_at_from"
		params["created_at_from"] = req.CreatedAtFrom
	}

	if req.CreatedAtTo != "" {
		filter += " AND created_at <= :created_at_to"
		params["created_at_to"] = req.CreatedAtTo
	}

	countQuery := `SELECT count(1) FROM clients WHERE deleted_at IS NULL` + filter
	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(&resp.Count)
	if err != nil {
		return nil, fmt.Errorf("error while scanning count: %w", err)
	}

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			photo,
			phone,
			birth_date::text,
			discount_type,
			discount_amount,
			total_orders_count,
			total_orders_sum,
			last_ordered_date::text,
			created_at,
			updated_at
		FROM clients
		WHERE deleted_at IS NULL` + filter +
		` ORDER BY created_at DESC LIMIT :limit OFFSET :offset`
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
		return nil, fmt.Errorf("error while getting rows: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			client    user_service.Clients
			createdAt sql.NullString
			updatedAt sql.NullString
		)
		err = rows.Scan(
			&client.Id,
			&client.Firstname,
			&client.Lastname,
			&client.Photo,
			&client.Phone,
			&client.BirthDate,
			&client.DiscountType,
			&client.DiscountAmount,
			&client.TotalOrdersCount,
			&client.TotalOrdersSum,
			&client.LastOrderedDate,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning client: %w", err)
		}

		if createdAt.Valid {
			client.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			client.UpdatedAt = updatedAt.String
		}

		resp.Clients = append(resp.Clients, &client)
	}

	return &resp, nil
}

func (b *clientRepo) Update(c context.Context, req *user_service.UpdateClientsRequest) (string, error) {

	query := `
				UPDATE "clients" 
				SET 
				"first_name" = $1,
				"last_name" = $2,
				"photo" = $3,  
				"phone" = $4,
				"birth_date"=$5,
				"discount_type"=$6,
				"discount_amount"=$7,
				"updated_at" = NOW() 
				WHERE id = $8`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Firstname,
		req.Lastname,
		req.Photo,
		req.Phone,
		req.BirthDate,
		req.DiscountType,
		req.DiscountAmount,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update client: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("client with ID %d not found", req.Id)
	}

	return fmt.Sprintf("client with ID %d updated", req.Id), nil
}

func (b *clientRepo) Delete(c context.Context, req *user_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "clients" 
				SET 
				"deleted_at" = NOW() 
				WHERE id = $1 AND "deleted_at" IS NULL`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to delete client: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("client with ID %s not found", req.Id)
	}

	return "deleted", nil
}

func (b *clientRepo) UpdateOrder(c context.Context, req *user_service.UpdateClientsOrderRequest) (string, error) {
	query := `
		UPDATE "clients" 
		SET 
		"last_ordered_date" = NOW(),
		"total_orders_sum" = "total_orders_sum" + $1,
		"total_orders_count" = "total_orders_count" + $2,
		"updated_at" = NOW() 
		WHERE id = $3`

	result, err := b.db.Exec(
		c,
		query,
		req.TotalOrdersSum,
		req.TotalOrdersCount,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update client order: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == -1 {
		return "", fmt.Errorf("client with ID %d not found", req.Id)
	}

	return fmt.Sprintf("client with ID %d updated", req.Id), nil
}
