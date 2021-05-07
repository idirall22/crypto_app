package service_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/auth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListTransactions(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.ListTransactionsParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Fail, admin not set the user_id",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "admin",
			}),
			params: model.ListTransactionsParams{
				Pagination: model.Pagination{
					Page:  0,
					Items: 10,
				},
				SearchTransactionBy: model.SearchTransactionBy{
					Address: func() *string {
						address := uuid.NewString()
						return &address
					}(),
				},
			},
			mock: func(ctx context.Context) {
				mockRepo.On("ListTransactions", ctx,
					mock.MatchedBy(func(input model.ListTransactionsParams) bool {
						return true
					})).Return([]model.Transaction{}, nil).Times(1)
			},
			compare: func(err error) {
				require.Error(t, err)
			},
		},
		{
			desc: "Success",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "admin",
			}),
			params: model.ListTransactionsParams{
				Pagination: model.Pagination{
					Page:  0,
					Items: 10,
				},
				SearchTransactionBy: model.SearchTransactionBy{
					Address: func() *string {
						address := uuid.NewString()
						return &address
					}(),
				},
				UserID: 1,
			},
			mock: func(ctx context.Context) {
				mockRepo.On("ListTransactions", ctx,
					mock.MatchedBy(func(input model.ListTransactionsParams) bool {
						return true
					})).Return([]model.Transaction{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			desc: "Success",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "user",
			}),
			params: model.ListTransactionsParams{
				Pagination: model.Pagination{
					Page:  0,
					Items: 10,
				},
				SearchTransactionBy: model.SearchTransactionBy{
					Address: func() *string {
						address := uuid.NewString()
						return &address
					}(),
				},
			},
			mock: func(ctx context.Context) {
				mockRepo.On("ListTransactions", ctx,
					mock.MatchedBy(func(input model.ListTransactionsParams) bool {
						return true
					})).Return([]model.Transaction{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		_, err := serviceTest.ListTransactions(c.ctx, c.params)
		c.compare(err)
	}
}
func TestSendMoney(t *testing.T) {
	ctx := context.Background()
	currency := gofakeit.RandString(model.DefaultCurrencies)
	senderAddress := gofakeit.UUID()

	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.SendMoneyParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Fail",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "user",
			}),
			params: model.SendMoneyParams{},
			mock:   func(ctx context.Context) {},
			compare: func(err error) {
				require.Error(t, err)
			},
		},
		{
			desc: "Success",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "user",
			}),
			params: model.SendMoneyParams{
				Amount:           gofakeit.Price(1, 10),
				Currency:         currency,
				SenderAddress:    senderAddress,
				RecipientAddress: gofakeit.UUID(),
			},
			mock: func(ctx context.Context) {
				mockRepo.On("GetWallet", ctx,
					mock.MatchedBy(func(input model.GetWalletParams) bool {
						return true
					})).Return(model.Wallet{
					ID:       1,
					Currency: currency,
					UserID:   1,
					Address:  senderAddress,
					Amount:   100,
				}, nil).Times(1)

				mockRepo.On("SendMoney", ctx,
					mock.MatchedBy(func(input model.SendMoneyParams) bool {
						return true
					})).Return(model.Transaction{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		_, err := serviceTest.SendMoney(c.ctx, c.params)
		c.compare(err)
	}
}
