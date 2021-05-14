package gmail

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/idirall22/crypto_app/notify/config"
	"github.com/idirall22/crypto_app/notify/service/model"
	"go.uber.org/zap"
)

type Gmail struct {
	logger *zap.Logger
	cfg    *config.Config
}

// NewGmail create Gmail adapter
func NewGmail(logger *zap.Logger, cfg *config.Config) *Gmail {
	return &Gmail{
		logger: logger,
		cfg:    cfg,
	}
}

// SendRegisterUserConfirmationEmail send register user email using Gmail.
func (g *Gmail) SendRegisterUserConfirmationEmail(
	ctx context.Context, args model.RegisterUserConfirmationEmailParams) error {

	if g.cfg.GMailEmail == "" || g.cfg.GMailPassword == "" {
		g.logger.Warn("Email and email password no set")
		return nil
	}

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
		g.cfg.GMailEmail,
		args.Email,
		args.Subject,
		args.Body,
	)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", g.cfg.GMailSMTP, g.cfg.GMailSMTPPort),
		smtp.PlainAuth("", g.cfg.GMailEmail, g.cfg.GMailPassword, g.cfg.GMailSMTP),
		g.cfg.GMailEmail, []string{args.Email}, []byte(msg),
	)

	if err != nil {
		g.logger.Warn("smtp error: " + err.Error())
		return err
	}

	g.logger.Info("email sent")

	return nil
}
