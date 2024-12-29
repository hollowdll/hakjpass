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
It is possible to list passwords by the password entry id or password group using flags --id and --group.
By default the passwords are hidden, but can be shown with --show flag.
`,
		Example: `# List by password group
hakjpass password ls -g group1

# List by id
hakjpass password ls --id 0193fe31-7675-761e-a7b8-5ee4663ddcd1

# Show the password
hakjpass password ls --show`,
		Run: func(cmd *cobra.Command, args []string) {
			listPasswords(cmd)
		},
	}
)

func init() {
	cmdPasswordLs.Flags().StringVarP(&group, "group", "g", "", "Password group")
	cmdPasswordLs.Flags().StringVar(&id, "id", "", "Password entry ID")
	cmdPasswordLs.Flags().BoolVarP(&showPassword, "show", "s", false, "Show password")
	cmdPasswordLs.Flags().BoolVarP(&numberOnly, "number-only", "N", false, "Show only the number of results")
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

	if numberOnly {
		fmt.Printf("Number of results: %d\n", len(passwordEntries))
	} else {
		var builder strings.Builder
		for _, passwordEntry := range passwordEntries {
			password := hidePassword(passwordEntry.Password)
			if showPassword {
				password = passwordEntry.Password
			}
			builder.WriteString(
				fmt.Sprintf("{\n  ID: %s\n  Group: %s\n  Username: %s\n  Password: %s\n  Description: %s\n}\n",
					passwordEntry.Id,
					passwordEntry.Group,
					passwordEntry.Username,
					password,
					passwordEntry.Description))
		}
		fmt.Print(builder.String())
	}
}

func hidePassword(password string) string {
	return strings.Repeat("*", len(password))
}
