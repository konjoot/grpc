package sessions

import (
	"errors"
)

var (
	ErrTokenNotFound = errors.New("token not found")
	ErrEmptyRequest  = errors.New("empty request")
	ErrNotFound      = errors.New("user not found")
	ErrWrongCreds    = errors.New("bad users creds")
)
