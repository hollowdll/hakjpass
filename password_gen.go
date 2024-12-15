package hakjpass

import (
	"crypto/rand"
	"math/big"

	"github.com/hollowdll/hakjpass/errors"
)

const (
	PasswordLowerChars       = "abcdefghijklmnopqrstuvwxyz"
	PasswordUpperChars       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	PasswordNumberChars      = "0123456789"
	PasswordSpecialChars     = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	MinPasswordLength    int = 4
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

// GenerateRandomSecurePassword generates a random secure password using the specified length and opts.
// The generated password includes at least one character of each character set in the opts.
// This function uses the crypto package in the Go standard library for cryptographically secure random values.
func GenerateRandomSecurePassword(length int, opts *PasswordOptions) (string, error) {
	if length < MinPasswordLength {
		return "", errors.ErrPasswordTooShort
	}

	password := []byte{}
	charsetAll := opts.LowerChars + opts.UpperChars + opts.NumberChars + opts.SpecialChars
	charsets := []string{opts.LowerChars, opts.UpperChars, opts.NumberChars, opts.SpecialChars}

	// include at least one character from each character set
	for _, charset := range charsets {
		index, err := randomIndex(len(charset))
		if err != nil {
			return "", err
		}
		password = append(password, charset[index])
	}

	// fill the remaining characters
	for i := MinPasswordLength; i < length; i++ {
		index, err := randomIndex(len(charsetAll))
		if err != nil {
			return "", err
		}
		password = append(password, charsetAll[index])
	}

	// randomly shuffle the password to avoid predictable patterns
	if err := randomlyShuffleChars(password); err != nil {
		return "", err
	}

	return string(password), nil
}

// randomIndex generates a random index between 0 and max. max is exclusive.
func randomIndex(max int) (int, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(num.Int64()), nil
}

func randomlyShuffleChars(slice []byte) error {
	for i := range slice {
		j, err := randomIndex(len(slice))
		if err != nil {
			return err
		}
		slice[i], slice[j] = slice[j], slice[i]
	}
	return nil
}
