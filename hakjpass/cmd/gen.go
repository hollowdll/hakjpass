package cmd

import (
	"fmt"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdGen = &cobra.Command{
		Use:   "gen",
		Short: "Generate password",
		Long: `Generate a random secure password with the specified length.
Minimum length is 4. If the length is not specified the default length 20 is used.
The generated password includes at least one upper case letter, lower case letter,
number and special character.
`,
		Example: `# Generate random secure password with default length 20
hakjpass gen

# Generate random secure password with length 30
hakjpass gen --length 30`,
		Run: func(cmd *cobra.Command, args []string) {
			generateRandomSecurePassword()
		},
	}
	passwordLength int = 20
)

func init() {
	cmdGen.Flags().IntVarP(&passwordLength, "length", "l", 20, "Length of the password. Minimum is 4")
}

func generateRandomSecurePassword() {
	if passwordLength < hakjpass.MinPasswordLength {
		fmt.Printf("Minimum length is %d\n", hakjpass.MinPasswordLength)
	} else {
		fmt.Printf("Generating a random secure password with length %d...\n", passwordLength)
		generatedPassword, err := hakjpass.GenerateRandomSecurePassword(passwordLength, hakjpass.DefaultPasswordOptions())
		if err != nil {
			fmt.Printf("Error generating password: %v\n", err)
		}
		fmt.Printf("%s\n", generatedPassword)
	}
}
