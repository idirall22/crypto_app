package port_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/stretchr/testify/mock"
)

func TestListWallets(t *testing.T) {
	testCases := []portTestCase{
		{
			desc:   "Success, user list its wallets",
			url:    "/wallets",
			method: http.MethodGet,
			request: func(req *http.Request) *http.Request {
				req.Header.Add("Authorization", "bearer "+userBearerTokens.AccessToken)
				return req
			},
			status: http.StatusOK,
			mock: func() {
				mockService.On("ListWallets", context.Background(),
					mock.MatchedBy(func(input model.ListWalletsParams) bool {
						return true
					})).Return([]model.Wallet{}, nil).Times(1)
			},
			body: func() (*strings.Reader, string) {
				return &strings.Reader{}, ""
			},
		},

		{
			desc:   "Success, admin list user wallets",
			url:    "/wallets",
			method: http.MethodGet,
			request: func(req *http.Request) *http.Request {
				q := req.URL.Query()
				q.Add("user_id", "1")
				req.URL.RawQuery = q.Encode()

				req.Header.Add("Authorization", "bearer "+adminBearerTokens.AccessToken)
				return req
			},
			status: http.StatusOK,
			mock: func() {
				mockService.On("ListWallets", context.Background(),
					mock.MatchedBy(func(input model.ListWalletsParams) bool {
						return true
					})).Return([]model.Wallet{}, nil).Times(1)
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
