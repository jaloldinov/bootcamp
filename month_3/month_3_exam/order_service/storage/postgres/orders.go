package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"order_service/pkg/helper"

	order_service "order_service/genproto"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrder(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (b *orderRepo) UpdateStatus(c context.Context, req *order_service.UpdateOrderStatusRequest) (string, error) {
	// Get the previous status
	prevStatus, err := b.GetOrderStatus(c, int(req.Id), req.OrderId)
	if err != nil {
		return "", fmt.Errorf("failed to get previous status: %w", err)
	}

	// Check if the previous status matches
	switch req.Status {
	case "courier_accepted":
		if prevStatus != "accepted" {
			return "", fmt.Errorf("previous status must be 'accepted' to update to 'courier_accepted'")
		}
	case "ready_in_branch":
		if prevStatus != "courier_accepted" {
			return "", fmt.Errorf("previous status must be 'ready_in_branch' to update to 'courier_accepted'")
		}
	case "on_way":
		if prevStatus != "courier_accepted" {
			return "", fmt.Errorf("previous status must be 'ready_in_branch' to update to 'on_way'")
		}
	case "finished":
		if prevStatus != "on_way" {
			return "", fmt.Errorf("previous status must be 'on_way' to update to 'finished'")
		}
	}

	query := `
		UPDATE "orders" 
		SET "status" = $1, "updated_at" = NOW()
		WHERE id = $2 AND "order_id" = $3 AND "deleted_at" IS NULL`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Status,
		req.Id,
		req.OrderId,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update status: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("order with ID %d not found", req.Id)
	}

	return fmt.Sprintf("status changed from '%s' to '%s'", prevStatus, req.Status), nil
}

func (b *orderRepo) GetListByCourierId(c context.Context, req *order_service.IdRequest) (*order_service.ListOrderResponse, error) {
	var (
		resp   order_service.ListOrderResponse
		err    error
		filter string = ` WHERE deleted_at IS NULL AND status IN ('courier_accepted', 'ready_in_branch', 'on_way')  `
		params        = make(map[string]interface{})
	)
	if req.Id != "" {
		filter += " AND courier_id = :courier_id"
		params["courier_id"] = req.Id
	}

	countQuery := `SELECT count(1) FROM "orders"  ` + filter

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
			"order_id", 
			"client_id",    
			"branch_id", 
			"type",
			"address",
			"courier_id",
			"price",
			"delivery_price",
			"discount",
			"status",
			"payment_type",
			"created_at",
			"updated_at" 
			FROM "orders"  ` + filter

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(c, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	defer rows.Close()

	var createdAt sql.NullString
	var updatedAt sql.NullString

	for rows.Next() {
		var order order_service.Order

		err = rows.Scan(
			&order.Id,
			&order.OrderId,
			&order.ClientId,
			&order.BranchId,
			&order.Type,
			&order.Address,
			&order.CourierId,
			&order.Price,
			&order.DeliveryPrice,
			&order.Discount,
			&order.Status,
			&order.PaymentType,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning order err: %w", err)
		}

		if createdAt.Valid {
			order.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			order.UpdatedAt = createdAt.String
		}
		resp.Orders = append(resp.Orders, &order)
	}

	return &resp, nil
}

func (b *orderRepo) Create(c context.Context, req *order_service.CreateOrderRequest) (string, error) {
	order_id := helper.GenerateUniqueID()

	query := `
		INSERT INTO "orders"(
			"order_id", 
			"client_id",    
			"branch_id", 
			"type",
			"address",
			"courier_id",
			"price",
			"delivery_price",
			"discount",
			"status",
			"payment_type",
			"created_at"
			)
		VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9,  'accepted', $10,NOW()) RETURNING "order_id"
	`

	var orderID string
	err := b.db.QueryRow(c, query,
		order_id,
		req.ClientId,
		req.BranchId,
		req.Type,
		req.Address,
		req.CourierId,
		req.Price,
		req.DeliveryPrice,
		req.Discount,
		req.PaymentType,
	).Scan(&orderID)
	if err != nil {
		return "", fmt.Errorf("failed to create order: %w", err)
	}

	return fmt.Sprintf("created with orderID: %s", orderID), nil

}

func (b *orderRepo) Get(c context.Context, req *order_service.IdRequest) (resp *order_service.Order, err error) {
	query := `
		SELECT 
			"id",
			"order_id", 
			"client_id",    
			"branch_id", 
			"type",
			"address",
			"courier_id",
			"price",
			"delivery_price",
			"discount",
			"status",
			"payment_type",
			"created_at",
			"updated_at" 
		FROM "orders" 
		WHERE "order_id"=$1 AND "deleted_at" IS NULL `

	var (
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	order := order_service.Order{}
	err = b.db.QueryRow(c, query, &req.Id).Scan(
		&order.Id,
		&order.OrderId,
		&order.ClientId,
		&order.BranchId,
		&order.Type,
		&order.Address,
		&order.CourierId,
		&order.Price,
		&order.DeliveryPrice,
		&order.Discount,
		&order.Status,
		&order.PaymentType,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	if createdAt.Valid {
		order.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		order.UpdatedAt = updatedAt.String
	}

	return &order, nil
}

func (b *orderRepo) GetList(c context.Context, req *order_service.ListOrderRequest) (*order_service.ListOrderResponse, error) {
	var (
		resp   order_service.ListOrderResponse
		err    error
		filter string = ` WHERE deleted_at IS NULL  `
		params        = make(map[string]interface{})
	)

	if req.OrderId != "" {
		filter += " AND order_id = :order_id"
		params["order_id"] = req.OrderId
	}

	if req.ClientId != 0 {
		filter += " AND client_id = :client_id"
		params["client_id"] = req.ClientId
	}

	if req.BranchId != 0 {
		filter += " AND branch_id = :branch_id"
		params["branch_id"] = req.BranchId
	}

	if req.CourierId != 0 {
		filter += " AND courier_id = :courier_id"
		params["courier_id"] = req.CourierId
	}

	if req.PaymentType != "" {
		filter += " AND payment_type = :payment_type"
		params["payment_type"] = req.PaymentType
	}

	if req.DeliveryType != "" {
		filter += " AND type = :delivery_type"
		params["delivery_type"] = req.DeliveryType
	}

	if req.PriceFrom != 0 {
		filter += " AND price >= :from_price"
		params["from_price"] = req.PriceFrom
	}

	if req.PriceTo != 0 {
		filter += " AND price <= :to_price"
		params["to_price"] = req.PriceTo
	}

	countQuery := `SELECT count(1) FROM "orders"  ` + filter

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
			"order_id", 
			"client_id",    
			"branch_id", 
			"type",
			"address",
			"courier_id",
			"price",
			"delivery_price",
			"discount",
			"status",
			"payment_type",
			"created_at",
			"updated_at" 
			FROM "orders"  ` + filter

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

	var createdAt sql.NullString
	var updatedAt sql.NullString

	for rows.Next() {
		var order order_service.Order

		err = rows.Scan(
			&order.Id,
			&order.OrderId,
			&order.ClientId,
			&order.BranchId,
			&order.Type,
			&order.Address,
			&order.CourierId,
			&order.Price,
			&order.DeliveryPrice,
			&order.Discount,
			&order.Status,
			&order.PaymentType,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning order err: %w", err)
		}

		if createdAt.Valid {
			order.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			order.UpdatedAt = createdAt.String
		}
		resp.Orders = append(resp.Orders, &order)
	}

	return &resp, nil
}

func (b *orderRepo) Update(c context.Context, req *order_service.UpdateOrderRequest) (string, error) {

	query := `
				UPDATE "orders" 
				SET 
				"client_id" = $1,   
				"branch_id" = $2,
				"type" = $3,
				"address" = $4,
				"courier_id" = $5,
				"price" = $6,
				"delivery_price" = $7,
				"discount" = $8,
				"status" = $9,
				"updated_at" = NOW()
				WHERE id = $10  AND "order_id" = $11 AND "deleted_at" IS NULL`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.ClientId,
		req.BranchId,
		req.Type,
		req.Address,
		req.CourierId,
		req.Price,
		req.DeliveryPrice,
		req.Discount,
		req.Status,
		req.Id,
		req.OrderId,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update order: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("order with ID %d not found", req.Id)
	}

	return fmt.Sprintf("order with ID %d updated", req.Id), nil
}

func (b *orderRepo) Delete(c context.Context, req *order_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "orders" 
				SET 
				"deleted_at" = NOW() 
				WHERE id = $1  AND "deleted_at" IS NULL`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to delete order: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("order with ID %s not found", req.Id)
	}

	return "deleted", nil
}

func (b *orderRepo) GetOrderStatus(c context.Context, ID int, orderID string) (string, error) {
	query := `
		SELECT "status" FROM "orders"
		WHERE id = $1 AND order_id = $2 AND "deleted_at" IS NULL`

	var status string
	err := b.db.QueryRow(c, query, ID, orderID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("order with ID %s not found", orderID)
		}
		return "", fmt.Errorf("failed to get order status: %w", err)
	}

	return status, nil
}
