package auth

import "errors"

var (
	ErrorAuthNotFound       = errors.New("authorization not found")
	ErrorAuthHeaderNotMatch = errors.New("authorization header must start with 'Bearer'")
	ErrorAuthTokenNotFound  = errors.New("token not found")
	ErrorAuthInvalidType    = errors.New("authorization header must be bearer token")
	ErrorAuthForbidden      = errors.New("forbidden")
	ErrorToGetJWTPayload    = errors.New("could not parse context payload")
)
