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
	IsActive         string    `db:"is_active" json:"is_active"`
	ConfirmationLink string    `db:"confirmation_link" json:"-"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

type RegisterUserParams struct {
	Email                string `json:"email" validate:"required,email"`
	FirstName            string `json:"first_name" validate:"required,alphaunicode"`
	LastName             string `json:"last_name" validate:"required,alphaunicode"`
	Password             string `json:"password" validate:"required,min=8,max=50"`
	IpAddress            string `json:"ip_address" validate:"required,ip"`
	UserAgent            string `json:"user_agent" validate:"required"`
	XXX_IsActive         bool   `json:"-"`
	XXX_ConfirmationLink string `json:"-"`
}

func (u *RegisterUserParams) Sanitize(s *bluemonday.Policy) {
	u.FirstName = s.Sanitize(u.FirstName)
	u.LastName = s.Sanitize(u.LastName)
	u.Email = s.Sanitize(u.Email)
	u.Password = s.Sanitize(u.Password)
	u.IpAddress = s.Sanitize(u.IpAddress)
	u.UserAgent = s.Sanitize(u.UserAgent)
}

type ActivateUserParams struct {
	ConfirmationLink string `json:"confirmation_link" validate:"required,uuid4"`
	XXX_IsActive     bool   `json:"-"`
}

type GetUserParams struct {
	UserID int32 `json:"id"`
}
