package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/account/utils"
	"github.com/idirall22/crypto_app/auth"
)

func (s *ServiceAccount) RegisterUser(ctx context.Context, args model.RegisterUserParams) error {

	args.Sanitize(s.sanitizer)
	err := s.validator.Struct(args)
	if err != nil {
		return ErrorInvalidRequestData
	}

	passwordHash, err := utils.HashPassword(args.Password)
	if err != nil {
		return ErrorInvalidRequestData
	}

	args.XXX_ConfirmationLink = uuid.New().String()
	args.XXX_IsActive = false
	args.XXX_DefaultCurrency = model.DefaultCurrencies
	args.XXX_DefaultRole = "user"
	args.XXX_DefaultWalletAmount = 100
	args.XXX_PasswordHash = passwordHash

	for range model.DefaultCurrencies {
		_, pub, err := utils.GenerateWallet()
		if err != nil {
			return ErrorCreateWallet
		}
		args.XXX_WalletAddresses = append(args.XXX_WalletAddresses, pub)
	}

	_, err = s.repo.RegisterUser(ctx, args)
	if err != nil {
		s.logger.Warn(err.Error())
		return err
	}
	return err
}

func (s *ServiceAccount) LoginUser(ctx context.Context, args model.LoginUserParams) (auth.TokenInfos, error) {

	var tokens auth.TokenInfos
	args.Sanitize(s.sanitizer)

	err := s.validator.Struct(args)
	if err != nil {
		return tokens, ErrorInvalidRequestData
	}

	user, err := s.repo.GetUser(ctx, model.GetUserParams{Email: args.Email})
	if err != nil {
		return tokens, err
	}

	if !user.IsActive {
		return tokens, ErrorUserAccountNoActive
	}

	err = utils.CheckPassword(args.Password, user.PasswordHash)
	if err != nil {
		return tokens, ErrorGetUser
	}

	// publish email to event store
	s.emailChan <- model.EmailEvent{
		Email:     user.Email,
		FirstName: user.FirstName,
		Subject:   "Sign up",
		Body:      fmt.Sprintf("Thank you for register\nPlease confirm you email using this confirmation token: %s", user.ConfirmationLink),
	}

	return s.token.CreatePairToken(user.ID, user.Role)
}

func (s *ServiceAccount) ActivateAccount(ctx context.Context, args model.ActivateAccountParams) error {
	err := s.validator.Struct(args)
	if err != nil {
		return ErrorInvalidRequestData
	}

	args.XXX_IsActive = true
	_, err = s.repo.ActivateAccount(ctx, args)

	return err
}

func (s *ServiceAccount) GetUser(ctx context.Context, args model.GetUserParams) (model.User, error) {
	err := s.validator.Struct(args)
	if err != nil {
		return model.User{}, ErrorInvalidRequestData
	}
	return s.repo.GetUser(ctx, args)
}
