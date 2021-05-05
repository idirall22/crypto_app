package service

import (
	"context"

	"github.com/idirall22/crypto_app/account/auth"
	"github.com/idirall22/crypto_app/account/service/model"
)

func (s *ServiceAccount) ListWallets(ctx context.Context, args model.ListWalletsParams) ([]model.Wallet, error) {
	var wallets []model.Wallet

	payload, err := auth.GetPayloadFromContext(ctx)
	if err != nil {
		return wallets, err
	}

	if payload.Role != "admin" {
		args.UserID = payload.UserID
	}

	return s.repo.ListWallets(ctx, args)
}
