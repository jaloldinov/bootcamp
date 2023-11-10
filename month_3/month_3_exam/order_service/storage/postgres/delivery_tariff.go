package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"order_service/pkg/helper"

	tariff_service "order_service/genproto"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type tariffRepo struct {
	db *pgxpool.Pool
}

func NewDeliveryTariff(db *pgxpool.Pool) *tariffRepo {
	return &tariffRepo{
		db: db,
	}
}

func (b *tariffRepo) Create(c context.Context, req *tariff_service.CreateDeliveryTariffRequest) (string, error) {
	var tariffID int

	if req.TariffType == "fixed" {
		query := `
		INSERT INTO "delivery_tarif"(
			"name",    
			"type", 
			"base_price",
			"created_at"
			)
		VALUES ($1, $2, $3,  NOW()) RETURNING "id"
	`

		err := b.db.QueryRow(c, query,
			req.Name,
			req.TariffType,
			req.BasePrice,
		).Scan(&tariffID)
		if err != nil {
			return "", fmt.Errorf("failed to create fixed tariff: %w", err)
		}

		return fmt.Sprintf("created fixed tariffID: %d", tariffID), nil

	} else if req.TariffType == "alternative" {
		query := `
		INSERT INTO "delivery_tarif"(
			"name",    
			"type", 
			"created_at"
			)
		VALUES ($1, $2,  NOW()) RETURNING "id"
	`
		err := b.db.QueryRow(c, query,
			req.Name,
			req.TariffType,
		).Scan(&tariffID)
		if err != nil {
			return "", fmt.Errorf("failed to create alternative tariff: %w", err)
		}

		query2 := `
			INSERT INTO "delivery_tarif_values"(
				"delivery_tarif_id",
				"from_price",
				"to_price",
				"price"
			) VALUES ($1, $2, $3, $4)
		`

		_, err = b.db.Exec(c, query2,
			tariffID,
			req.Values.FromPrice,
			req.Values.ToPrice,
			req.Values.Price,
		)
		if err != nil {
			return "", fmt.Errorf("failed to create alternative tariff values: %w", err)
		}
		return fmt.Sprintf("created alternative tariffID: %d", tariffID), nil
	}
	return "", fmt.Errorf("something went wrong")
}

func (b *tariffRepo) Get(c context.Context, req *tariff_service.IdRequest) (resp *tariff_service.DeliveryTariff, err error) {
	query := `
	SELECT 
 	   t."id",
	   t."name", 
  	   t."type",    
		COALESCE(t."base_price"::numeric, 0) AS base_price,
		COALESCE(tv."from_price"::numeric, 0) AS from_price,
		COALESCE(tv."to_price"::numeric, 0) AS to_price,
		COALESCE(tv."price"::numeric, 0) AS price,
	    t."created_at",
   		t."updated_at" 
	FROM "delivery_tarif" t
	FULL JOIN "delivery_tarif_values" tv ON t.id = tv.delivery_tarif_id 
	WHERE t."id" = $1 AND t."deleted_at" IS NULL;`

	var (
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	tariff := tariff_service.DeliveryTariff{}
	tariff.Values = &tariff_service.DeliveryTariffValues{}
	err = b.db.QueryRow(c, query, &req.Id).Scan(
		&tariff.Id,
		&tariff.Name,
		&tariff.TariffType,
		&tariff.BasePrice,
		&tariff.Values.FromPrice,
		&tariff.Values.ToPrice,
		&tariff.Values.Price,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("tariff not found")
		}
		return nil, fmt.Errorf("failed to get tariff: %w", err)
	}

	if createdAt.Valid {
		tariff.CreatedAt = createdAt.String
	}
	if updatedAt.Valid {
		tariff.UpdatedAt = updatedAt.String
	}

	return &tariff, nil
}

func (b *tariffRepo) GetList(c context.Context, req *tariff_service.ListDeliveryTariffRequest) (*tariff_service.ListDeliveryTariffResponse, error) {
	var (
		err    error
		filter string = ` WHERE t.deleted_at IS NULL  `
		params        = make(map[string]interface{})
	)
	resp := tariff_service.ListDeliveryTariffResponse{
		DeliveryTariffs: make([]*tariff_service.DeliveryTariff, 0),
		Count:           0,
	}

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :name || '%' "
		params["name"] = req.Search
	}

	if req.TarifType != "" {
		filter += " AND type = :type"
		params["type"] = req.TarifType
	}

	countQuery := `SELECT count(1) FROM "delivery_tarif" t ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `
	SELECT 
 	   t."id",
	   t."name", 
  	   t."type",    
		COALESCE(t."base_price"::numeric, 0) AS base_price,
		COALESCE(tv."from_price"::numeric, 0) AS from_price,
		COALESCE(tv."to_price"::numeric, 0) AS to_price,
		COALESCE(tv."price"::numeric, 0) AS price,
	    t."created_at",
   		t."updated_at" 
	FROM "delivery_tarif" t
	FULL JOIN "delivery_tarif_values" tv ON t.id = tv.delivery_tarif_id ` + filter

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
		var tariff tariff_service.DeliveryTariff
		tariff.Values = &tariff_service.DeliveryTariffValues{}

		err = rows.Scan(
			&tariff.Id,
			&tariff.Name,
			&tariff.TariffType,
			&tariff.BasePrice,
			&tariff.Values.FromPrice,
			&tariff.Values.ToPrice,
			&tariff.Values.Price,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning tariff err: %w", err)
		}

		if createdAt.Valid {
			tariff.CreatedAt = createdAt.String

		}
		if updatedAt.Valid {
			tariff.UpdatedAt = createdAt.String
		}
		resp.DeliveryTariffs = append(resp.DeliveryTariffs, &tariff)
	}

	return &resp, nil
}

func (b *tariffRepo) Update(c context.Context, req *tariff_service.UpdateDeliveryTariffRequest) (string, error) {

	if req.TariffType == "fixed" {
		query := `
			UPDATE "delivery_tarif" 
			SET 
			"name" = $1,   
			"type" = $2,
			"base_price" = $3,
			"updated_at" = NOW()
			WHERE id = $4 AND "deleted_at" IS NULL`

		result, err := b.db.Exec(
			context.Background(),
			query,
			req.Name,
			req.TariffType,
			req.BasePrice,
			req.Id,
		)

		if err != nil {
			return "", fmt.Errorf("failed to update tariff: %w", err)
		}

		if result.RowsAffected() == 0 {
			return "", fmt.Errorf("tariff with ID %d not found", req.Id)
		}

		return fmt.Sprintf("tariff with ID %d updated", req.Id), nil

	} else if req.TariffType == "alternative" {
		query := `
		UPDATE "delivery_tarif" 
		SET 
		"name" = $1,   
		"type" = $2,
		"updated_at" = NOW()
		WHERE id = $3 AND "deleted_at" IS NULL`

		result, err := b.db.Exec(
			context.Background(),
			query,
			req.Name,
			req.TariffType,
			req.Id,
		)

		if err != nil {
			return "", fmt.Errorf("failed to update tariff: %w", err)
		}
		if result.RowsAffected() == 0 {
			return "", fmt.Errorf("tariff with ID %d not found", req.Id)
		}

		query2 := `
			UPDATE "delivery_tarif_values" 
			SET 
			"from_price" = $1,   
			"to_price" = $2,
			"price" = $3
		WHERE "delivery_tarif_id" = $4 
		`
		result, err = b.db.Exec(
			context.Background(),
			query2,
			req.Values.FromPrice,
			req.Values.ToPrice,
			req.Values.Price,
			req.Id,
		)

		if err != nil {
			return "", fmt.Errorf("failed to update tariff values: %w", err)
		}

		if result.RowsAffected() == 0 {
			return "", fmt.Errorf("tariff values with ID %d not found", req.Id)
		}

		return fmt.Sprintf("tariff and values with ID %d updated", req.Id), nil
	}
	return "", fmt.Errorf("someting went wrong")
}

func (b *tariffRepo) Delete(c context.Context, req *tariff_service.IdRequest) (resp string, err error) {

	query := `
				UPDATE "delivery_tarif" 
				SET 
				"deleted_at" = NOW() 
				WHERE id = $1  AND "deleted_at" IS NULL`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to delete tariff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("tariff with ID %s not found", req.Id)
	}

	return "deleted", nil
}
