package auth

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/idirall22/crypto_app/account/config"
)

// TokenInfos struct
type TokenInfos struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// JWTGenerator is a JSON Web Token maker
type JWTGenerator struct {
	signKey     *rsa.PrivateKey
	verifiedKey *rsa.PublicKey
}

// NewJWTGenerator creates a new JWTGenerator
func NewJWTGenerator(cfg *config.Config) (*JWTGenerator, error) {
	var JWTGenerator = &JWTGenerator{}

	signBytes, err := ioutil.ReadFile(cfg.JwtPrivatePath)
	if err != nil {
		return JWTGenerator, err
	}

	verifyBytes, err := ioutil.ReadFile(cfg.JwtPublicPath)
	if err != nil {
		return JWTGenerator, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return JWTGenerator, err
	}
	verifiedKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return JWTGenerator, err
	}
	JWTGenerator.signKey = signKey
	JWTGenerator.verifiedKey = verifiedKey

	return JWTGenerator, nil
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
	ti.AccessToken = token
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

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	token, err := jwtToken.SignedString(g.signKey)
	return token, err
}

// VerifyToken checks if the token is valid or not
func (g *JWTGenerator) VerifyToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, ErrInvalidToken
		}
		return g.verifiedKey, nil
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
