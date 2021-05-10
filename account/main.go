package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"gopkg.in/tylerb/graceful.v1"

	amqpeventStore "github.com/idirall22/crypto_app/account/adapters/event/amqp"
	pgrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres"
	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/port"
	"github.com/idirall22/crypto_app/account/service"
	"github.com/idirall22/crypto_app/auth"
)

func main() {
	cfg := config.New()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a logger: %v", err))
	}

	// Connect to database
	dbConn := connectToDatabase(cfg)
	defer dbConn.Close()
	repo := pgrepo.NewPostgresRepo(dbConn)

	// Connect to event store
	esConn := connectToEventStore(cfg)
	defer esConn.Close()
	event := amqpeventStore.NewAmqpEventStore(logger, esConn)

	// create jwt manager
	jwtGen, err := auth.NewJWTGenerator(cfg.JwtPrivatePath, cfg.JwtPublicPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a JWT generator: %v", err))
	}

	ctx, fn := context.WithCancel(context.Background())
	defer fn()
	service := service.NewServiceAccount(logger, repo, event, jwtGen)
	go func() {
		err := service.Start(ctx)
		if err != nil {
			logger.Warn(err.Error())
		}
	}()

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

func connectToEventStore(cfg *config.Config) *amqp.Connection {
	conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			cfg.RabbitMQUser,
			cfg.RabbitMQPassword,
			cfg.RabbitMQHost,
			cfg.RabbitMQPort,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
