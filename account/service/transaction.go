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
		s.logger.Warn(err.Error())
		return tran, ErrorInvalidRequestData
	}

	args.XXX_Commission = 0.01
	args.XXX_UserID = payload.UserID

	tran, err = s.repo.SendMoney(ctx, args)
	if err != nil {
		return tran, ErrorInternalError
	}

	return tran, nil
}
