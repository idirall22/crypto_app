package service

import (
	"context"

	"github.com/idirall22/crypto_app/account/auth"
	"github.com/idirall22/crypto_app/account/service/model"
)

func (s *ServiceAccount) ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error) {
	var trans []model.Transaction

	payload, err := auth.GetPayloadFromContext(ctx)
	if err != nil {
		return trans, err
	}

	if payload.Role != "admin" {
		args.UserID = payload.UserID
	}

	err = s.validator.Struct(args)
	if err != nil {
		return trans, ErrorInvalidRequestData
	}

	return s.repo.ListTransactions(ctx, args)
}

func (s *ServiceAccount) SendMoney(ctx context.Context, args model.SendMoneyParams) (model.Transaction, error) {

	var tran model.Transaction

	payload, err := auth.GetPayloadFromContext(ctx)
	if err != nil {
		return tran, err
	}

	err = s.validator.Struct(args)
	if err != nil {
		return tran, ErrorInvalidRequestData
	}

	wallet, err := s.GetWallet(ctx, model.GetWalletParams{
		Address: args.SenderAddress,
	})
	if err != nil {
		return tran, err
	}

	args.XXX_Commission = 0.01
	args.XXX_UserID = payload.UserID

	err = checkIfCanSendMoney(wallet, args)
	if err != nil {
		return tran, err
	}

	tran, err = s.repo.SendMoney(ctx, args)
	if err != nil {
		return tran, ErrorInternalError
	}

	return tran, nil
}

func checkIfCanSendMoney(w model.Wallet, t model.SendMoneyParams) error {
	if w.Currency != t.Currency {
		return ErrorCurrencyNotMatch
	}
	if w.Amount < t.Amount+t.XXX_Commission {
		return ErrorNotenoughMoney
	}
	return nil
}
