package hakjpass

import passwordstoragepb "github.com/hollowdll/hakjpass/pb"

// PasswordStorage is the interface used for managing passwords.
type PasswordStorage interface {
	// SavePassword stores a new password entry.
	SavePassword(passwordEntry *passwordstoragepb.PasswordEntry) error
	// GetPasswords returns password entries.
	GetPasswords() ([]*passwordstoragepb.PasswordEntry, error)
}

type HakjpassStorage struct {
	storageFilePath       string
	storageBackupFilePath string
	symmetricKeyFilePath  string
	symmetricKeyPassword  string
}

func NewHakjpassStorage() *HakjpassStorage {
	return &HakjpassStorage{
		storageFilePath:       "",
		storageBackupFilePath: "",
		symmetricKeyFilePath:  "",
		symmetricKeyPassword:  "",
	}
}

func (s *HakjpassStorage) SavePassword(passwordEntry *passwordstoragepb.PasswordEntry) error {
	storageData, err := readFile(s.storageFilePath)
	if err != nil {
		return err
	}
	passwordEntryList, err := deserializePasswordEntryListFromBinary(storageData)
	if err != nil {
		return err
	}
	passwordEntryList.PasswordEntries = append(passwordEntryList.PasswordEntries, passwordEntry)
	updatedStorageData, err := serializePasswordEntryListToBinary(passwordEntryList)
	if err != nil {
		return err
	}
	err = writeToFile(s.storageFilePath, updatedStorageData, PasswordStorageFilePermission)
	if err != nil {
		return err
	}
	return nil
}

func (s *HakjpassStorage) GetPasswords() ([]*passwordstoragepb.PasswordEntry, error) {
	storageData, err := readFile(s.storageFilePath)
	if err != nil {
		return nil, err
	}
	passwordEntryList, err := deserializePasswordEntryListFromBinary(storageData)
	if err != nil {
		return nil, err
	}
	return passwordEntryList.PasswordEntries, nil
}
