package port

import (
	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/service"
	"github.com/labstack/echo/v4"
)

type EchoPort struct {
	service service.IService
	engin   *echo.Echo
	cfg     *config.Config
}

func NewEchoPort(cfg *config.Config, service service.IService, e *echo.Echo) *EchoPort {
	return &EchoPort{
		cfg:     cfg,
		service: service,
		engin:   e,
	}
}
