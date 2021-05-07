package pgrepo

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

const (

	// NotNullViolationError Postgres error
	NotNullViolationError = "23502"

	// ForeignKeyViolation Postgres error
	ForeignKeyViolation = "23503"

	// UniqueViolation Postgres error
	UniqueViolation = "23505"
)

var (
	// ErrorInternalError Error server
	ErrorInternalError = errors.New("internal error")

	// ErrorAlreadyExists error when the record already exists.
	ErrorAlreadyExists = errors.New("user already exists")

	// ErrorNotExists error when there is no entity.
	ErrorNotExists = errors.New("the entity does not exist")
)

// ParseRepoErrors parse database errors and return
func parseError(err error) error {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return ErrorNotExists
	}

	pqErr, ok := err.(*pq.Error)
	if ok {
		switch pqErr.Code.Name() {
		case NotNullViolationError:
			return ErrorInternalError

		case ForeignKeyViolation:
			return ErrorInternalError

		case UniqueViolation:
			return ErrorAlreadyExists
		}
	}
	return ErrorInternalError
}
