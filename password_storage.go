package hakjpass

import (
	"path/filepath"
	"strings"

	passwordstoragepb "github.com/hollowdll/hakjpass/pb"
)

// PasswordStorage is the interface used for managing passwords.
type PasswordStorage interface {
	// SavePassword stores a new password entry.
	SavePassword(passwordEntry *passwordstoragepb.PasswordEntry) error
	// GetPasswords returns password entries.
	GetPasswords() ([]*passwordstoragepb.PasswordEntry, error)
	// DeletePasswords removes the password entry with the id.
	// Returns false if the password entry is not found.
	DeletePasswordById(id string) (bool, error)
	// DeletePasswordsByGroup removes the password entries in the group.
	// Returns false if the group does not exist.
	DeletePasswordsByGroup(group string) (bool, error)
	// EditPasswordById edits the fields of a password entry.
	// Returns false if the password entry is not found.
	EditPasswordById(id string, passwordEntryFields *PasswordEntryFields) (bool, error)
}

type HakjpassStorage struct {
	storageFilePath       string
	storageBackupFilePath string
	symmetricKeyFilePath  string
	symmetricKeyPassword  string
}

func NewHakjpassStorage() (PasswordStorage, error) {
	dataDir, err := createDataDirIfNotExists()
	if err != nil {
		return nil, err
	}

	return &HakjpassStorage{
		storageFilePath:       filepath.Join(dataDir, PasswordStorageFileName),
		storageBackupFilePath: filepath.Join(dataDir, PasswordStorageBackupFileName),
		symmetricKeyFilePath:  "",
		symmetricKeyPassword:  "",
	}, nil
}

func (s *HakjpassStorage) SavePassword(passwordEntry *passwordstoragepb.PasswordEntry) error {
	passwordEntryList, err := s.readPasswordStorageFile(PasswordStorageFilePermission)
	if err != nil {
		return err
	}
	passwordEntryList.PasswordEntries = append(passwordEntryList.PasswordEntries, passwordEntry)
	err = s.writePasswordStorageFile(PasswordStorageFilePermission, passwordEntryList)
	if err != nil {
		return err
	}
	return nil
}

func (s *HakjpassStorage) GetPasswords() ([]*passwordstoragepb.PasswordEntry, error) {
	passwordEntryList, err := s.readPasswordStorageFile(PasswordStorageFilePermission)
	if err != nil {
		return nil, err
	}
	return passwordEntryList.PasswordEntries, nil
}

func (s *HakjpassStorage) DeletePasswordById(id string) (bool, error) {
	passwordEntryList, err := s.readPasswordStorageFile(PasswordStorageFilePermission)
	if err != nil {
		return false, err
	}
	counter := 0
	found := false
	for _, passwordEntry := range passwordEntryList.PasswordEntries {
		if passwordEntry.Id != id {
			passwordEntryList.PasswordEntries[counter] = passwordEntry
			counter++
		} else {
			found = true
		}
	}
	if !found {
		return false, nil
	}
	passwordEntryList.PasswordEntries = passwordEntryList.PasswordEntries[:counter]
	err = s.writePasswordStorageFile(PasswordStorageFilePermission, passwordEntryList)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *HakjpassStorage) DeletePasswordsByGroup(group string) (bool, error) {
	passwordEntryList, err := s.readPasswordStorageFile(PasswordStorageFilePermission)
	if err != nil {
		return false, err
	}
	counter := 0
	found := false
	for _, passwordEntry := range passwordEntryList.PasswordEntries {
		if passwordEntry.Group != group {
			passwordEntryList.PasswordEntries[counter] = passwordEntry
			counter++
		} else {
			found = true
		}
	}
	if !found {
		return false, nil
	}
	passwordEntryList.PasswordEntries = passwordEntryList.PasswordEntries[:counter]
	err = s.writePasswordStorageFile(PasswordStorageFilePermission, passwordEntryList)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *HakjpassStorage) EditPasswordById(id string, passwordEntryFields *PasswordEntryFields) (bool, error) {
	passwordEntryList, err := s.readPasswordStorageFile(PasswordStorageFilePermission)
	if err != nil {
		return false, err
	}
	found := false
	for _, passwordEntry := range passwordEntryList.PasswordEntries {
		if passwordEntry.Id == id {
			found = true
			if passwordEntryFields.Password != nil {
				passwordEntry.Password = strings.TrimSpace(*passwordEntryFields.Password)
			}
			if passwordEntryFields.Group != nil {
				passwordEntry.Group = strings.TrimSpace(*passwordEntryFields.Group)
			}
			if passwordEntryFields.Username != nil {
				passwordEntry.Username = strings.TrimSpace(*passwordEntryFields.Username)
			}
			if passwordEntryFields.Description != nil {
				passwordEntry.Description = strings.TrimSpace(*passwordEntryFields.Description)
			}
			break
		}
	}
	if !found {
		return false, nil
	}
	err = s.writePasswordStorageFile(PasswordStorageFilePermission, passwordEntryList)
	if err != nil {
		return false, err
	}
	return true, nil
}
