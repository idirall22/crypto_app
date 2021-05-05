package pgrepo

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/jmoiron/sqlx"
)

func (p *PostgresRepo) RegisterUser(ctx context.Context, args model.RegisterUserParams) (model.User, error) {
	var user model.User

	err := p.execTX(ctx, func(tx *sqlx.Tx) error {
		tx.
		return nil
	})

	return user, err
}

func (p *PostgresRepo) ActivateUser(ctx context.Context, args model.ActivateUserParams) (model.User, error) {
	var user model.User
	return user, nil
}

func (p *PostgresRepo) GetUser(ctx context.Context, args model.GetUserParams) (model.User, error) {
	var user model.User
	return user, nil
}
