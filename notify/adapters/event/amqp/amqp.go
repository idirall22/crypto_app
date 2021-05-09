package amqpeventStore

import (
	"context"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type AmqpEventStore struct {
	logger *zap.Logger
	conn   *amqp.Connection
}

func NewAmqpEventStore(conn *amqp.Connection) *AmqpEventStore {
	return &AmqpEventStore{
		conn: conn,
	}
}

func (a *AmqpEventStore) ReceiveEmail(ctx context.Context,
	queueName string, fn func(d amqp.Delivery) error) error {

	ch, err := a.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	delivery, err := ch.Consume(queueName, "", true, false, false, false, amqp.Table{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case result := <-delivery:
			err := fn(result)
			if err != nil {
				a.logger.Warn(err.Error())
			}
		}
	}
}

func (a *AmqpEventStore) ReceiveNotification(ctx context.Context,
	queueName string, fn func(d amqp.Delivery) error) error {

	ch, err := a.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	delivery, err := ch.Consume(queueName, "", true, false, false, false, amqp.Table{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case result := <-delivery:
			err := fn(result)
			if err != nil {
				a.logger.Warn(err.Error())
			}
		}
	}
}
