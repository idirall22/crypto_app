package model

import (
	"time"

	"github.com/microcosm-cc/bluemonday"
)

type User struct {
	ID               int32     `db:"id" json:"id"`
	Email            string    `db:"email" json:"email"`
	FirstName        string    `db:"first_name" json:"first_name"`
	LastName         string    `db:"last_name" json:"last_name"`
	IsActive         bool      `db:"is_active" json:"is_active"`
	Role             string    `db:"role" json:"role"`
	ConfirmationLink string    `db:"confirmation_link" json:"-"`
	PasswordHash     string    `db:"password_hash" json:"-"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

type RegisterUserParams struct {
	Email                   string   `json:"email" validate:"required,email"`
	FirstName               string   `json:"first_name" validate:"required,alphaunicode"`
	LastName                string   `json:"last_name" validate:"required,alphaunicode"`
	Password                string   `json:"password" validate:"required,min=8,max=50"`
	IpAddress               string   `json:"ip_address" validate:"required,ip"`
	UserAgent               string   `json:"user_agent" validate:"required"`
	XXX_PasswordHash        string   `json:"-"`
	XXX_IsActive            bool     `json:"-"`
	XXX_ConfirmationLink    string   `json:"-"`
	XXX_DefaultCurrency     []string `json:"-"`
	XXX_DefaultWalletAmount float64  `json:"-"`
	XXX_WalletAddresses     []string `json:"-"`
	XXX_DefaultRole         string   `json:"-"`
}

func (u *RegisterUserParams) Sanitize(s *bluemonday.Policy) {
	u.FirstName = s.Sanitize(u.FirstName)
	u.LastName = s.Sanitize(u.LastName)
	u.Email = s.Sanitize(u.Email)
	u.Password = s.Sanitize(u.Password)
	u.IpAddress = s.Sanitize(u.IpAddress)
	u.UserAgent = s.Sanitize(u.UserAgent)
}

type LoginUserParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

func (u *LoginUserParams) Sanitize(s *bluemonday.Policy) {
	u.Email = s.Sanitize(u.Email)
	u.Password = s.Sanitize(u.Password)
}

type ActivateAccountParams struct {
	ConfirmationLink string `json:"confirmation_link" validate:"required,uuid4"`
	XXX_IsActive     bool   `json:"-"`
}

type GetUserParams struct {
	UserID int32  `json:"id" validate:"omitempty,gt=0"`
	Email  string `json:"email" validate:"omitempty,email"`
}
