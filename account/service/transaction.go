package service

import (
	"context"
	"fmt"

	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/auth"
)

func (s *ServiceAccount) ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error) {
	var trans []model.Transaction

	args.Sanitize(s.sanitizer)

	args.Page, args.Items = model.Pagination(args.Page, args.Items)

	err := s.validator.Struct(args)
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

	senderWallet, err := s.repo.GetWalletByAddress(ctx, model.GetWalletParams{
		Address: args.SenderAddress,
	})
	if err != nil {
		return tran, err
	}
	recpWallet, err := s.repo.GetWalletByAddress(ctx, model.GetWalletParams{
		Address: args.RecipientAddress,
	})
	if err != nil {
		return tran, err
	}

	args.XXX_Commission = 0.01
	args.XXX_UserID = payload.UserID

	err = checkIfCanSendMoney(senderWallet, args)
	if err != nil {
		return tran, err
	}

	tran, err = s.repo.SendMoney(ctx, args)
	if err != nil {
		return tran, err
	}

	{
		s.notificationsChan <- model.NotificationEvent{
			UserID:    senderWallet.UserID,
			Type:      "transaction",
			Title:     "sent money",
			Content:   fmt.Sprintf("sent %f %s from %s to %s", tran.Amount, tran.Currency, tran.SenderAddress, tran.RecipientAddress),
			CreatedAt: tran.CreatedAt,
		}

		s.notificationsChan <- model.NotificationEvent{
			UserID:    recpWallet.UserID,
			Type:      "transaction",
			Title:     "receive money",
			Content:   fmt.Sprintf("receive %f %s from %s to %s", tran.Amount, tran.Currency, tran.RecipientAddress, tran.SenderAddress),
			CreatedAt: tran.CreatedAt,
		}
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
