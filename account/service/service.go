package service

import (
	"github.com/go-playground/validator/v10"
	iemail "github.com/idirall22/crypto_app/account/adapters/email"
	irepository "github.com/idirall22/crypto_app/account/adapters/repository"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
)

type ServiceAccount struct {
	logger    *zap.Logger
	repo      irepository.IRepository
	email     iemail.IEmail
	validator *validator.Validate
	sanitizer *bluemonday.Policy
}

func NewServiceAccount(logger *zap.Logger, repo irepository.IRepository, email iemail.IEmail) *ServiceAccount {
	validator := validator.New()
	sanitizer := bluemonday.UGCPolicy()
	return &ServiceAccount{
		logger:    logger,
		repo:      repo,
		email:     email,
		validator: validator,
		sanitizer: sanitizer,
	}
}
