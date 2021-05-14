package service

import (
	"context"
	"encoding/json"

	"github.com/idirall22/crypto_app/notify/service/model"
	"github.com/streadway/amqp"
)

// ReceiveNotification logic to do when the service receive a notification
func (s *Service) ReceiveNotification(ctx context.Context) error {
	return s.eventStore.ReceiveNotification(ctx, "notification", func(d amqp.Delivery) error {
		s.logger.Info("Receive new notification...")

		var notif model.Notification
		err := json.Unmarshal(d.Body, &notif)
		if err != nil {
			return err
		}

		s.Connections[notif.UserID].Notif <- notif
		return nil
	})

}
