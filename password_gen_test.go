package hakjpass

import (
	"testing"

	"github.com/hollowdll/hakjpass/errors"
	"github.com/hollowdll/hakjpass/internal/common"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomSecurePassword(t *testing.T) {
	opts := &PasswordOptions{
		LowerChars:   "abc",
		UpperChars:   "ABC",
		NumberChars:  "123",
		SpecialChars: "!?-",
	}

	for _, length := range []int{4, 12, 20, 50, 99, 101, 200, 999} {
		password, err := GenerateRandomSecurePassword(length, opts)
		require.NoError(t, err)
		require.Equal(t, length, len(password))

		charsets := []string{opts.LowerChars, opts.UpperChars, opts.NumberChars, opts.SpecialChars}
		for _, charset := range charsets {
			require.True(t, common.ContainsCharFromCharset(password, charset))
		}
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
