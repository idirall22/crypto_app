package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/idirall22/crypto_app/auth"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestContextPayload(t *testing.T) {
	m, err := auth.NewJWTGenerator("../rsa/key.pem", "../rsa/public.pem")
	require.NoError(t, err)

	token, err := m.CreateToken(1, "user")
	require.NoError(t, err)

	fmt.Println(token)
	engine := echo.New()
	engine.GET("/", func(c echo.Context) error {
		payload, err := auth.GetPayloadFromContext(auth.Context(c))
		require.NoError(t, err)
		require.NotNil(t, payload)
		return nil
	}, m.JwtMiddleware())

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Add("Authorization", "bearer "+token)
	engine.ServeHTTP(rec, req)

	fmt.Println(rec.Code)
}
