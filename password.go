package hakjpass

import (
	"github.com/google/uuid"
	passwordstoragepb "github.com/hollowdll/hakjpass/pb"
	"google.golang.org/protobuf/proto"
)

func NewPasswordEntry(password string, username string, description string, group string) (*passwordstoragepb.PasswordEntry, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &passwordstoragepb.PasswordEntry{
		Id:          id.String(),
		Password:    password,
		Username:    "",
		Description: "",
		Group:       group,
	}, nil
}

func FindPasswordEntryById(passwordEntries []*passwordstoragepb.PasswordEntry, id string) *passwordstoragepb.PasswordEntry {
	for _, passwordEntry := range passwordEntries {
		if passwordEntry.Id == id {
			return passwordEntry
		}
	}
	return nil
}

func FindPasswordEntriesByGroup(passwordEntries []*passwordstoragepb.PasswordEntry, group string) []*passwordstoragepb.PasswordEntry {
	matchingPasswordEntries := []*passwordstoragepb.PasswordEntry{}
	for _, passwordEntry := range passwordEntries {
		if passwordEntry.Group == group {
			matchingPasswordEntries = append(matchingPasswordEntries, passwordEntry)
		}
	}
	return matchingPasswordEntries
}

func serializePasswordEntryListToBinary(passwordEntryList *passwordstoragepb.PasswordEntryList) ([]byte, error) {
	data, err := proto.Marshal(passwordEntryList)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func deserializePasswordEntryListFromBinary(data []byte) (*passwordstoragepb.PasswordEntryList, error) {
	passwordEntryList := &passwordstoragepb.PasswordEntryList{}
	err := proto.Unmarshal(data, passwordEntryList)
	if err != nil {
		return nil, err
	}
	return passwordEntryList, nil
}
