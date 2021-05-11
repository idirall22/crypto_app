package port

import (
	"net/http"

	pgrepo "github.com/idirall22/crypto_app/account/adapters/repository/postgres"
	"github.com/idirall22/crypto_app/account/service"
)

// ParseError return the error and the status code
func parseError(err error) (int, string) {
	switch err {

	case service.ErrorInvalidRequestData:
		return http.StatusBadRequest, service.ErrorInvalidRequestData.Error()

	case service.ErrorGetUser:
		return http.StatusBadRequest, service.ErrorGetUser.Error()

	case service.ErrorUserAccountNoActive:
		return http.StatusForbidden, service.ErrorUserAccountNoActive.Error()

	case service.ErrorCurrencyNotMatch:
		return http.StatusForbidden, service.ErrorCurrencyNotMatch.Error()

	case service.ErrorNotenoughMoney:
		return http.StatusForbidden, service.ErrorNotenoughMoney.Error()

	case service.ErrorCreateWallet:
		return http.StatusForbidden, service.ErrorCreateWallet.Error()

	case pgrepo.ErrorNotExists:
		return http.StatusNotFound, pgrepo.ErrorNotExists.Error()

	case pgrepo.ErrorAlreadyExists:
		return http.StatusConflict, pgrepo.ErrorAlreadyExists.Error()

	case service.ErrorAccountBlocked:
		return http.StatusLocked, service.ErrorAccountBlocked.Error()

	default:
		return http.StatusInternalServerError, pgrepo.ErrorInternalError.Error()
	}
}
