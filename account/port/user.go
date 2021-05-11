package port

import (
	"net/http"

	"github.com/idirall22/crypto_app/account/service"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/labstack/echo/v4"
)

func (p *EchoPort) RegisterUser(c echo.Context) error {
	var params model.RegisterUserParams
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, service.ErrorInvalidRequestData.Error())
	}

	params.XXX_IpAddress = c.Request().RemoteAddr
	params.XXX_UserAgent = c.Request().UserAgent()

	link, err := p.service.RegisterUser(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(parseError(err))
	}

	return c.JSON(http.StatusCreated, map[string]string{"confirmation_link": link})
}

func (p *EchoPort) LoginUser(c echo.Context) error {
	var params model.LoginUserParams
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, service.ErrorInvalidRequestData.Error())
	}

	tokens, err := p.service.LoginUser(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(parseError(err))
	}

	return c.JSON(http.StatusOK, tokens)
}

func (p *EchoPort) ActivateAccount(c echo.Context) error {
	var params model.ActivateAccountParams
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, service.ErrorInvalidRequestData.Error())
	}

	err = p.service.ActivateAccount(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(parseError(err))
	}
	return c.JSON(http.StatusOK, "success")
}

func (p *EchoPort) GetUser(c echo.Context) error {
	var params model.GetUserParams
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, service.ErrorInvalidRequestData.Error())
	}

	user, err := p.service.GetUser(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(parseError(err))
	}

	return c.JSON(http.StatusOK, user)
}
