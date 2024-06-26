package constants

import (
	"errors"

	"github.com/lib/pq"
)

var ErrNotFound = errors.New("no row affected")
var ErrConflict = errors.New("error conflict")
var ErrInternalServer = errors.New("errror checking row affected")
var ErrWrongPassword = errors.New("wrong password verification")

const (
	UniqueViolationExistData    = pq.ErrorCode("23505") // 'unique_violation'
	UniqueViolationNotExistData = pq.ErrorCode("23503") // 'schema_and_data_statement_mixing_not_supported'
)
