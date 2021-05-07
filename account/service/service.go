package service

import (
	"github.com/go-playground/validator/v10"
	irepository "github.com/idirall22/crypto_app/account/adapters/repository"
	"github.com/idirall22/crypto_app/account/auth"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
)

type ServiceAccount struct {
	logger    *zap.Logger
	repo      irepository.IRepository
	token     auth.TokenGenerator
	validator *validator.Validate
	sanitizer *bluemonday.Policy
}

func NewServiceAccount(
	logger *zap.Logger,
	repo irepository.IRepository,
	token auth.TokenGenerator,
) *ServiceAccount {

	validator := validator.New()
	sanitizer := bluemonday.UGCPolicy()
	return &ServiceAccount{
		logger:    logger,
		repo:      repo,
		token:     token,
		validator: validator,
		sanitizer: sanitizer,
	}
}
