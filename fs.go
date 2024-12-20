package hakjpass

import (
	"os"
	"path/filepath"
)

const (
	PasswordStorageFileName       string = "hakjpass_storage"
	HakjpassDataDirName           string = "hakjpass-data"
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

func readFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func createDataDirIfNotExists() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	dataDir := filepath.Join(configDir, HakjpassDataDirName)

	err = os.MkdirAll(dataDir, dataDirPermission)
	if err != nil {
		return err
	}

	err = os.Chmod(dataDir, dataDirPermission)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(filepath string) (bool, error) {
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
