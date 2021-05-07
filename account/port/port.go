package port

import "github.com/idirall22/crypto_app/account/service"

type EchoPort struct {
	service service.IService
}

func NewEchoPort(service service.IService) *EchoPort {
	return &EchoPort{
		service: service,
	}
}
