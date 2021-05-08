package model

import (
	"time"

	"github.com/microcosm-cc/bluemonday"
)

type Transaction struct {
	ID               int32     `db:"id" json:"id"`
	Amount           float64   `db:"amount" json:"amount"`
	Comission        float64   `db:"commission" json:"commission"`
	Currency         string    `db:"currency" json:"currency"`
	SenderAddress    string    `db:"sender_address" json:"sender_address"`
	RecipientAddress string    `db:"recipient_address" json:"recipient_address"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

type ListTransactionsParams struct {
	Address  string    `json:"address" query:"address" validate:"required"`
	FromDate *Datetime `json:"from_date" query:"from_date" validate:"omitempty"`
	ToDate   *Datetime `json:"to_date" query:"to_date" validate:"omitempty"`
	SortBy   string    `json:"sort_by" query:"sort_by" validate:"omitempty,oneof=desc asc"`
	Page     int32     `json:"page" query:"page"`
	Items    int32     `json:"items" query:"items"`
}

func (p *ListTransactionsParams) Sanitize(s *bluemonday.Policy) {
	p.Address = s.Sanitize(p.Address)
}

func (p *ListTransactionsParams) NormalizeInput() {
	p.Page, p.Items = Pagination(p.Page, p.Items)
	if p.SortBy == "" {
		p.SortBy = "desc"
	}
}

type SendMoneyParams struct {
	Amount           float64 `json:"amount" validate:"required,gt=0"`
	Currency         string  `json:"currency" validate:"required,gt=0"`
	SenderAddress    string  `json:"sender_address" validate:"required,uuid4"`
	RecipientAddress string  `json:"recipient_address" validate:"required,uuid4"`
	XXX_Commission   float64 `json:"-"`
	XXX_UserID       int32   `json:"-"`
}
