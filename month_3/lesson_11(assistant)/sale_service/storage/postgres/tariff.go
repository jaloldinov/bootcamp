package postgres

import (
	"context"
	"fmt"
	"staff_service/pkg/helper"
	"time"

	pb "staff_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type tariffRepo struct {
	db *pgxpool.Pool
}

func NewTariff(db *pgxpool.Pool) *tariffRepo {
	return &tariffRepo{
		db: db,
	}
}

func (b *tariffRepo) CreateTariff(c context.Context, req *pb.CreateTariffRequest) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "tariffs"(
			"id", 
			"name", 
			"type",
			"amount_for_cash",
			"amount_for_card",
			"created_at")
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.TariffType,
		req.AmountForCash,
		req.AmountForCard,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create tariff: %w", err)
	}

	return id, nil
}

func (b *tariffRepo) GetTariff(c context.Context, req *pb.IdRequest) (resp *pb.Tariff, err error) {
	query := `
			SELECT 
				"id", 
				"name", 
				"type",
				"amount_for_cash",
				"amount_for_card",
				"created_at",
				"updated_at" 
			FROM "tariffs" 
			WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt time.Time
	)

	tariff := pb.Tariff{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&tariff.Id,
		&tariff.Name,
		&tariff.TariffType,
		&tariff.AmountForCash,
		&tariff.AmountForCard,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("tariff not found")
		}
		return nil, fmt.Errorf("failed to get tariff: %w", err)
	}

	tariff.CreatedAt = createdAt.Format(time.RFC3339)
	tariff.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &tariff, nil
}

func (b *tariffRepo) GetAllTariff(c context.Context, req *pb.ListTariffRequest) (*pb.ListTariffResponse, error) {
	var (
		resp   pb.ListTariffResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM "tariffs" WHERE true ` + filter

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
				"name", 
				"type",
				"amount_for_cash",
				"amount_for_card",
				"created_at",
				"updated_at" 
			FROM "tariffs" 
			WHERE true` + filter

	query += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(c, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	defer rows.Close()

	createdAt := time.Time{}
	updatedAt := time.Time{}
	for rows.Next() {
		var tariff pb.Tariff

		err = rows.Scan(
			&tariff.Id,
			&tariff.Name,
			&tariff.TariffType,
			&tariff.AmountForCash,
			&tariff.AmountForCard,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning tariff err: %w", err)
		}

		tariff.CreatedAt = createdAt.Format(time.RFC3339)
		tariff.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.Tariffs = append(resp.Tariffs, &tariff)
	}

	return &resp, nil
}

func (b *tariffRepo) UpdateTariff(c context.Context, req *pb.UpdateTariffRequest) (string, error) {

	query := `
				UPDATE "tariffs" 
				SET 
				"name" = $1, 
				"type" = $2,
				"amount_for_cash" = $3,
				"amount_for_card" = $4,
				"updated_at" = NOW() 
				WHERE id = $5 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.TariffType,
		req.AmountForCash,
		req.AmountForCard,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update tariff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("tariff with ID %s not found", req.Id)
	}

	return fmt.Sprintf("tariff with ID %s updated", req.Id), nil
}

func (b *tariffRepo) DeleteTariff(c context.Context, req *pb.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM "tariffs" 
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
		return "", fmt.Errorf("tariff with ID %s not found", req.Id)

	}

	return fmt.Sprintf("tariff with ID %s deleted", req.Id), nil
}
