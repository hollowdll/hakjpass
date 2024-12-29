package errors

import "errors"

var (
	ErrPasswordTooShort    = errors.New("password length is too short")
	ErrEncryptedKeyInvalid = errors.New("encrypted key is invalid")
	ErrCiphertextTooShort  = errors.New("ciphertext is too short")
)
