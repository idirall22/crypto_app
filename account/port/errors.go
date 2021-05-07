package port

import (
	"net/http"

	pgrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres"
	"github.com/idirall22/crypto_app/account/service"
)

// ParseError return the error and the status code
func parseError(err error) (int, error) {

	switch err {
	case service.ErrorInvalidRequestData:
		// 400
		return http.StatusBadRequest, service.ErrorInvalidRequestData
	case pgrepo.ErrorNotExists:
		// 404
		return http.StatusNotFound, pgrepo.ErrorNotExists
	case pgrepo.ErrorAlreadyExists:
		// 409
		return http.StatusConflict, pgrepo.ErrorAlreadyExists
	default:
		// 500
		return http.StatusInternalServerError, pgrepo.ErrorInternalError
	}
}
