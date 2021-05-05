package service

import "errors"

var (
	ErrorInvalidRequestData = errors.New("invalid requirest data")
	ErrorInternalError      = errors.New("internal error")
)
