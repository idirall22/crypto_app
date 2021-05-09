package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/idirall22/crypto_app/notify/service/model"
	"github.com/streadway/amqp"
)

func (s *Service) ReceiveEmail(ctx context.Context) error {
	return s.eventStore.ReceiveEmail(ctx, "email", func(d amqp.Delivery) error {
		s.logger.Info("Receive new email...")

		var email model.RegisterUserConfirmationEmailParams
		err := json.Unmarshal(d.Body, &email)
		if err != nil {
			return err
		}

		ctxt, f := context.WithTimeout(context.Background(), time.Second*10)
		defer f()

		go func() {
			err := s.email.SendRegisterUserConfirmationEmail(ctxt, email)
			if err != nil {
				s.logger.Warn("Error to send email: " + err.Error())
			}
		}()

		return nil
	})
}
