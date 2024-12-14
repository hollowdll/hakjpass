package errors

import "errors"

var (
	ErrPasswordTooShort = errors.New("password length is too short")
)
