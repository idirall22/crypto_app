package service

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/idirall22/crypto_app/notify/service/model"
)

type IService interface {
	Start(ctx context.Context) error
	ReceiveEmail(ctx context.Context) error
	ReceiveNotification(ctx context.Context) error
	Subscribe(ctx context.Context, conn *websocket.Conn) (<-chan model.Notification, error)
	Unsubscribe(ctx context.Context) error
}
