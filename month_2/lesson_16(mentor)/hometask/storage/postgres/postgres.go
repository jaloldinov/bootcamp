package postgres

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db       *pgxpool.Pool
	branches *branchRepo
	// staffes     *staffRepo
	// sales       *saleRepo
	// transaction *transactionRepo
	// staffTarifs *staffTarifRepo
}

func NewConnection(cfg *config.Config) (storage.StorageI, error) {
	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnections

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (b *store) Branch() storage.BranchesI {
	if b.branches == nil {
		b.branches = NewBranchRepo(b.db)
	}
	return b.branches
}
func (s *store) Close() {
	s.db.Close()
}

// func (s *store) Staff() storage.StaffesI {
// 	return s.staffes
// }

// func (s *store) Sales() storage.SalesI {
// 	return s.sales
// }

// func (s *store) Transaction() storage.TransactionI {
// 	return s.transaction
// }

// func (s *store) StaffTarif() storage.StaffTarifsI {
// 	return s.staffTarifs
// }
