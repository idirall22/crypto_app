package port

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/idirall22/crypto_app/auth"
	"github.com/idirall22/crypto_app/notify/config"
	"github.com/idirall22/crypto_app/notify/service"
	"github.com/labstack/echo/v4"
)

type EchoPort struct {
	service  service.IService
	engin    *echo.Echo
	cfg      *config.Config
	upgrader *websocket.Upgrader
}

// NewEchoPort create EchoPort
func NewEchoPort(cfg *config.Config, service service.IService, e *echo.Echo) *EchoPort {
	return &EchoPort{
		cfg:     cfg,
		service: service,
		engin:   e,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (p *EchoPort) Healthy(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}

// Notification Upgrade a connection to websocket and push notifications
func (p *EchoPort) Notification(c echo.Context) error {

	conn, err := p.upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		return echo.NewHTTPError(http.StatusUpgradeRequired, err.Error())
	}
	defer conn.Close()

	// Subscribe user
	notifChan, err := p.service.Subscribe(auth.Context(c), conn)
	if err != nil {
		return echo.NewHTTPError(http.StatusUpgradeRequired, "error to subscribe")
	}
	defer p.service.Unsubscribe(auth.Context(c))

	for {
		err = conn.WriteJSON(<-notifChan)
		if err != nil {
			return err
		}
		fmt.Println("Sent to websocket")
	}
}
