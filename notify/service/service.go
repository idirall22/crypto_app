package service

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	iemail "github.com/idirall22/crypto_app/notify/adapters/email"
	"github.com/idirall22/crypto_app/notify/adapters/event"
	"github.com/idirall22/crypto_app/notify/service/model"
	"go.uber.org/zap"
)

type WebsocketConnection struct {
	Conn  *websocket.Conn
	Notif chan model.Notification
}
type Service struct {
	logger      *zap.Logger
	email       iemail.IEmail
	eventStore  event.IEventStore
	Connections map[int32]WebsocketConnection
	sync.RWMutex
}

func NewService(logger *zap.Logger, email iemail.IEmail, eventStore event.IEventStore) *Service {
	return &Service{
		logger:      logger,
		email:       email,
		eventStore:  eventStore,
		Connections: make(map[int32]WebsocketConnection),
	}
}

func (s *Service) Start(ctx context.Context) error {
	go func() {
		err := s.ReceiveEmail(ctx)
		if err != nil {
			s.logger.Warn("Close receive email: " + err.Error())
		}
	}()

	go func() {
		err := s.ReceiveNotification(ctx)
		if err != nil {
			s.logger.Warn("Close receive email: " + err.Error())
		}
	}()

	<-ctx.Done()
	return nil
}
