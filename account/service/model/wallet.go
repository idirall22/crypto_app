package model

type Wallet struct {
	ID       int32  `db:"id" json:"id"`
	Currency string `db:"currency" json:"currency"`
	UserID   string `db:"user_id" json:"user_id"`
	Address  string `db:"address" json:"address"`
	Amount   string `db:"amount" json:"amount"`
}

type ListWalletsParams struct {
	UserID int32 `json:"user_id"`
}
