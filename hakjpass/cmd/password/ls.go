package password

import (
	"fmt"
	"strings"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdPasswordLs = &cobra.Command{
		Use:   "ls",
		Short: "List saved passwords",
		Long: `List saved passwords. If no flags specified, this lists all the saved passwords.
It is possible to list passwords by password group by specifying the group with a flag.
`,
		Run: func(cmd *cobra.Command, args []string) {
			listPasswords(cmd)
		},
	}
)

func init() {
	cmdPasswordLs.Flags().StringVarP(&group, "group", "g", "", "Password group")
}

func listPasswords(cmd *cobra.Command) {
	hakjpassStorage, err := hakjpass.NewHakjpassStorage()
	cobra.CheckErr(err)
	passwordEntries, err := hakjpassStorage.GetPasswords()
	cobra.CheckErr(err)

	var builder strings.Builder
	for _, passwordEntry := range passwordEntries {
		builder.WriteString(
			fmt.Sprintf("ID: %s\nGroup: %s\nUsername: %s\nPassword: %s\nDescription: %s\n\n",
				passwordEntry.Id,
				passwordEntry.Group,
				passwordEntry.Username,
				hidePassword(passwordEntry.Password),
				passwordEntry.Description))
	}
	fmt.Print(builder.String())
}

func hidePassword(password string) string {
	return strings.Repeat("*", len(password))
}
