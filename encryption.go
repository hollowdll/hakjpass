package hakjpass

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strings"

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
	encryptedKey, salt, err := encryptEncryptionKeyWithPassword(key, password)
	if err != nil {
		return "", err
	}
	return combineSaltAndEncryptedKey(salt, encryptedKey), nil
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

func combineSaltAndEncryptedKey(salt string, encryptedKey string) string {
	return salt + ":" + encryptedKey
}

func splitCombinedSaltAndEncryptedKey(combination string) (string, string) {
	substrings := strings.Split(combination, ":")
	if len(substrings) == 2 {
		return substrings[0], substrings[1]
	} else {
		return "", ""
	}
}

// encryptEncryptionKeyWithPassword encrypts a key with a password.
// Returns the salt and encrypted key as base64 strings.
func encryptEncryptionKeyWithPassword(key []byte, password string) (string, string, error) {
	salt, err := generateRandomBytes(saltSize)
	if err != nil {
		return "", "", err
	}
	derivedKey := deriveKey(password, salt)
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(key))
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(iv); err != nil {
		return "", "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], key)

	return base64.StdEncoding.EncodeToString(ciphertext), base64.StdEncoding.EncodeToString(salt), nil
}
