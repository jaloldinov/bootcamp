package postgres

import (
	"context"
	"fmt"
	"staff_service/config"
	"staff_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db     *pgxpool.Pool
	tariff *tariffRepo
	staff  *staffRepo
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

func (d *strg) Staff() storage.StaffI {
	if d.staff == nil {
		d.staff = NewStaff(d.db)
	}
	return d.staff
}

func (d *strg) Tariff() storage.TariffI {
	if d.tariff == nil {
		d.tariff = NewTariff(d.db)
	}
	return d.tariff
}
