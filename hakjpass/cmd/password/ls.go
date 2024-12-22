package password

import (
	"fmt"
	"strings"

	"github.com/hollowdll/hakjpass"
	passwordstoragepb "github.com/hollowdll/hakjpass/pb"
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
	cmdPasswordLs.Flags().StringVar(&id, "id", "", "Password entry ID")
	cmdPasswordLs.Flags().BoolVar(&showPassword, "show-password", false, "Show password")
}

func listPasswords(cmd *cobra.Command) {
	hakjpassStorage, err := hakjpass.NewHakjpassStorage()
	cobra.CheckErr(err)
	passwordEntries, err := hakjpassStorage.GetPasswords()
	cobra.CheckErr(err)

	if cmd.Flags().Changed("group") {
		passwordEntries = hakjpass.FindPasswordEntriesByGroup(passwordEntries, group)
	}

	if cmd.Flags().Changed("id") {
		passwordEntry := hakjpass.FindPasswordEntryById(passwordEntries, id)
		if passwordEntry != nil {
			passwordEntries = []*passwordstoragepb.PasswordEntry{passwordEntry}
		} else {
			passwordEntries = []*passwordstoragepb.PasswordEntry{}
		}
	}

	var builder strings.Builder
	for _, passwordEntry := range passwordEntries {
		password := hidePassword(passwordEntry.Password)
		if showPassword {
			password = passwordEntry.Password
		}
		builder.WriteString(
			fmt.Sprintf("ID: %s\nGroup: %s\nUsername: %s\nPassword: %s\nDescription: %s\n\n",
				passwordEntry.Id,
				passwordEntry.Group,
				passwordEntry.Username,
				password,
				passwordEntry.Description))
	}
	fmt.Print(builder.String())
}

func hidePassword(password string) string {
	return strings.Repeat("*", len(password))
}
