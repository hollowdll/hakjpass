package hakjpass

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hollowdll/hakjpass/internal/common"
	passwordstoragepb "github.com/hollowdll/hakjpass/pb"
)

const (
	PasswordStorageFileName       string = "hakjpass_storage"
	PasswordStorageBackupFileName string = "hakjpass_storage.bak"
	EncryptionKeyFileName         string = "hakjpass_storage_key"
	HakjpassDataDirName           string = ".hakjpass-data"
	dataDirPermission                    = 0700
	PasswordStorageFilePermission        = 0600
	EncryptionKeyFilePermission          = 0600
)

func writeToFile(filepath string, data []byte, perm os.FileMode) error {
	err := os.WriteFile(filepath, data, perm)
	if err != nil {
		return err
	}
	return nil
}

func readFile(filepath string, perm os.FileMode) ([]byte, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, perm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	if fileSize == 0 {
		return []byte{}, nil
	}

	data := make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func createDataDirIfNotExists() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dataDir := filepath.Join(homeDir, HakjpassDataDirName)

	err = os.MkdirAll(dataDir, dataDirPermission)
	if err != nil {
		return "", err
	}

	err = os.Chmod(dataDir, dataDirPermission)
	if err != nil {
		return "", err
	}

	return dataDir, nil
}

// FileExists returns true if the given filepath is an existing file.
// non-nil error is returned if this check fails.
func FileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (s *HakjpassStorage) writePasswordStorageFile(perm os.FileMode, passwordEntryList *passwordstoragepb.PasswordEntryList) error {
	data, err := serializePasswordEntryListToBinary(passwordEntryList)
	if err != nil {
		return err
	}
	encryptedData, err := encryptData(data, s.encryptionKey)
	if err != nil {
		return err
	}
	err = writeToFile(s.storageFilePath, encryptedData, perm)
	if err != nil {
		return err
	}
	return nil
}

func (s *HakjpassStorage) readPasswordStorageFile(perm os.FileMode) (*passwordstoragepb.PasswordEntryList, error) {
	encryptedData, err := readFile(s.storageFilePath, perm)
	if err != nil {
		return nil, err
	}
	decryptedData, err := decryptData(encryptedData, s.encryptionKey)
	if err != nil {
		return nil, err
	}
	passwordEntryList, err := deserializePasswordEntryListFromBinary(decryptedData)
	if err != nil {
		return nil, err
	}
	return passwordEntryList, nil
}

func (s *HakjpassStorage) readEncryptionKeyFile(perm os.FileMode) error {
	encryptedKey, err := readFile(s.encryptionKeyFilePath, perm)
	if err != nil {
		return err
	}
	encryptedKeyStr := string(encryptedKey)
	password := ""

	if len(encryptedKeyStr) == 0 {
		fmt.Printf(`No encryption key found, a new key will be created.
The key will be encrypted with a password that is needed to access the password storage.
If you lose the password or the key, you cannot access the password storage anymore.
In this case you need to create a new password storage file and a new key.
The user is advised to keep the password safe and to backup the key and the password storage file.
It is also possible to rotate the key by using the commands under 'hakjpass key'.
Run command 'hakjpass key --help' for more information.
When rotating the key, the user should replace the password storage file backup with the new file
because the old file cannot be decrypted anymore with the new key.\n`)
		password, err = common.PromptEncryptionKeyPassword()
		if err != nil {
			return err
		}
		encryptedKeyStr, err = GenerateEncryptionKeyWithPassword(password)
		if err != nil {
			return err
		}
		err = writeToFile(s.encryptionKeyFilePath, []byte(encryptedKeyStr), EncryptionKeyFilePermission)
		if err != nil {
			return err
		}
	} else {
		password, err = common.PromptEncryptionKeyPassword()
		if err != nil {
			return err
		}
	}

	key, err := DecryptEncryptionKeyWithPassword(encryptedKeyStr, password)
	if err != nil {
		return err
	}
	s.encryptionKey = key

	return nil
}
