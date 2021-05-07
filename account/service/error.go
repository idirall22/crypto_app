package service

import "errors"

var (
	ErrorInvalidRequestData = errors.New("invalid requirest data")
	ErrorInternalError      = errors.New("internal error")
	ErrorGetUser            = errors.New("email or password not valid")

	ErrorCurrencyNotMatch = errors.New("sender wallet currency not match receiver wallet currency")
	ErrorNotenoughMoney   = errors.New("you don't have enough money to make the transaction")
)
