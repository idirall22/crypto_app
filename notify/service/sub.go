package service

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/idirall22/crypto_app/auth"
	"github.com/idirall22/crypto_app/notify/service/model"
)

func (s *Service) Subscribe(ctx context.Context, conn *websocket.Conn) (<-chan model.Notification, error) {
	s.logger.Info("New user subscribe to ws")
	payload, err := auth.GetPayloadFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("-----------------", payload.UserID)

	s.Lock()
	defer s.Unlock()
	_, ok := s.Connections[payload.UserID]
	if !ok {
		s.Connections[payload.UserID] = WebsocketConnection{
			Conn:  conn,
			Notif: make(chan model.Notification, 8),
		}
	}
	return s.Connections[payload.UserID].Notif, nil
}

func (s *Service) Unsubscribe(ctx context.Context) error {
	payload, err := auth.GetPayloadFromContext(ctx)
	if err != nil {
		return err
	}
	s.Lock()
	defer s.Unlock()
	_, ok := s.Connections[payload.UserID]
	if ok {
		delete(s.Connections, payload.UserID)
	}
	return nil
}
