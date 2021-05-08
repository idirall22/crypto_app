package service_test

import (
	"log"
	"os"
	"testing"

	mockrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres/mock"
	"github.com/idirall22/crypto_app/account/auth"
	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/service"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var (
	mockRepo    *mockrepo.IRepository
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

	token, err := auth.NewJWTGenerator(cfg)
	if err != nil {
		log.Fatal(err)
	}

	mockRepo = &mockrepo.IRepository{}
	serviceTest = service.NewServiceAccount(logger, mockRepo, token)

	os.Exit(m.Run())
}
