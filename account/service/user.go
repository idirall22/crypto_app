package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/idirall22/crypto_app/account/service/model"
)

func (s *ServiceAccount) RegisterUser(ctx context.Context, args model.RegisterUserParams) error {

	args.Sanitize(s.sanitizer)
	err := s.validator.Struct(args)
	if err != nil {
		return ErrorInvalidRequestData
	}

	args.XXX_ConfirmationLink = uuid.New().String()
	args.XXX_IsActive = false

	newUser, err := s.repo.RegisterUser(ctx, args)
	if err != nil {
		return ErrorInternalError
	}

	// send email
	go func(user model.User) {
		err := s.email.EmailRegisterUserConfirmation(user.Email, user.FirstName, user.ConfirmationLink)
		if err != nil {
			s.logger.Warn(fmt.Sprintf("Error to send RegisterUserConfirmationEmail:%v", err))
		}
	}(newUser)

	return nil
}

func (s *ServiceAccount) LoginUser(ctx context.Context, args model.RegisterUserParams) error {
	// add redis filtring to check how many times user tried to login
	return nil
}

func (s *ServiceAccount) ActivateUser(ctx context.Context, args model.ActivateUserParams) error {
	err := s.validator.Struct(args)
	if err != nil {
		return ErrorInvalidRequestData
	}

	args.XXX_IsActive = true
	_, err = s.repo.ActivateUser(ctx, args)

	return err
}

func (s *ServiceAccount) GetUser(ctx context.Context, args model.GetUserParams) (model.User, error) {
	return s.repo.GetUser(ctx, args)
}
