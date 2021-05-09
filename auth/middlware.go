package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// JwtMiddleware checks if the jwt is valid then the permissions
func (g *JWTGenerator) JwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token, err := g.getTokenAuthHeader(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			payload, err := g.VerifyToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			c.Set(PKey, payload)
			return next(c)
		}
	}
}

func (g JWTGenerator) getTokenAuthHeader(c echo.Context) (string, error) {
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		return "", ErrorAuthNotFound
	}

	parts := strings.Split(auth, " ")

	if strings.ToLower(parts[0]) != "bearer" {
		return "", ErrorAuthHeaderNotMatch
	}

	if len(parts) == 1 {
		return "", ErrorAuthTokenNotFound
	}

	if len(parts) > 2 {
		return "", ErrorAuthInvalidType
	}
	return parts[1], nil
}
