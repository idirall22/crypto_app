package service_test

import (
	"log"
	"os"
	"testing"

	mockevent "github.com/idirall22/crypto_app/account/adapters/event/amqp/mock"
	mockrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres/mock"
	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/service"
	"github.com/idirall22/crypto_app/auth"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var (
	mockRepo    *mockrepo.IRepository
	mockEvent   *mockevent.IEventStore
	serviceTest *service.ServiceAccount
)

func TestMain(m *testing.M) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New()
	cfg.JwtPrivatePath = "../../../rsa/key.pem"
	cfg.JwtPublicPath = "../../../rsa/public.pem"

	token, err := auth.NewJWTGenerator(cfg.JwtPrivatePath, cfg.JwtPublicPath)
	if err != nil {
		log.Fatal(err)
	}

	mockRepo = &mockrepo.IRepository{}
	mockEvent = &mockevent.IEventStore{}
	serviceTest = service.NewServiceAccount(logger, mockRepo, mockEvent, token)

	os.Exit(m.Run())
}
