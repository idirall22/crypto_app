package service_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/account/utils"
	"github.com/idirall22/crypto_app/auth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListTransactions(t *testing.T) {

	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.ListTransactionsParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Success",
			ctx:  context.Background(),
			params: func() model.ListTransactionsParams {
				_, pubAdds, err := utils.GenerateWallet()
				require.NoError(t, err)
				return model.ListTransactionsParams{
					Page:    0,
					Items:   10,
					Address: pubAdds,
					SortBy:  "desc",
				}
			}(),
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
				mockRepo.On("GetWalletByAddress", ctx,
					mock.MatchedBy(func(input model.GetWalletParams) bool {
						return true
					})).Return(model.Wallet{
					ID:       1,
					Currency: currency,
					UserID:   1,
					Address:  senderAddress,
					Amount:   100,
				}, nil).Times(2)
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
