package service

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/auth"
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

	err = s.validator.Struct(args)
	if err != nil {
		return wallets, ErrorInvalidRequestData
	}
	return s.repo.ListWallets(ctx, args)
}

func (s *ServiceAccount) GetWallet(ctx context.Context, args model.GetWalletParams) (model.Wallet, error) {
	var wallets model.Wallet

	payload, err := auth.GetPayloadFromContext(ctx)
	if err != nil {
		return wallets, err
	}

	err = s.validator.Struct(args)
	if err != nil {
		return wallets, ErrorInvalidRequestData
	}

	if payload.Role != "admin" {
		args.XXX_UserID = payload.UserID
	}

	return s.repo.GetWallet(ctx, args)
}
