package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdPaths = &cobra.Command{
		Use:   "paths",
		Short: "List file paths",
		Long:  `List paths to the files used by the program.`,
		Run: func(cmd *cobra.Command, args []string) {
			listFilepaths()
		},
	}
)

func listFilepaths() {
	dataDir, err := hakjpass.GetDataDirPath()
	cobra.CheckErr(err)
	fmt.Printf("Data directory: %s\n", dataDir)
	fmt.Printf("Password storage file: %s\n", filepath.Join(dataDir, hakjpass.PasswordStorageFileName))
	fmt.Printf("Encryption key file: %s\n", filepath.Join(dataDir, hakjpass.EncryptionKeyFileName))
}
