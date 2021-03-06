package service

import (
	"errors"
)

var (
	// ErrorAccountBlocked when bruitforce login
	ErrorAccountBlocked = errors.New("the account is blocked for 3 min")

	// ErrorInvalidRequestData request data not valid
	ErrorInvalidRequestData = errors.New("invalid requirest data")

	// ErrorUserAccountNoActive error when the user account is not active.
	ErrorUserAccountNoActive = errors.New("user account is not active")

	// ErrorGetUser error when Email or password provided not valid.
	ErrorGetUser = errors.New("email or password not valid")

	// ErrorToGetJWTPayload when there is and error to parse jwt payload from the context
	ErrorToGetJWTPayload = errors.New("could not parse context payload")

	// ErrorCurrencyNotMatch sender wallet currency not match receiver wallet currency
	ErrorCurrencyNotMatch = errors.New("sender wallet currency not match receiver wallet currency")

	// ErrorNotenoughMoney you don't have enough money to make the transaction
	ErrorNotenoughMoney = errors.New("you don't have enough money to make the transaction")

	// ErrorCreateWallet error to create a wallet
	ErrorCreateWallet = errors.New("error to create a wallet")
)
