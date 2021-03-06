package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/idirall22/crypto_app/auth"
	"github.com/idirall22/crypto_app/notify/adapters/email/gmail"
	amqpeventStore "github.com/idirall22/crypto_app/notify/adapters/event/amqp"
	"github.com/idirall22/crypto_app/notify/config"
	"github.com/idirall22/crypto_app/notify/port"
	"github.com/idirall22/crypto_app/notify/service"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	// create config
	cfg := config.New()

	// create zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a logger: %v", err))
	}

	// create gmail adapter
	g := gmail.NewGmail(logger, cfg)

	// create event store adapter
	conn := connectEventStore(cfg)
	defer conn.Close()
	eventStore := amqpeventStore.NewAmqpEventStore(conn)

	// create jwt manager
	jwtGen, err := auth.NewJWTGenerator(cfg.JwtPrivatePath, cfg.JwtPublicPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error to create a JWT generator: %v", err))
	}

	// create service logic
	svc := service.NewService(logger, g, eventStore)
	ctx, fn := context.WithCancel(context.Background())
	defer fn()
	go svc.Start(ctx)

	// create port
	e := echo.New()
	echoPort := port.NewEchoPort(cfg, svc, e)
	echoPort.InitRoutes(jwtGen)

	// start server
	logger.Info(fmt.Sprintf("Server started at %s", cfg.Port))
	graceful.Run(":"+cfg.Port, 6*time.Second, e)
}

func connectEventStore(cfg *config.Config) *amqp.Connection {
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
