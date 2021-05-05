package pgrepo

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
)

func (p *PostgresRepo) ListWallets(ctx context.Context, args model.ListWalletsParams) ([]model.Wallet, error) {
	var wallets []model.Wallet
	return wallets, nil
}
