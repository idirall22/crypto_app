package amqpeventStore

import (
	"context"
	"encoding/json"

	"github.com/idirall22/crypto_app/account/service/model"
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

func (a *AmqpEventStore) PublishEmail(
	ctx context.Context,
	queueName string, emailChan <-chan model.EmailEvent,
) error {

	ch, err := a.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case email := <-emailChan:
			a.logger.Warn("Publish new email")
			body, err := json.Marshal(email)
			if err != nil {
				return err
			}
			err = ch.Publish("", queue.Name, false, false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			if err != nil {
				return err
			}
		}
	}
}

func (a *AmqpEventStore) PublishNotification(ctx context.Context,
	queueName string, notifyChan <-chan model.NotificationEvent) error {

	ch, err := a.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case notif := <-notifyChan:
			a.logger.Warn("Publish new notification")
			body, err := json.Marshal(notif)
			if err != nil {
				return err
			}
			err = ch.Publish("", queue.Name, false, false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			if err != nil {
				return err
			}
		}
	}
}
