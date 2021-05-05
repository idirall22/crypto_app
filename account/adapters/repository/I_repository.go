package irepository

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
)

type IRepository interface {
	RegisterUser(ctx context.Context, args model.RegisterUserParams) (model.User, error)
	ActivateUser(ctx context.Context, args model.ActivateUserParams) (model.User, error)
	GetUser(ctx context.Context, args model.GetUserParams) (model.User, error)
	ListWallets(ctx context.Context, args model.ListWalletsParams) ([]model.Wallet, error)
	ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error)
	SendMoney(ctx context.Context, args model.SendMoneyParams) (model.Transaction, error)
}
