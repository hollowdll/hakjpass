package hakjpass

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hollowdll/hakjpass/errors"
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

// GetDataDirPath returns the path to the data directory
// where the files used by the program should be in.
func GetDataDirPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, HakjpassDataDirName), nil
}

func writeFile(filepath string, data []byte, perm os.FileMode) error {
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

func encryptAndWriteFile(filepath string, data []byte, perm os.FileMode, key []byte) error {
	encryptedData, err := encryptData(data, key)
	if err != nil {
		return err
	}
	err = writeFile(filepath, encryptedData, perm)
	if err != nil {
		return err
	}
	return nil
}

func readFileAndDecrypt(filepath string, perm os.FileMode, key []byte) ([]byte, error) {
	encryptedData, err := readFile(filepath, perm)
	if err != nil {
		return nil, err
	}
	decryptedData, err := decryptData(encryptedData, key)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

func readFileWithoutCreating(filepath string) ([]byte, error) {
	return os.ReadFile(filepath)
}

func createDataDirIfNotExists() (string, error) {
	dataDir, err := GetDataDirPath()
	if err != nil {
		return "", err
	}

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
	err = encryptAndWriteFile(s.storageFilePath, data, perm, s.encryptionKey)
	if err != nil {
		return err
	}

	return nil
}

func (s *HakjpassStorage) readPasswordStorageFile(perm os.FileMode) (*passwordstoragepb.PasswordEntryList, error) {
	decryptedData, err := readFileAndDecrypt(s.storageFilePath, perm, s.encryptionKey)
	if err != nil {
		return nil, err
	}

	var passwordEntryList *passwordstoragepb.PasswordEntryList = nil
	if len(decryptedData) <= 0 {
		fmt.Println("No password storage found, creating a new one...")
		passwordEntryList = NewPasswordEntryList()
		s.writePasswordStorageFile(PasswordStorageFilePermission, passwordEntryList)
		fmt.Println("Password storage created! You can access it with the encryption key and the password you set")
	} else {
		passwordEntryList, err = deserializePasswordEntryListFromBinary(decryptedData)
		if err != nil {
			return nil, errors.ErrInvalidFileOrPassword
		}
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
		fmt.Printf(`No encryption key found, a new key will be created...
The key will be encrypted with a password that is needed to access the password storage.
If you lose the password or the key, you cannot access the password storage anymore.
In this case you need to create a new password storage file and a new key.
The user is advised to keep the password safe and to backup the key and the password storage file.
It is also possible to rotate the key by using the commands under 'hakjpass key'.
Run command 'hakjpass key --help' for more information.
When rotating the key, the user should replace the password storage file backup with the new file
because the old file cannot be decrypted anymore with the new key.

`)
		password, err = common.PromptPasswordWithPrecedingText("Enter password for the encryption key")
		if err != nil {
			return err
		}
		encryptedKeyStr, err = GenerateEncryptionKeyWithPassword(password)
		if err != nil {
			return err
		}
		err = writeFile(s.encryptionKeyFilePath, []byte(encryptedKeyStr), EncryptionKeyFilePermission)
		if err != nil {
			return err
		}
	} else {
		password, err = common.PromptPasswordWithPrecedingText("Enter password for the encryption key")
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
