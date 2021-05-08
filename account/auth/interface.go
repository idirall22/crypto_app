package auth

import "github.com/labstack/echo/v4"

// TokenGenerator interface
type TokenGenerator interface {

	// CreatePairToken creates a new token and refresh_token
	CreatePairToken(userID int32, Role string) (TokenInfos, error)

	// CreateToken creates a new token
	CreateToken(userID int32, role string) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)

	// JwtMiddleware validate jwt.
	JwtMiddleware() echo.MiddlewareFunc
}
