package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"gopkg.in/tylerb/graceful.v1"

	amqpeventStore "github.com/idirall22/crypto_app/account/adapters/event/amqp"
	redismem "github.com/idirall22/crypto_app/account/adapters/memory/redis"
	pgrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres"
	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/port"
	"github.com/idirall22/crypto_app/account/service"
	"github.com/idirall22/crypto_app/auth"
)

func main() {
	// create config
	cfg := config.New()

	// create zap logger
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

	// Connect with memory store
	rConn := connectToMemoryStore(cfg)
	defer rConn.Close()
	memory := redismem.NewRedisMemory(rConn)

	// create jwt manager
	jwtGen, err := auth.NewJWTGenerator(cfg.JwtPrivatePath, cfg.JwtPublicPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a JWT generator: %v", err))
	}

	ctx, fn := context.WithCancel(context.Background())
	defer fn()

	// create new service logic
	service := service.NewServiceAccount(logger, repo, event, memory, jwtGen)
	go func() {
		err := service.Start(ctx)
		if err != nil {
			logger.Warn(err.Error())
		}
	}()

	// create new port
	e := echo.New()
	echoPort := port.NewEchoPort(cfg, service, e)
	echoPort.InitRoutes(jwtGen)

	// run server
	logger.Info(fmt.Sprintf("Server started at %s", cfg.Port))
	graceful.Run(":"+cfg.Port, 5*time.Second, e)
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

func connectToMemoryStore(cfg *config.Config) *redis.Client {
	conn := redis.NewClient(
		&redis.Options{
			Addr:       fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
			MaxRetries: 3,
		},
	)
	return conn
}
