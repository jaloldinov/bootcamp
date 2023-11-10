package postgres

import (
	"context"
	"fmt"
	"user_service/config"
	"user_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db      *pgxpool.Pool
	branch  *branchRepo
	user    *userRepo
	courier *courierRepo
	client  *clientRepo
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

func (d *strg) Branch() storage.BranchI {
	if d.branch == nil {
		d.branch = NewBranch(d.db)
	}
	return d.branch
}

func (d *strg) Users() storage.UsersI {
	if d.user == nil {
		d.user = NewUser(d.db)
	}
	return d.user
}

func (d *strg) Couriers() storage.CouriersI {
	if d.courier == nil {
		d.courier = NewCourier(d.db)
	}
	return d.courier
}

func (d *strg) Clients() storage.ClientsI {
	if d.client == nil {
		d.client = NewClient(d.db)
	}
	return d.client
}
