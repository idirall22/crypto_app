package pgrepo

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/jmoiron/sqlx"
)

var listTransactionStmt = `
SELECT * FROM transaction 
WHERE
(sender_address=COALESCE($1, sender_address) OR recipient_address=COALESCE($1, recipient_address))
AND created_at >= COALESCE($2,created_at)
AND created_at <= COALESCE($3,created_at)
ORDER BY
CASE WHEN $4 = 'desc' THEN created_at  END desc,
CASE WHEN $4 = 'asc'  THEN created_at  END asc
OFFSET $5 LIMIT $6

`

// SELECT * FROM transaction
// WHERE
// (sender_address=COALESCE($1, sender_address) OR recipient_address=COALESCE($1, recipient_address))
// AND created_at >= COALESCE($2,created_at)
// AND created_at <= COALESCE($3,created_at)
// ORDER BY created_at
// 	CASE WHEN 'desc' THEN

// OFFSET $4 LIMIT $5

var updateWalletStmt = `
UPDATE wallet SET
amount=(amount+$1)::decimal
WHERE address=$2 AND currency=$3
`

var createTransactionStmt = `
INSERT INTO transaction
(amount, commission, currency, sender_address, recipient_address)
VALUES ($1,$2,$3,$4,$5)
RETURNING id, amount, commission, currency, sender_address, recipient_address, created_at
`

func (p *PostgresRepo) ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error) {
	var trans []model.Transaction
	err := p.db.SelectContext(ctx, &trans, listTransactionStmt,
		args.Address,
		args.FromDate,
		args.ToDate,
		args.SortBy,
		args.Page,
		args.Items,
	)
	return trans, parseError(err)
}

func (p *PostgresRepo) SendMoney(ctx context.Context, args model.SendMoneyParams) (model.Transaction, error) {
	var tran model.Transaction

	err := p.execTX(ctx, func(tx *sqlx.Tx) error {
		err := tx.GetContext(ctx, &tran, createTransactionStmt,
			args.Amount,
			args.XXX_Commission,
			args.Currency,
			args.SenderAddress,
			args.RecipientAddress,
		)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, updateWalletStmt,
			args.Amount,
			args.RecipientAddress,
			args.Currency,
		)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, updateWalletStmt,
			-(args.Amount + args.XXX_Commission),
			args.SenderAddress,
			args.Currency,
		)

		return err
	})

	return tran, parseError(err)
}
