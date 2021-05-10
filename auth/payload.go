package auth

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	PKey string = "payload"

	// DefaultTokenDuration the duration for a token is 15min
	DefaultTokenDuration = time.Minute * 15

	// DefaultRefreshTokenDuration the duration for a refresh_token is 7 days
	DefaultRefreshTokenDuration = time.Hour * 24 * 7
)

var (
	// ErrInvalidToken when the token is not valid
	ErrInvalidToken = errors.New("token is invalid")

	// ErrExpiredToken when the token has expired
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	UserID    int32     `json:"uid"`
	Role      string    `json:"uro"`
	IssuedAt  time.Time `json:"ist"`
	ExpiredAt time.Time `json:"ext"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(userID int32, role string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		UserID:    userID,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// GetPayloadFromContext get payload from context
func GetPayloadFromContext(ctx echo.Context) (*Payload, error) {
	var (
		payload *Payload
		ok      bool
	)

	payload, ok = ctx.Get(PKey).(*Payload)
	if !ok {
		return payload, ErrorToGetJWTPayload
	}

	return payload, nil
}
