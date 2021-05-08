package model

type Wallet struct {
	ID       int32   `db:"id" json:"id"`
	Currency string  `db:"currency" json:"currency"`
	UserID   int32   `db:"user_id" json:"user_id"`
	Address  string  `db:"address" json:"address"`
	Amount   float64 `db:"amount" json:"amount"`
}

type ListWalletsParams struct {
	UserID int32 `param:"user_id" validate:"required,gt=0"`
}

type GetWalletParams struct {
	Address    string `json:"address" validate:"required,uuid4"`
	XXX_UserID int32  `json:"-"`
}
