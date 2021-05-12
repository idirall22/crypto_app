package irepository

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
)

type IRepository interface {
	RegisterUser(ctx context.Context, args model.RegisterUserParams) (model.User, error)
	ActivateAccount(ctx context.Context, args model.ActivateAccountParams) (model.User, error)
	GetUser(ctx context.Context, args model.GetUserParams) (model.User, error)
	ListWallets(ctx context.Context, args model.ListWalletsParams) ([]model.Wallet, error)
	GetWallet(ctx context.Context, args model.GetWalletParams) (model.Wallet, error)
	GetWalletByAddress(ctx context.Context, args model.GetWalletParams) (model.Wallet, error)
	ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error)
	SendMoney(ctx context.Context, args model.SendMoneyParams) (model.Transaction, error)
}
