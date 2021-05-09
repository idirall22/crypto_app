package iemail

import (
	"context"

	"github.com/idirall22/crypto_app/notify/service/model"
)

type IEmail interface {
	SendRegisterUserConfirmationEmail(ctx context.Context, args model.RegisterUserConfirmationEmailParams) error
}
