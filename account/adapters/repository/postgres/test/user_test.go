package pgrepo_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/stretchr/testify/require"
)

var defaultCurrencies = []string{"bitcoin", "ether"}

func createUserForTest(t *testing.T) model.User {
	ctx := context.Background()
	params := model.RegisterUserParams{
		Email:                   gofakeit.Email(),
		FirstName:               gofakeit.FirstName(),
		LastName:                gofakeit.LastName(),
		Password:                gofakeit.Password(true, true, true, false, false, 12),
		IpAddress:               gofakeit.IPv4Address(),
		UserAgent:               gofakeit.UserAgent(),
		XXX_PasswordHash:        gofakeit.Password(true, true, true, false, false, 12),
		XXX_IsActive:            false,
		XXX_ConfirmationLink:    uuid.NewString(),
		XXX_DefaultCurrency:     defaultCurrencies,
		XXX_DefaultWalletAmount: 100,
		XXX_WalletAddresses:     []string{uuid.New().String(), uuid.New().String()},
		XXX_DefaultRole:         "user",
	}
	res, err := repoTest.RegisterUser(ctx, params)
	require.NoError(t, err)
	require.NotZero(t, res)
	require.False(t, res.IsActive)
	return res
}
func TestRegisterUser(t *testing.T) {
	createUserForTest(t)
}
func TestActivateAccount(t *testing.T) {
	user := createUserForTest(t)
	ctx := context.Background()

	testCases := []struct {
		desc    string
		params  model.ActivateAccountParams
		compare func(res model.User, err error)
	}{
		{
			desc: "fail, invalid confirmation link",
			params: model.ActivateAccountParams{
				ConfirmationLink: uuid.NewString(),
				XXX_IsActive:     true,
			},
			compare: func(res model.User, err error) {},
		},
		{
			desc: "success",
			params: model.ActivateAccountParams{
				ConfirmationLink: user.ConfirmationLink,
				XXX_IsActive:     true,
			},
			compare: func(res model.User, err error) {
				require.NoError(t, err)
				require.True(t, res.IsActive)
				require.Zero(t, res.ConfirmationLink)
			},
		},
	}

	for _, c := range testCases {
		res, err := repoTest.ActivateAccount(ctx, c.params)
		c.compare(res, err)
	}
}
func TestGetUser(t *testing.T) {
	user := createUserForTest(t)
	ctx := context.Background()

	testCases := []struct {
		desc    string
		params  model.GetUserParams
		compare func(res model.User, err error)
	}{
		{
			desc: "Success, get user by id",
			params: model.GetUserParams{
				UserID: user.ID,
			},
			compare: func(res model.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.ID, res.ID)
				require.Equal(t, user.FirstName, res.FirstName)
				require.Equal(t, user.LastName, res.LastName)
				require.Equal(t, user.Email, res.Email)
				require.Equal(t, user.Role, res.Role)
				require.Equal(t, user.IsActive, res.IsActive)
			},
		},
		{
			desc: "Success, get user by email",
			params: model.GetUserParams{
				Email: user.Email,
			},
			compare: func(res model.User, err error) {
				require.NoError(t, err)
				require.Equal(t, user.ID, res.ID)
				require.Equal(t, user.FirstName, res.FirstName)
				require.Equal(t, user.LastName, res.LastName)
				require.Equal(t, user.Email, res.Email)
				require.Equal(t, user.Role, res.Role)
				require.Equal(t, user.IsActive, res.IsActive)
			},
		},
	}
	for _, c := range testCases {
		res, err := repoTest.GetUser(ctx, c.params)
		c.compare(res, err)
	}
}
