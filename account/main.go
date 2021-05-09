package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gopkg.in/tylerb/graceful.v1"

	pgrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres"
	"github.com/idirall22/crypto_app/account/auth"
	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/port"
	"github.com/idirall22/crypto_app/account/service"
)

func main() {
	cfg := config.New()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a logger: %v", err))
	}

	db := connectToDatabase(cfg)
	defer db.Close()

	repo := pgrepo.NewPostgresRepo(db)

	jwtGen, err := auth.NewJWTGenerator(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a JWT generator: %v", err))
	}

	service := service.NewServiceAccount(logger, repo, jwtGen)

	e := echo.New()
	echoPort := port.NewEchoPort(cfg, service, e)
	echoPort.InitRoutes(jwtGen)

	logger.Info(fmt.Sprintf("Server started at %s", cfg.Port))

	graceful.Run(":"+cfg.Port, 5*time.Second, e)
	// err = graceful.ListenAndServe(e.Server, 5*time.Second)
	// if err != nil {
	// 	logger.Info(err.Error())
	// }
}

func connectToDatabase(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Connect(cfg.RepositoryConfig())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
