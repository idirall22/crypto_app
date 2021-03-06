package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	ievent "github.com/idirall22/crypto_app/account/adapters/event"
	imemory "github.com/idirall22/crypto_app/account/adapters/memory"
	irepository "github.com/idirall22/crypto_app/account/adapters/repository"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/auth"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
)

type ServiceAccount struct {
	logger            *zap.Logger
	repo              irepository.IRepository
	eventStore        ievent.IEventStore
	memoryStore       imemory.IMemoryStore
	token             auth.TokenGenerator
	validator         *validator.Validate
	sanitizer         *bluemonday.Policy
	notificationsChan chan model.NotificationEvent
	emailChan         chan model.EmailEvent
}

func NewServiceAccount(
	logger *zap.Logger,
	repo irepository.IRepository,
	eventStore ievent.IEventStore,
	memoryStore imemory.IMemoryStore,
	token auth.TokenGenerator,
) *ServiceAccount {

	validator := validator.New()
	sanitizer := bluemonday.UGCPolicy()
	return &ServiceAccount{
		logger:            logger,
		repo:              repo,
		eventStore:        eventStore,
		memoryStore:       memoryStore,
		token:             token,
		validator:         validator,
		sanitizer:         sanitizer,
		notificationsChan: make(chan model.NotificationEvent, 1024),
		emailChan:         make(chan model.EmailEvent, 1024),
	}
}

func (s *ServiceAccount) Start(ctx context.Context) error {
	go func() {
		err := s.PublishEmail(ctx, "email")
		if err != nil {
			s.logger.Warn("Close publish email: " + err.Error())
		}
	}()

	go func() {
		err := s.PublishNotification(ctx, "notification")
		if err != nil {
			s.logger.Warn("Close publish notification: " + err.Error())
		}
	}()

	<-ctx.Done()
	return nil
}
