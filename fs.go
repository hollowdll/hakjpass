package hakjpass

import (
	"os"
	"path/filepath"
)

const (
	PasswordStorageFileName       string = "hakjpass_storage"
	PasswordStorageBackupFileName string = "hakjpass_storage.bak"
	HakjpassDataDirName           string = ".hakjpass-data"
	dataDirPermission                    = 0700
	PasswordStorageFilePermission        = 0600
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
