package service_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/idirall22/crypto_app/account/service"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/account/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.RegisterUserParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Fail, invalid email",
			ctx:  context.Background(),
			params: model.RegisterUserParams{
				Email:     gofakeit.FirstName(),
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Password:  gofakeit.Password(true, true, true, false, false, 12),
				IpAddress: gofakeit.IPv4Address(),
				UserAgent: gofakeit.UserAgent(),
			},
			mock: func(ctx context.Context) {},
			compare: func(err error) {
				require.Equal(t, service.ErrorInvalidRequestData.Error(), err.Error())
			},
		},
		{
			desc: "Success",
			ctx:  context.Background(),
			params: model.RegisterUserParams{
				Email:     gofakeit.Email(),
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Password:  gofakeit.Password(true, true, true, false, false, 12),
				IpAddress: gofakeit.IPv4Address(),
				UserAgent: gofakeit.UserAgent(),
			},
			mock: func(ctx context.Context) {
				mockRepo.On("RegisterUser", ctx,
					mock.MatchedBy(func(input model.RegisterUserParams) bool {
						return true
					})).Return(model.User{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		err := serviceTest.RegisterUser(c.ctx, c.params)
		c.compare(err)
	}
}
func TestLoginUser(t *testing.T) {
	password := gofakeit.Password(true, true, true, false, false, 12)
	passwordHash, err := utils.HashPassword(password)
	require.NoError(t, err)

	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.LoginUserParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Success",
			ctx:  context.Background(),
			params: model.LoginUserParams{
				Email:    gofakeit.Email(),
				Password: password,
			},
			mock: func(ctx context.Context) {
				mockRepo.On("GetUser", ctx,
					mock.MatchedBy(func(input model.GetUserParams) bool {
						return true
					})).Return(model.User{ID: 1, Role: "user", PasswordHash: passwordHash}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		_, err := serviceTest.LoginUser(c.ctx, c.params)
		c.compare(err)
	}
}

func TestActivateAccount(t *testing.T) {
	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.ActivateAccountParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Fail, invalid confirmation_link",
			ctx:  context.Background(),
			params: model.ActivateAccountParams{
				ConfirmationLink: "",
			},
			mock: func(ctx context.Context) {},
			compare: func(err error) {
				require.Equal(t, service.ErrorInvalidRequestData.Error(), err.Error())
			},
		},
		{
			desc: "Success",
			ctx:  context.Background(),
			params: model.ActivateAccountParams{
				ConfirmationLink: uuid.NewString(),
			},
			mock: func(ctx context.Context) {
				mockRepo.On("ActivateAccount", ctx,
					mock.MatchedBy(func(input model.ActivateAccountParams) bool {
						return true
					})).Return(model.User{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		err := serviceTest.ActivateAccount(c.ctx, c.params)
		c.compare(err)
	}

}

func TestGetUser(t *testing.T) {
	testCases := []struct {
		desc    string
		ctx     context.Context
		params  model.GetUserParams
		mock    func(ctx context.Context)
		compare func(err error)
	}{
		{
			desc: "Success, get user by id",
			ctx:  context.Background(),
			params: model.GetUserParams{
				UserID: 1,
			},
			mock: func(ctx context.Context) {
				mockRepo.On("GetUser", ctx,
					mock.MatchedBy(func(input model.GetUserParams) bool {
						return true
					})).Return(model.User{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			desc: "Success, get user by email",
			ctx:  context.Background(),
			params: model.GetUserParams{
				Email: gofakeit.Email(),
			},
			mock: func(ctx context.Context) {
				mockRepo.On("GetUser", ctx,
					mock.MatchedBy(func(input model.GetUserParams) bool {
						return true
					})).Return(model.User{}, nil).Times(1)
			},
			compare: func(err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, c := range testCases {
		c.mock(c.ctx)
		_, err := serviceTest.GetUser(c.ctx, c.params)
		c.compare(err)
	}
}
