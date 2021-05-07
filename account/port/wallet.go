package port

import (
	"net/http"

	"github.com/idirall22/crypto_app/account/service"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/labstack/echo/v4"
)

func (p *EchoPort) ListWallets(c echo.Context) error {
	var params model.ListWalletsParams
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, service.ErrorInvalidRequestData.Error())
	}

	wallets, err := p.service.ListWallets(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(parseError(err))
	}

	return c.JSON(http.StatusOK, wallets)
}
