package port_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/auth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	testCases := []portTestCase{
		{
			desc:   "Success",
			url:    "/register",
			method: http.MethodPost,
			request: func(req *http.Request) *http.Request {
				return req
			},
			status: http.StatusCreated,
			mock: func() {
				mockService.On("RegisterUser", context.Background(),
					mock.MatchedBy(func(input model.RegisterUserParams) bool {
						return true
					})).Return("", nil).Times(1)
			},
			body: func() (*strings.Reader, string) {
				data, err := json.Marshal(model.RegisterUserParams{
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
					Email:     gofakeit.Email(),
					Password:  gofakeit.Password(true, true, true, false, false, 12),
				})
				require.NoError(t, err)
				return strings.NewReader(string(data)), ""
			},
		},
	}

	for _, c := range testCases {
		c.run(t)
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []portTestCase{
		{
			desc:   "Success",
			url:    "/login",
			method: http.MethodPost,
			request: func(req *http.Request) *http.Request {
				return req
			},
			status: http.StatusOK,
			mock: func() {
				mockService.On("LoginUser", context.Background(),
					mock.MatchedBy(func(input model.LoginUserParams) bool {
						return true
					})).Return(auth.TokenInfos{}, nil).Times(1)
			},
			body: func() (*strings.Reader, string) {
				data, err := json.Marshal(model.LoginUserParams{
					Email:    gofakeit.Email(),
					Password: gofakeit.Password(true, true, true, false, false, 12),
				})
				require.NoError(t, err)
				return strings.NewReader(string(data)), ""
			},
		},
	}

	for _, c := range testCases {
		c.run(t)
	}
}
func TestActivateAccount(t *testing.T) {
	testCases := []portTestCase{
		{
			desc:   "Success",
			url:    "/activate_account",
			method: http.MethodPost,
			request: func(req *http.Request) *http.Request {
				return req
			},
			status: http.StatusOK,
			mock: func() {
				mockService.On("ActivateAccount", context.Background(),
					mock.MatchedBy(func(input model.ActivateAccountParams) bool {
						return true
					})).Return(nil).Times(1)
			},
			body: func() (*strings.Reader, string) {
				data, err := json.Marshal(model.ActivateAccountParams{
					ConfirmationLink: gofakeit.UUID(),
				})
				require.NoError(t, err)
				return strings.NewReader(string(data)), ""
			},
		},
	}

	for _, c := range testCases {
		c.run(t)
	}
}
func TestGetUser(t *testing.T) {
	testCases := []portTestCase{
		{
			desc:   "Success",
			url:    "/profile/1",
			method: http.MethodGet,
			request: func(req *http.Request) *http.Request {
				req.Header.Add("Authorization", "bearer "+userBearerTokens.AccessToken)
				return req
			},
			status: http.StatusOK,
			mock: func() {
				mockService.On("GetUser", context.Background(),
					mock.MatchedBy(func(input model.GetUserParams) bool {
						return true
					})).Return(model.User{}, nil).Times(1)
			},
			body: func() (*strings.Reader, string) {
				return &strings.Reader{}, ""
			},
		},
	}

	for _, c := range testCases {
		c.run(t)
	}
}
