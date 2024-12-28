package key

import (
	"encoding/hex"
	"fmt"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdKeyNew = &cobra.Command{
		Use:   "new",
		Short: "Create a new encryption key",
		Long: `Create a new encryption key. A key requires a password that will be used to encrypt it.
A symmetric encryption key is needed to encrypt and decrypt the password storage file so it can be stored securely.
This command outputs the encrypted key. The user is responsible to remember the password.
The passwords storage file cannot be accessed without the password.
`,
		Run: func(cmd *cobra.Command, args []string) {
			createEncryptionKey()
		},
	}
)

func init() {

}

func createEncryptionKey() {
	encryptionKey, err := hakjpass.GenerateEncryptionKey()
	cobra.CheckErr(err)
	fmt.Println(hex.EncodeToString(encryptionKey))
}
