package hakjpass

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"github.com/hollowdll/hakjpass/errors"
	"golang.org/x/crypto/pbkdf2"
)

const (
	keySize    = 32
	saltSize   = 16
	iterations = 100000
)

// GenerateEncryptionKeyWithPassword generates an AES256 encryption key.
// The returned key is encrypted with the password and is a combination of
// base64 encoded salt and key separated with colon (:).
func GenerateEncryptionKeyWithPassword(password string) (string, error) {
	key, err := generateRandomBytes(keySize)
	if err != nil {
		return "", err
	}
	salt, err := generateRandomBytes(saltSize)
	if err != nil {
		return "", err
	}

	derivedKey := deriveKey(password, salt)
	ciphertext, err := encryptData(key, derivedKey)
	if err != nil {
		return "", err
	}

	return combineSaltAndCiphertext(salt, ciphertext), nil
}

// DecryptEncryptionKeyWithPassword decrypts the encypted encryption key with the password
// from the salt and key combination. It returns the plaintext key.
func DecryptEncryptionKeyWithPassword(encryptedKeyCombination string, password string) ([]byte, error) {
	salt, ciphertext, err := splitCombinedSaltAndCiphertext(encryptedKeyCombination)
	if err != nil {
		return nil, err
	}

	derivedKey := deriveKey(password, salt)
	key, err := decryptData(ciphertext, derivedKey)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func generateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}

func deriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)
}

func combineSaltAndCiphertext(salt []byte, ciphertext []byte) string {
	return base64.StdEncoding.EncodeToString(salt) + ":" + base64.StdEncoding.EncodeToString(ciphertext)
}

func splitCombinedSaltAndCiphertext(combination string) ([]byte, []byte, error) {
	substrings := strings.Split(combination, ":")
	if len(substrings) == 2 {
		salt, err := base64.StdEncoding.DecodeString(substrings[0])
		if err != nil {
			return nil, nil, err
		}
		ciphertext, err := base64.StdEncoding.DecodeString(substrings[1])
		if err != nil {
			return nil, nil, err
		}
		return salt, ciphertext, nil
	} else {
		return nil, nil, errors.ErrEncryptedKeyInvalid
	}
}

func encryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func decryptData(ciphertext []byte, key []byte) ([]byte, error) {
	if len(ciphertext) <= 0 {
		return []byte{}, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.ErrCiphertextTooShort
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
