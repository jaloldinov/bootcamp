package postgres

import (
	"context"
	"fmt"
	order_service "order_service/genproto"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OrderProductRepo struct {
	db *pgxpool.Pool
}

func NewOrderProductRepo(db *pgxpool.Pool) *OrderProductRepo {
	return &OrderProductRepo{
		db: db,
	}
}

func (r *OrderProductRepo) Create(ctx context.Context, req *order_service.OrderProductCreateReq) (*order_service.OrderProductCreateResp, error) {
	var id int
	query := `
	INSERT INTO order_products(
		order_id,
		product_id,
		quantity,
		price
	) VALUES($1,$2,$3,$4)
	RETURNING id;`

	if err := r.db.QueryRow(ctx, query,
		req.OrderId,
		req.ProductId,
		req.Quantity,
		req.Price,
	).Scan(&id); err != nil {
		return nil, err
	}

	return &order_service.OrderProductCreateResp{Msg: "success"}, nil
}

func (r *OrderProductRepo) GetList(ctx context.Context, req *order_service.OrderProductGetListReq) (*order_service.OrderProductGetListResp, error) {
	var (
		filter  = " WHERE deleted_at IS NULL "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		order_id,
		product_id,
		quantity,
		price
	FROM order_products `

	if req.OrderId != "" {
		filter += " AND order_id=" + req.OrderId + " "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM order_products` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &order_service.OrderProductGetListResp{}
	for rows.Next() {
		var orderProduct = order_service.OrderProduct{}
		if err := rows.Scan(
			&orderProduct.OrderId,
			&orderProduct.ProductId,
			&orderProduct.Quantity,
			&orderProduct.Price,
		); err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &orderProduct)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *OrderProductRepo) GetById(ctx context.Context, req *order_service.OrderProductIdReq) (*order_service.OrderProduct, error) {
	query := `
	SELECT 
		order_id,
		product_id,
		quantity,
		price
	FROM order_products
	WHERE order_id = $1 AND product_id = $2;`

	var orderProduct = order_service.OrderProduct{}
	if err := r.db.QueryRow(ctx, query, req.OrderId, req.ProductId).Scan(
		&orderProduct.OrderId,
		&orderProduct.ProductId,
		&orderProduct.Quantity,
		&orderProduct.Price,
	); err != nil {
		return nil, err
	}

	return &orderProduct, nil
}

func (r *OrderProductRepo) Update(ctx context.Context, req *order_service.OrderProductUpdateReq) (*order_service.OrderProductUpdateResp, error) {
	query := `
	UPDATE order_products
	SET 
		quantity=$3
	WHERE order_id = $1 AND product_id = $2;`

	resp, err := r.db.Exec(ctx, query,
		req.OrderId,
		req.ProductId,
		req.Quantity,
	)

	if err != nil {
		return nil, err
	}

	if resp.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}

	return &order_service.OrderProductUpdateResp{Msg: "success"}, nil
}

func (r *OrderProductRepo) Delete(ctx context.Context, req *order_service.OrderProductIdReq) (*order_service.Response, error) {
	query := `
	DELETE FROM order_products
	WHERE order_id = $1 AND product_id = $2;`

	_, err := r.db.Exec(ctx, query, req.OrderId, req.ProductId)
	if err != nil {
		return nil, err
	}

	return &order_service.Response{Message: "success"}, nil
}
