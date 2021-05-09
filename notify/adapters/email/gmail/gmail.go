package gmail

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"github.com/idirall22/crypto_app/notify/config"
	"github.com/idirall22/crypto_app/notify/service/model"
)

type Gmail struct {
	cfg *config.Config
}

func NewGmail(cfg *config.Config) *Gmail {
	return &Gmail{
		cfg: cfg,
	}
}
func (g *Gmail) SendRegisterUserConfirmationEmail(
	ctx context.Context, args model.RegisterUserConfirmationEmailParams) error {

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
		g.cfg.GMailEmail,
		args.Email,
		args.Subject,
		args.Body,
	)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", g.cfg.GMailSMTP, g.cfg.GMailSMTPPort),
		smtp.PlainAuth("", g.cfg.GMailEmail, g.cfg.GMailPassword, g.cfg.GMailSMTP),
		g.cfg.GMailEmail, []string{args.Email}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Print("sent")

	return nil
}
