package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/idirall22/crypto_app/account/service/model"
	"github.com/idirall22/crypto_app/account/utils"
	"github.com/idirall22/crypto_app/auth"
)

const (
	MaxLoginAttempts = 3
	LoginBlockTime   = time.Second * 180
)

func (s *ServiceAccount) RegisterUser(ctx context.Context, args model.RegisterUserParams) (string, error) {

	args.Sanitize(s.sanitizer)
	err := s.validator.Struct(args)
	if err != nil {
		return "", ErrorInvalidRequestData
	}

	passwordHash, err := utils.HashPassword(args.Password)
	if err != nil {
		return "", ErrorInvalidRequestData
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
			return "", ErrorCreateWallet
		}
		args.XXX_WalletAddresses = append(args.XXX_WalletAddresses, pub)
	}

	user, err := s.repo.RegisterUser(ctx, args)
	if err != nil {
		s.logger.Warn(err.Error())
		return "", err
	}

	s.emailChan <- model.EmailEvent{
		Email:     user.Email,
		FirstName: user.FirstName,
		Subject:   "register",
		Body:      fmt.Sprintf("Confirmation link: %s", user.ConfirmationLink),
	}

	return user.ConfirmationLink, err
}

func (s *ServiceAccount) LoginUser(ctx context.Context, args model.LoginUserParams) (auth.TokenInfos, error) {

	var tokens auth.TokenInfos
	args.Sanitize(s.sanitizer)

	err := s.validator.Struct(args)
	if err != nil {
		return tokens, ErrorInvalidRequestData
	}

	fmt.Println("-----------------------------------1")
	// check if user is not blocked and can login
	err = s.CheckIfUserCanLogin(args.Email)
	if err != nil {
		fmt.Println("2----------------------------", err)
		if err == ErrorAccountBlocked {
			return tokens, ErrorAccountBlocked
		}
		return tokens, err
	}

	user, err := s.repo.GetUser(ctx, model.GetUserParams{Email: args.Email})
	if err != nil {
		fmt.Println("3----------------------------", err)
		return tokens, err
	}

	if !user.IsActive {
		fmt.Println("4----------------------------")
		return tokens, ErrorUserAccountNoActive
	}

	err = utils.CheckPassword(args.Password, user.PasswordHash)
	if err != nil {
		fmt.Println("5----------------------------", err)
		return tokens, ErrorGetUser
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

func (s *ServiceAccount) CheckIfUserCanLogin(email string) error {
	b := strings.Builder{}
	b.WriteString("login-")
	b.WriteString(email)
	key := b.String()

	attempts, err := s.memoryStore.GetLoginAttemps(key)
	if err != nil {
		return err
	}

	if attempts >= MaxLoginAttempts {
		return ErrorAccountBlocked
	}

	attempts += 1
	err = s.memoryStore.SetLoginAttemps(key, attempts, LoginBlockTime)
	if err != nil {
		return err
	}

	return nil
}
