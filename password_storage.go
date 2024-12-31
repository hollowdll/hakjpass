package hakjpass

import (
	"path/filepath"
	"strings"

	"github.com/hollowdll/hakjpass/internal/common"
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
	// RotateEncryptionKey decrypts the password storage file
	// and encrypts it again with the new encryption key.
	// After this it replaces the current key with the new key.
	RotateEncryptionKey(newKeyPath string) error
}

type HakjpassStorage struct {
	storageFilePath       string
	storageBackupFilePath string
	encryptionKeyFilePath string
	encryptionKey         []byte
}

func NewHakjpassStorage() (PasswordStorage, error) {
	dataDir, err := createDataDirIfNotExists()
	if err != nil {
		return nil, err
	}

	return &HakjpassStorage{
		storageFilePath:       filepath.Join(dataDir, PasswordStorageFileName),
		storageBackupFilePath: filepath.Join(dataDir, PasswordStorageBackupFileName),
		encryptionKeyFilePath: filepath.Join(dataDir, EncryptionKeyFileName),
		encryptionKey:         []byte{},
	}, nil
}

func (s *HakjpassStorage) SavePassword(passwordEntry *passwordstoragepb.PasswordEntry) error {
	err := s.readEncryptionKeyFile(EncryptionKeyFilePermission)
	if err != nil {
		return err
	}
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
	err := s.readEncryptionKeyFile(EncryptionKeyFilePermission)
	if err != nil {
		return nil, err
	}
	passwordEntryList, err := s.readPasswordStorageFile(PasswordStorageFilePermission)
	if err != nil {
		return nil, err
	}
	return passwordEntryList.PasswordEntries, nil
}

func (s *HakjpassStorage) DeletePasswordById(id string) (bool, error) {
	err := s.readEncryptionKeyFile(EncryptionKeyFilePermission)
	if err != nil {
		return false, err
	}
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
	err := s.readEncryptionKeyFile(EncryptionKeyFilePermission)
	if err != nil {
		return false, err
	}
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
	err := s.readEncryptionKeyFile(EncryptionKeyFilePermission)
	if err != nil {
		return false, err
	}
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

func (s *HakjpassStorage) RotateEncryptionKey(newKeyPath string) error {
	if err := s.readEncryptionKeyFile(EncryptionKeyFilePermission); err != nil {
		return err
	}
	encryptedNewKey, err := readFileWithoutCreating(newKeyPath)
	if err != nil {
		return err
	}
	newKeyPassword, err := common.PromptPasswordWithPrecedingText("Enter password for the new encryption key")
	if err != nil {
		return err
	}
	newKey, err := DecryptEncryptionKeyWithPassword(string(encryptedNewKey), newKeyPassword)
	if err != nil {
		return err
	}
	data, err := s.readPasswordStorageFile(PasswordStorageFilePermission)
	if err != nil {
		return err
	}
	s.encryptionKey = newKey
	err = s.writePasswordStorageFile(PasswordStorageFilePermission, data)
	if err != nil {
		return err
	}
	err = writeFile(s.encryptionKeyFilePath, encryptedNewKey, EncryptionKeyFilePermission)
	if err != nil {
		return err
	}

	return nil
}
