package postgres

import (
	"context"
	"fmt"
	"order_service/config"
	"order_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db             *pgxpool.Pool
	order          *orderRepo
	deliveryTariff *tariffRepo
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

func (d *strg) Order() storage.OrderI {
	if d.order == nil {
		d.order = NewOrder(d.db)
	}
	return d.order
}

func (d *strg) DeliveryTariff() storage.DeliveryTariffI {
	if d.deliveryTariff == nil {
		d.deliveryTariff = NewDeliveryTariff(d.db)
	}
	return d.deliveryTariff
}
