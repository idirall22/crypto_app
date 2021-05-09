package event

import (
	"context"

	"github.com/streadway/amqp"
)

type IEventStore interface {
	ReceiveEmail(ctx context.Context, queueName string, fn func(d amqp.Delivery) error) error
	ReceiveNotification(ctx context.Context, queueName string, fn func(d amqp.Delivery) error) error
}
