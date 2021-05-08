package auth

import (
	"errors"
	"net/http"
	"time"
)

// DefaultCookieExpire time for a cookie to expire 1 day.
const DefaultCookieExpire = time.Hour * 24

// ErrorSetCookie error when the args passes to the set cookie are not pair.
var ErrorSetCookie = errors.New("error args should be a pair of [key, value] list")

// NewCookie create new cookie http cookie
func NewCookie(name, value, domain, path string, httpOnly, secure bool) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		Expires:  time.Now().Add(DefaultCookieExpire),
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: http.SameSiteStrictMode,
	}
}
