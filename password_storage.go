package hakjpass

// PasswordStorage is the interface used for managing passwords.
type PasswordStorage interface {
	// SavePassword stores a new password.
	SavePassword(passwordEntry PasswordEntry)
}

type HakjpassStorage struct {
	storageFilePath      string
	symmetricKeyFilePath string
	symmetricKeyPassword string
}

func (s HakjpassStorage) SavePassword(passwordEntry PasswordEntry) {}
