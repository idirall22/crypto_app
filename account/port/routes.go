package port

import (
	"github.com/idirall22/crypto_app/account/auth"
)

func (p *EchoPort) InitRoutes(jwtGen auth.TokenGenerator) {
	p.engin.POST("/register", p.RegisterUser)
	p.engin.POST("/login", p.LoginUser)
	p.engin.POST("/activate_account", p.ActivateAccount)
	p.engin.GET("/profile/:user_id", p.GetUser, jwtGen.JwtMiddleware())

	p.engin.POST("/wallets", p.SendMoney, jwtGen.JwtMiddleware())
	p.engin.GET("/transactions", p.ListTransactions, jwtGen.JwtMiddleware())
	p.engin.GET("/wallets", p.ListWallets, jwtGen.JwtMiddleware())
}
