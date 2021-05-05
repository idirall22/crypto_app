package auth

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	minSecretKeySize = 32
	issuer           = "crypto_app.com"
)

// TokenInfos struct
type TokenInfos struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// JWTGenerator is a JSON Web Token maker
type JWTGenerator struct {
	secretKey string
}

// NewJWTGenerator creates a new JWTGenerator
func NewJWTGenerator(secretKey string) (*JWTGenerator, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTGenerator{secretKey}, nil
}

// CreatePairToken creates a new token and refresh_token
func (g *JWTGenerator) CreatePairToken(userID int32, Role string) (TokenInfos, error) {
	ti := TokenInfos{}
	token, err := g.createToken(userID, Role)
	if err != nil {
		return ti, err
	}
	refreshToken, err := g.createToken(userID, Role)
	if err != nil {
		return ti, err
	}
	ti.RefreshToken = refreshToken
	ti.Token = token
	return ti, nil
}

// CreateToken creates a new token
func (g *JWTGenerator) CreateToken(userID int32, role string) (string, error) {
	return g.createToken(userID, role)
}

func (g *JWTGenerator) createToken(userID int32, role string) (string, error) {
	payload, err := NewPayload(userID, role, DefaultTokenDuration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(g.secretKey))
	return token, err
}

// VerifyToken checks if the token is valid or not
func (g *JWTGenerator) VerifyToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(g.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
