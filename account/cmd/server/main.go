package main

import (
	"fmt"
	"log"

	pgrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres"
	"github.com/idirall22/crypto_app/account/auth"
	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/port"
	"github.com/idirall22/crypto_app/account/service"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a logger: %v", err))
	}

	db := connectToDatabase(cfg)
	repo := pgrepo.NewPostgresRepo(db)

	jwtGen := auth.NewJWTGenerator(cfg)

	service := service.NewServiceAccount(logger, repo, jwtGen)

	echoPort := port.NewEchoPort(service)

	e := echo.New()
}

func connectToDatabase(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Connect(cfg.RepositoryConfig())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
