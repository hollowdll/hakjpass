package hakjpass

import (
	"fmt"

	"github.com/hollowdll/hakjpass/errors"
)

const (
	PasswordLowerChars   = "abcdefghijklmnopqrstuvwxyz"
	PasswordUpperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	PasswordNumberChars  = "0123456789"
	PasswordSpecialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

type PasswordOptions struct {
	LowerChars   string
	UpperChars   string
	NumberChars  string
	SpecialChars string
}

func DefaultPasswordOptions() *PasswordOptions {
	return &PasswordOptions{
		LowerChars:   PasswordLowerChars,
		UpperChars:   PasswordUpperChars,
		NumberChars:  PasswordNumberChars,
		SpecialChars: PasswordSpecialChars,
	}
}

func GenerateRandomSecurePassword(length int, opts *PasswordOptions) (string, error) {
	if length < 4 {
		return "", errors.ErrPasswordTooShort
	}

	// TODO logic

	return "", nil
}
