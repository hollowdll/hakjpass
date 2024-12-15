package hakjpass

import (
	"testing"

	"github.com/hollowdll/hakjpass/errors"
	"github.com/hollowdll/hakjpass/internal/common"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomSecurePassword(t *testing.T) {
	length := 4
	opts := &PasswordOptions{
		LowerChars:   "abc",
		UpperChars:   "ABC",
		NumberChars:  "123",
		SpecialChars: "!?-",
	}
	password, err := GenerateRandomSecurePassword(length, opts)
	require.NoError(t, err)
	require.Equal(t, length, len(password))

	charsets := []string{opts.LowerChars, opts.UpperChars, opts.NumberChars, opts.SpecialChars}
	for _, charset := range charsets {
		require.True(t, common.ContainsCharFromCharset(password, charset))
	}
}

func TestGenerateRandomSecurePasswordTooShort(t *testing.T) {
	length := 3
	opts := DefaultPasswordOptions()
	password, err := GenerateRandomSecurePassword(length, opts)
	require.Error(t, err)
	require.Equal(t, errors.ErrPasswordTooShort, err)
	require.Equal(t, "", password)
}
