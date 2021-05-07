package service

import (
	"context"

	"github.com/idirall22/crypto_app/account/auth"
	"github.com/idirall22/crypto_app/account/service/model"
)

type IService interface {
	RegisterUser(ctx context.Context, args model.RegisterUserParams) error
	LoginUser(ctx context.Context, args model.LoginUserParams) (auth.TokenInfos, error)
	ActivateAccount(ctx context.Context, args model.ActivateAccountParams) error
	GetUser(ctx context.Context, args model.GetUserParams) (model.User, error)

	ListWallets(ctx context.Context, args model.ListWalletsParams) ([]model.Wallet, error)
	GetWallet(ctx context.Context, args model.GetWalletParams) (model.Wallet, error)

	ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error)
	SendMoney(ctx context.Context, args model.SendMoneyParams) (model.Transaction, error)
}
