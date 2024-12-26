package password

import (
	"fmt"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdPasswordEdit = &cobra.Command{
		Use:   "edit ID",
		Short: "Edit a saved password",
		Long: `Edit the fields of a saved password entry. The fields can be modified by giving the ID of the password entry.
The new password, password group, username and description can be modified by specifying them with flags.
The password is always prompted when modifying it.
`,
		Example: `# Prompt to enter the new password
hakjpass password edit 0193fe31-7675-761e-a7b8-5ee4663ddcd1 -p

# Modify only the group
hakjpass password edit 0193fe31-7675-761e-a7b8-5ee4663ddcd1 -g newgroup`,
		Args: cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			editPassword(cmd, args[0])
		},
	}
)

func init() {
	cmdPasswordEdit.Flags().BoolVarP(&enterPassword, "password", "p", false, "Prompt to enter new password")
	cmdPasswordEdit.Flags().StringVarP(&group, "group", "g", "", "New password group")
	cmdPasswordEdit.Flags().StringVarP(&username, "username", "u", "", "New username linked to the password")
	cmdPasswordEdit.Flags().StringVarP(&description, "description", "d", "", "New password description for additional info")
}

func editPassword(cmd *cobra.Command, id string) {
	passwordEntryFields := hakjpass.NewPasswordEntryFields()
	if cmd.Flags().Changed("password") {
		password := promptPassword()
		passwordEntryFields.Password = &password
	}
	if cmd.Flags().Changed("group") {
		passwordEntryFields.Group = &group
	}
	if cmd.Flags().Changed("username") {
		passwordEntryFields.Username = &username
	}
	if cmd.Flags().Changed("description") {
		passwordEntryFields.Description = &description
	}

	hakjpassStorage, err := hakjpass.NewHakjpassStorage()
	cobra.CheckErr(err)

	ok, err := hakjpassStorage.EditPasswordById(id, passwordEntryFields)
	cobra.CheckErr(err)
	if !ok {
		fmt.Println("Password entry not found with the ID")
	}
}
