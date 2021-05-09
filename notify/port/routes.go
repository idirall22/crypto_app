package port

import (
	"github.com/idirall22/crypto_app/notify/auth"
)

func (p *EchoPort) InitRoutes(jwtGen auth.TokenGenerator) {
	p.engin.GET("/healthy", p.Healthy)
	p.engin.GET("/ws", p.Notification)
}
