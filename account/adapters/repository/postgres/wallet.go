package pgrepo

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
)

var listUserWalletsStmt = `
SELECT
	id, address, user_id, amount, currency
FROM wallet
WHERE user_id=$1
`
var getUserWalletStmt = `
SELECT * FROM wallet WHERE address=$1 AND user_id=$2 LIMIT 1
`

func (p *PostgresRepo) ListWallets(ctx context.Context, args model.ListWalletsParams) ([]model.Wallet, error) {
	var wallets []model.Wallet
	err := p.db.SelectContext(ctx, &wallets, listUserWalletsStmt, args.UserID)
	return wallets, err
}

func (p *PostgresRepo) GetWallet(ctx context.Context, args model.GetWalletParams) (model.Wallet, error) {
	var wallet model.Wallet
	err := p.db.GetContext(ctx, &wallet, getUserWalletStmt, args.Address, args.XXX_UserID)
	return wallet, err
}
