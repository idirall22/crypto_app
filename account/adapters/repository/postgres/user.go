package pgrepo

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/jmoiron/sqlx"
)

var createUserStmt = `
INSERT INTO users (first_name, last_name, email, is_active, confirmation_link, password_hash, role)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING
id, first_name, last_name, email, is_active, confirmation_link, role, created_at;
`

var createUserMetadataStmt = `
INSERT INTO user_metadata (user_id, ip_address, user_agent) VALUES ($1,$2,$3)
`

var createWalletStmt = `
INSERT INTO wallet (currency, user_id, address, amount)
VALUES($1,$2,$3,$4)
`

var getUserStmt = `
SELECT
	id, first_name, last_name, email, is_active, confirmation_link, role, password_hash, created_at
FROM users
	WHERE id=COALESCE($1, id) OR email=COALESCE($2, email)
LIMIT 1
`
var activateAccountStmt = `
UPDATE users SET 
	is_active=$1,
	confirmation_link=''
WHERE confirmation_link=$2
RETURNING id, first_name, last_name, email, is_active, confirmation_link, created_at;
`

func (p *PostgresRepo) RegisterUser(ctx context.Context, args model.RegisterUserParams) (model.User, error) {
	var user model.User

	err := p.execTX(ctx, func(tx *sqlx.Tx) error {
		// create user
		err := tx.GetContext(ctx, &user, createUserStmt,
			args.FirstName,
			args.LastName,
			args.Email,
			args.XXX_IsActive,
			args.XXX_ConfirmationLink,
			args.XXX_PasswordHash,
			args.XXX_DefaultRole,
		)
		if err != nil {
			return err
		}

		// create user metadata
		_, err = tx.ExecContext(ctx, createUserMetadataStmt, user.ID, args.IpAddress, args.UserAgent)
		if err != nil {
			return err
		}

		// create wallets
		for i := 0; i < len(args.XXX_WalletAddresses); i++ {
			_, err = tx.ExecContext(ctx, createWalletStmt,
				args.XXX_DefaultCurrency[i],
				user.ID,
				args.XXX_WalletAddresses[i],
				args.XXX_DefaultWalletAmount,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return user, err
}

func (p *PostgresRepo) ActivateAccount(ctx context.Context, args model.ActivateAccountParams) (model.User, error) {
	var user model.User
	err := p.db.GetContext(ctx, &user, activateAccountStmt, args.XXX_IsActive, args.ConfirmationLink)
	return user, err
}

func (p *PostgresRepo) GetUser(ctx context.Context, args model.GetUserParams) (model.User, error) {
	var user model.User
	err := p.db.GetContext(ctx, &user, getUserStmt, args.UserID, args.Email)
	return user, err
}
