package pgrepo

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

// transaction wrapper execute transactions.
func (p *PostgresRepo) execTX(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := fn(tx); err != nil {
		return err
	}
	return nil
}
