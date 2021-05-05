package pgrepo

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
)

func (p *PostgresRepo) ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error) {
	var trans []model.Transaction
	return trans, nil
}

func (p *PostgresRepo) SendMoney(ctx context.Context, args model.SendMoneyParams) (model.Transaction, error) {
	var tran model.Transaction
	return tran, nil
}
