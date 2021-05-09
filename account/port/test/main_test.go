package port_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/idirall22/crypto_app/account/config"
	"github.com/idirall22/crypto_app/account/port"
	mockservice "github.com/idirall22/crypto_app/account/service/mock"
	"github.com/idirall22/crypto_app/auth"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var (
	portTest          *port.EchoPort
	mockService       *mockservice.IService
	engine            *echo.Echo
	userBearerTokens  auth.TokenInfos
	adminBearerTokens auth.TokenInfos
)

func TestMain(m *testing.M) {
	cfg := config.New()

	mockService = &mockservice.IService{}

	cfg.JwtPrivatePath = "../../../rsa/key.pem"
	cfg.JwtPublicPath = "../../../rsa/public.pem"

	genJWT, err := auth.NewJWTGenerator(cfg.JwtPrivatePath, cfg.JwtPublicPath)
	if err != nil {
		log.Fatal(err)
	}

	adminBearerTokens = genUserToken(genJWT, 1, "admin")
	userBearerTokens = genUserToken(genJWT, 2, "user")

	engine = echo.New()
	portTest = port.NewEchoPort(cfg, mockService, engine)

	portTest.InitRoutes(genJWT)

	os.Exit(m.Run())
}

type portTestCase struct {
	// description of the test case
	desc string
	// the url
	url string
	// methods
	method string
	// used to add queries
	request func(req *http.Request) *http.Request
	// the expected status
	status int
	// the expected mock function
	mock func()
	// body and content-type
	body func() (*strings.Reader, string)
}

func (c portTestCase) run(t *testing.T) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()

	// body
	var body *strings.Reader
	var contentType string
	if c.body != nil {
		body, contentType = c.body()
	} else {
		body = &strings.Reader{}
	}
	// body
	req, err := http.NewRequest(c.method, c.url, body)
	require.NoError(t, err, c.desc)

	// run mock
	c.mock()

	// add queries
	if c.request != nil {
		req = c.request(req)
	}

	// add application/json header
	if contentType == "" {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	} else {
		req.Header.Set(echo.HeaderContentType, contentType)
	}

	engine.ServeHTTP(rec, req)

	// check status
	require.Equal(t, c.status, rec.Code, c.desc)

	return rec
}

func genUserToken(genJWT *auth.JWTGenerator, id int32, role string) auth.TokenInfos {
	tokens, err := genJWT.CreatePairToken(id, role)
	if err != nil {
		log.Fatal(err)
	}
	return tokens
}
