package postgres

import (
	"catalog_service/config"
	"catalog_service/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db       *pgxpool.Pool
	category *categoryRepo
	product  *productRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}

	return &strg{
		db: pool,
	}, nil
}

func (d *strg) Category() storage.CategoryI {
	if d.category == nil {
		d.category = NewCategory(d.db)
	}
	return d.category
}

func (d *strg) Product() storage.ProductI {
	if d.product == nil {
		d.product = NewProduct(d.db)
	}
	return d.product
}
