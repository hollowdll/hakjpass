package hakjpass

import "github.com/google/uuid"

// PasswordEntry is a password entry in the password storage.
type PasswordEntry struct {
	Id          string
	Password    string
	Username    string
	Description string
	Group       string
}

func NewPasswordEntry(password string, username string, description string, group string) (*PasswordEntry, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &PasswordEntry{
		Id:          id.String(),
		Password:    password,
		Username:    "",
		Description: "",
		Group:       group,
	}, nil
}

// PasswordEntryList is a list of password entries
type PasswordEntryList struct {
	passwordEntries []PasswordEntry
}
