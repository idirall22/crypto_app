package pgrepo_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/stretchr/testify/require"
)

func sendMoneyForTest(t *testing.T) (model.Wallet, model.Wallet) {
	currency := gofakeit.RandomString(model.DefaultCurrencies)
	senderWallet := getWalletAddressByCurrency(listWalletsForTest(t), currency)
	recipientWallet := getWalletAddressByCurrency(listWalletsForTest(t), currency)

	ctx := context.Background()

	params := model.SendMoneyParams{
		Amount:           gofakeit.Price(1, 10),
		Currency:         currency,
		SenderAddress:    senderWallet.Address,
		RecipientAddress: recipientWallet.Address,
		XXX_Commission:   float64(gofakeit.Price(0.01, 1)),
		XXX_UserID:       senderWallet.UserID,
	}
	res, err := repoTest.SendMoney(ctx, params)
	require.NoError(t, err)
	require.NotZero(t, res)

	// validate wallets amount
	{
		res1, err := repoTest.GetWallet(ctx, model.GetWalletParams{
			Address:    senderWallet.Address,
			XXX_UserID: senderWallet.UserID,
		})
		require.NoError(t, err)
		require.NotZero(t, res1)
		require.Equal(t, float64(senderWallet.Amount-params.Amount-params.XXX_Commission), res1.Amount)

		res2, err := repoTest.GetWallet(ctx, model.GetWalletParams{
			Address:    recipientWallet.Address,
			XXX_UserID: recipientWallet.UserID,
		})
		require.NoError(t, err)
		require.NotZero(t, res2)
		require.Equal(t, (senderWallet.Amount + params.Amount), res2.Amount)
	}

	return senderWallet, recipientWallet
}

func TestSendMoney(t *testing.T) {
	sendMoneyForTest(t)
}

func TestListTransactions(t *testing.T) {
	ctx := context.Background()
	sender, _ := sendMoneyForTest(t)

	params := model.ListTransactionsParams{
		Address: sender.Address,
		Page:    0,
		Items:   10,
		SortBy:  "desc",
	}

	res, err := repoTest.ListTransactions(ctx, params)
	require.NoError(t, err)
	require.NotZero(t, res)
}

func getWalletAddressByCurrency(walltes []model.Wallet, currency string) model.Wallet {
	for _, w := range walltes {
		if w.Currency == currency {
			return w
		}
	}
	return model.Wallet{}
}
