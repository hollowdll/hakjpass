package hakjpass

import (
	"strings"

	"github.com/google/uuid"
	"github.com/hollowdll/hakjpass/internal/common"
	passwordstoragepb "github.com/hollowdll/hakjpass/pb"
	"google.golang.org/protobuf/proto"
)

type PasswordEntryFields struct {
	Id          *string
	Password    *string
	Group       *string
	Username    *string
	Description *string
}

func NewPasswordEntryFields() *PasswordEntryFields {
	return &PasswordEntryFields{
		Id:          nil,
		Password:    nil,
		Group:       nil,
		Username:    nil,
		Description: nil,
	}
}

func NewPasswordEntry(password string, username string, description string, group string) (*passwordstoragepb.PasswordEntry, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &passwordstoragepb.PasswordEntry{
		Id:          id.String(),
		Password:    strings.TrimSpace(password),
		Username:    strings.TrimSpace(username),
		Description: strings.TrimSpace(description),
		Group:       strings.TrimSpace(group),
	}, nil
}

func NewPasswordEntryList() *passwordstoragepb.PasswordEntryList {
	return &passwordstoragepb.PasswordEntryList{
		PasswordEntries: []*passwordstoragepb.PasswordEntry{},
		Id:              uuid.New().String(),
	}
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

func FindPasswordGroups(passwordEntries []*passwordstoragepb.PasswordEntry) []string {
	groups := []string{}
	for _, passwordEntry := range passwordEntries {
		if passwordEntry.Group != "" && !common.StringInSlice(passwordEntry.Group, groups) {
			groups = append(groups, passwordEntry.Group)
		}
	}
	return groups
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
