package pgrepo_test

import (
	"context"
	"testing"

	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/stretchr/testify/require"
)

func listWalletsForTest(t *testing.T) []model.Wallet {
	user := createUserForTest(t)
	ctx := context.Background()

	params := model.ListWalletsParams{
		UserID: user.ID,
	}

	res, err := repoTest.ListWallets(ctx, params)
	require.NoError(t, err)
	require.NotZero(t, res)
	require.Equal(t, len(model.DefaultCurrencies), len(res))
	for _, w := range res {
		require.Equal(t, user.ID, w.UserID)
	}
	return res
}

func TestListWallets(t *testing.T) {
	listWalletsForTest(t)
}

func TestGetWallet(t *testing.T) {
	userWalltes := listWalletsForTest(t)

	ctx := context.Background()

	testCases := []struct {
		desc    string
		params  model.GetWalletParams
		compare func(res model.Wallet, err error)
	}{
		{
			desc: "Success",
			params: model.GetWalletParams{
				Address:    userWalltes[0].Address,
				XXX_UserID: userWalltes[0].UserID,
			},
			compare: func(res model.Wallet, err error) {
				require.NoError(t, err)
				require.Equal(t, userWalltes[0].UserID, res.UserID)
				require.Equal(t, userWalltes[0].Address, res.Address)
			},
		},
	}

	for _, c := range testCases {
		res, err := repoTest.GetWallet(ctx, c.params)
		c.compare(res, err)
	}
}
