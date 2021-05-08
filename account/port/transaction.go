package port

import (
	"net/http"

	"github.com/idirall22/crypto_app/account/service"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/labstack/echo/v4"
)

func (p *EchoPort) ListTransactions(c echo.Context) error {
	var params model.ListTransactionsParams
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, service.ErrorInvalidRequestData.Error())
	}

	transactions, err := p.service.ListTransactions(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(parseError(err))
	}

	return c.JSON(http.StatusOK, transactions)
}

func (p *EchoPort) SendMoney(c echo.Context) error {
	var params model.SendMoneyParams
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, service.ErrorInvalidRequestData.Error())
	}

	transaction, err := p.service.SendMoney(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(parseError(err))
	}

	return c.JSON(http.StatusCreated, transaction)
}
