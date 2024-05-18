package constants

import "errors"

var ErrNotFound = errors.New("no row affected")
var ErrConflict = errors.New("error conflict")
var ErrInternalServer = errors.New("errror checking row affected")
