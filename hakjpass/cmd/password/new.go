package password

import (
	"github.com/hollowdll/hakjpass"
	"github.com/hollowdll/hakjpass/internal/common"
	"github.com/spf13/cobra"
)

var (
	cmdPasswordNew = &cobra.Command{
		Use:   "new",
		Short: "Save a new password",
		Long: `Save a new password. Password group, username and description can be specified with flags.
If they are not specified with flags, the user is prompted to enter them. Password is always prompted
and the input is not echoed to the terminal to improve security.
`,
		Run: func(cmd *cobra.Command, args []string) {
			saveNewPassword(cmd)
		},
	}
	group       string
	username    string
	description string
)

func init() {
	cmdPasswordNew.Flags().StringVarP(&group, "group", "g", "", "Password group")
	cmdPasswordNew.Flags().StringVarP(&username, "username", "u", "", "Username linked to the password")
	cmdPasswordNew.Flags().StringVarP(&description, "description", "d", "", "Password description for additional info")
}

func saveNewPassword(cmd *cobra.Command) {
	password := promptPassword()

	if !cmd.Flags().Changed("group") {
		groupInput, err := common.PromptInput("Group: ")
		cobra.CheckErr(err)
		group = groupInput
	}

	if !cmd.Flags().Changed("username") {
		usernameInput, err := common.PromptInput("Username: ")
		cobra.CheckErr(err)
		username = usernameInput
	}

	if !cmd.Flags().Changed("description") {
		descriptionInput, err := common.PromptInput("Description: ")
		cobra.CheckErr(err)
		description = descriptionInput
	}

	passwordEntry, err := hakjpass.NewPasswordEntry(password, username, description, group)
	cobra.CheckErr(err)

	hakjpassStorage, err := hakjpass.NewHakjpassStorage()
	cobra.CheckErr(err)

	err = hakjpassStorage.SavePassword(passwordEntry)
	cobra.CheckErr(err)
}
