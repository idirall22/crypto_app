package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/idirall22/crypto_app/account/auth"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListWallets(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.ListWalletsParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Fail, user_id not set by the admin",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "admin",
			}),
			params: model.ListWalletsParams{},
			mock: func(ctx context.Context) {
				mockRepo.On("ListWallets", ctx,
					mock.MatchedBy(func(input model.ListWalletsParams) bool {
						return true
					})).Return([]model.Wallet{}, nil).Times(1)
			},
			compare: func(err error) {
				require.Error(t, err)
			},
		},
		{
			desc: "Success, user list it's wallet",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "user",
			}),
			params: model.ListWalletsParams{},
			mock: func(ctx context.Context) {
				mockRepo.On("ListWallets", ctx,
					mock.MatchedBy(func(input model.ListWalletsParams) bool {
						return true
					})).Return([]model.Wallet{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			desc: "Success, admin list users wallet",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "admin",
			}),
			params: model.ListWalletsParams{
				UserID: 1,
			},
			mock: func(ctx context.Context) {
				mockRepo.On("ListWallets", ctx,
					mock.MatchedBy(func(input model.ListWalletsParams) bool {
						return true
					})).Return([]model.Wallet{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		_, err := serviceTest.ListWallets(c.ctx, c.params)
		c.compare(err)
	}
}

func TestGetWallet(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.GetWalletParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Fail, address not set",
			ctx: context.WithValue(ctx, auth.PKey, &auth.Payload{
				UserID: 1,
				Role:   "user",
			}),
			params: model.GetWalletParams{},
			mock: func(ctx context.Context) {
				mockRepo.On("GetWallet", ctx,
					mock.MatchedBy(func(input model.GetWalletParams) bool {
						return true
					})).Return(model.Wallet{}, nil).Times(1)
			},
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
			params: model.GetWalletParams{
				Address: uuid.NewString(),
			},
			mock: func(ctx context.Context) {
				mockRepo.On("GetWallet", ctx,
					mock.MatchedBy(func(input model.GetWalletParams) bool {
						return true
					})).Return(model.Wallet{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		_, err := serviceTest.GetWallet(c.ctx, c.params)
		c.compare(err)
	}
}
