package model

import (
	"time"
)

type Transaction struct {
	ID               int32     `db:"id" json:"id"`
	Amount           float64   `db:"amount" json:"amount"`
	Comission        float64   `db:"commission" json:"commission"`
	Currency         string    `db:"currency_id" json:"currency_id"`
	SenderAddress    string    `db:"sender_address" json:"sender_address"`
	RecipientAddress string    `db:"recipient_address" json:"recipient_address"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

type ListTransactionsParams struct {
	Pagination          Pagination          `json:"pagination"`
	SearchTransactionBy SearchTransactionBy `json:"search_transaction_by"`
	UserID              int32               `json:"user_id" validate:"required,gt=0"`
}

type SearchTransactionBy struct {
	Address  string    `json:"address" validate:"omitempty,uuid4"`
	FromDate time.Time `json:"from_date" validate:"omitempty"`
	ToDate   time.Time `json:"to_date" validate:"omitempty"`
}

type SendMoneyParams struct {
	Amount           float64 `json:"amount" validate:"required,gt=0"`
	Currency_id      int32   `json:"currency_id" validate:"required,gt=0"`
	SenderAddress    string  `json:"sender_address" validate:"required,uuid4"`
	RecipientAddress string  `json:"recipient_address" validate:"required,uuid4"`
	XXX_Commission   float64 `json:"commission" validate:"required,gt=0"`
	XXX_UserID       int32   `json:"-"`
}
