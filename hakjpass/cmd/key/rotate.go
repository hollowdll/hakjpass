package key

import (
	"fmt"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdKeyRotate = &cobra.Command{
		Use:   "rotate",
		Short: "Rotate an encryption key",
		Long: `Rotate an encryption key. This command takes a new key and rotates to use that.
It decrypts the password storage file with the old key and then encrypts it with the new key.
After that it rewrites the key file with the new key.
`,
		Run: func(cmd *cobra.Command, args []string) {
			rotateKey(cmd)
		},
	}
)

func init() {
	cmdKeyRotate.Flags().StringVarP(&newKeyPath, "new-key", "n", "", "File path to the new key")
}

func rotateKey(cmd *cobra.Command) {
	if cmd.Flags().Changed("new-key") {
		hakjpassStorage, err := hakjpass.NewHakjpassStorage()
		cobra.CheckErr(err)
		err = hakjpassStorage.RotateEncryptionKey(newKeyPath)
		cobra.CheckErr(err)
	} else {
		fmt.Println("No file path to the new key provided")
	}
}
