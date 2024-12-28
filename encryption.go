package hakjpass

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keySize    = 32
	saltSize   = 16
	iterations = 100000
)

func GenerateEncryptionKey() ([]byte, error) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GenerateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}

func deriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)
}

func encryptEncryptionKeyWithPassword(key []byte, password string) (string, string, error) {
	salt, err := GenerateRandomBytes(saltSize)
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
