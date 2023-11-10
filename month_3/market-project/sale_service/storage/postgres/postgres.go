package postgres

import (
	"context"
	"fmt"
	"sale_service/config"
	"sale_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db                    *pgxpool.Pool
	sale                  *saleRepo
	sale_product          *saleProductRepo
	staff_transaction     *staffTransactionRepo
	branch_pr_transaction *branchPrTranRepo
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

func (d *strg) Sale() storage.SaleI {
	if d.sale == nil {
		d.sale = NewSaleRepo(d.db)
	}
	return d.sale
}

func (d *strg) SaleProduct() storage.SaleProductI {
	if d.sale_product == nil {
		d.sale_product = NewSaleProductRepo(d.db)
	}
	return d.sale_product
}

func (d *strg) StaffTransaction() storage.StaffTransactionI {
	if d.staff_transaction == nil {
		d.staff_transaction = NewStaffTransactionRepo(d.db)
	}
	return d.staff_transaction
}

func (d *strg) BranchProductTransactions() storage.BranchPrTransactionI {
	if d.branch_pr_transaction == nil {
		d.branch_pr_transaction = NewBranchPrTranRepo(d.db)
	}
	return d.branch_pr_transaction
}
