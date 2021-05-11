package port_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/account/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListTransactions(t *testing.T) {
	testCases := []portTestCase{
		{
			desc:   "Success",
			url:    "/transactions",
			method: http.MethodGet,
			request: func(req *http.Request) *http.Request {
				q := req.URL.Query()
				q.Add("page", "1")
				q.Add("items", "25")
				q.Add("from_date", "2020-01-01")
				q.Add("to_date", "2020-01-01")
				q.Add("address", func() string {
					_, pub, err := utils.GenerateWallet()
					require.NoError(t, err)
					return pub
				}())
				q.Add("sort_by", "desc")
				req.URL.RawQuery = q.Encode()
				req.Header.Add("Authorization", "bearer "+userBearerTokens.AccessToken)
				return req
			},
			status: http.StatusOK,
			mock: func() {
				mockService.On("ListTransactions", mock.MatchedBy(func(input context.Context) bool {
					return true
				}),
					mock.MatchedBy(func(input model.ListTransactionsParams) bool {
						return true
					})).Return([]model.Transaction{}, nil).Times(1)
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

func TestSendMoney(t *testing.T) {
	testCases := []portTestCase{
		{
			desc:   "Success",
			url:    "/send_money",
			method: http.MethodPost,
			request: func(req *http.Request) *http.Request {
				req.Header.Add("Authorization", "bearer "+userBearerTokens.AccessToken)
				return req
			},
			status: http.StatusCreated,
			mock: func() {
				mockService.On("SendMoney", mock.MatchedBy(func(input context.Context) bool {
					return true
				}),
					mock.MatchedBy(func(input model.SendMoneyParams) bool {
						return true
					})).Return(model.Transaction{}, nil).Times(1)
			},
			body: func() (*strings.Reader, string) {
				data, err := json.Marshal(model.SendMoneyParams{
					Amount:           gofakeit.Price(1, 10),
					Currency:         gofakeit.RandString(model.DefaultCurrencies),
					SenderAddress:    gofakeit.UUID(),
					RecipientAddress: gofakeit.UUID(),
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
