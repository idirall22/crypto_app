package ievent

import (
	"context"

	"github.com/idirall22/crypto_app/account/service/model"
)

type IEventStore interface {
	PublishEmail(ctx context.Context, queueName string, emailChan <-chan model.EmailEvent) error
	PublishNotification(ctx context.Context, queueName string, emailChan <-chan model.NotificationEvent) error
}
