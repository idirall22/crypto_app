// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/idirall22/crypto_app/auth"

	mock "github.com/stretchr/testify/mock"

	model "github.com/idirall22/crypto_app/account/service/model"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// ActivateAccount provides a mock function with given fields: ctx, args
func (_m *IService) ActivateAccount(ctx context.Context, args model.ActivateAccountParams) error {
	ret := _m.Called(ctx, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.ActivateAccountParams) error); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, args
func (_m *IService) GetUser(ctx context.Context, args model.GetUserParams) (model.User, error) {
	ret := _m.Called(ctx, args)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, model.GetUserParams) model.User); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.GetUserParams) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWallet provides a mock function with given fields: ctx, args
func (_m *IService) GetWallet(ctx context.Context, args model.GetWalletParams) (model.Wallet, error) {
	ret := _m.Called(ctx, args)

	var r0 model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, model.GetWalletParams) model.Wallet); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(model.Wallet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.GetWalletParams) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTransactions provides a mock function with given fields: ctx, args
func (_m *IService) ListTransactions(ctx context.Context, args model.ListTransactionsParams) ([]model.Transaction, error) {
	ret := _m.Called(ctx, args)

	var r0 []model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, model.ListTransactionsParams) []model.Transaction); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.ListTransactionsParams) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWallets provides a mock function with given fields: ctx, args
func (_m *IService) ListWallets(ctx context.Context, args model.ListWalletsParams) ([]model.Wallet, error) {
	ret := _m.Called(ctx, args)

	var r0 []model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, model.ListWalletsParams) []model.Wallet); ok {
		r0 = rf(ctx, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.ListWalletsParams) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: ctx, args
func (_m *IService) LoginUser(ctx context.Context, args model.LoginUserParams) (auth.TokenInfos, error) {
	ret := _m.Called(ctx, args)

	var r0 auth.TokenInfos
	if rf, ok := ret.Get(0).(func(context.Context, model.LoginUserParams) auth.TokenInfos); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(auth.TokenInfos)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.LoginUserParams) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, args
func (_m *IService) RegisterUser(ctx context.Context, args model.RegisterUserParams) (string, error) {
	ret := _m.Called(ctx, args)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, model.RegisterUserParams) string); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.RegisterUserParams) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMoney provides a mock function with given fields: ctx, args
func (_m *IService) SendMoney(ctx context.Context, args model.SendMoneyParams) (model.Transaction, error) {
	ret := _m.Called(ctx, args)

	var r0 model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, model.SendMoneyParams) model.Transaction); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(model.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SendMoneyParams) error); ok {
		r1 = rf(ctx, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
