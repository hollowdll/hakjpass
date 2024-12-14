package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cmdGen = &cobra.Command{
		Use:   "gen",
		Short: "Generate password",
		Long: `Generate a random secure password with the specified length.
Minimum length is 12 and is used by default if the length is not specified.
`,
		Example: `# Generate random secure password with default length 12
hakjpass gen

# Generate random secure password with length 30
hakjpass gen --length 30`,
		Run: func(cmd *cobra.Command, args []string) {
			generateRandomSecurePassword()
		},
	}
	passwordLength int = 12
)

func init() {
	cmdGen.Flags().IntVarP(&passwordLength, "length", "l", 12, "Length of the password. Minimum is 12")
}

func generateRandomSecurePassword() {
	if passwordLength < 12 {
		fmt.Printf("Minimum length is 12\n")
	} else {
		fmt.Printf("TODO\n")
		fmt.Printf("Generating a random secure password with length %d...\n", passwordLength)
	}
}
