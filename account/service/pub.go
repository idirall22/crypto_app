package service

import (
	"context"
)

func (s *ServiceAccount) PublishEmail(ctx context.Context, queueName string) error {
	return s.eventStore.PublishEmail(ctx, queueName, s.emailChan)
}

func (s *ServiceAccount) PublishNotification(ctx context.Context, queueName string) error {
	return s.eventStore.PublishNotification(ctx, queueName, s.notificationsChan)
}
