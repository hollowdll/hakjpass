package password

import (
	"fmt"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdPasswordDelete = &cobra.Command{
		Use:   "delete",
		Short: "Delete passwords",
		Long: `Delete passwords. It is possible to delete passwords by specifying the password entry id
or to delete all the passwords from a password group.
`,
		Example: `# Delete by id
hakjpass password delete --id 0193fe31-7675-761e-a7b8-5ee4663ddcd1

# Delete all the passwords from a password group
hakjpass password delete -g group1`,
		Run: func(cmd *cobra.Command, args []string) {
			deletePasswords(cmd)
		},
	}
)

func init() {
	cmdPasswordDelete.Flags().StringVar(&id, "id", "", "Password entry ID")
	cmdPasswordDelete.Flags().StringVarP(&group, "group", "g", "", "Password group")
}

func deletePasswords(cmd *cobra.Command) {
	hakjpassStorage, err := hakjpass.NewHakjpassStorage()
	cobra.CheckErr(err)

	if cmd.Flags().Changed("id") {
		ok, err := hakjpassStorage.DeletePasswordById(id)
		cobra.CheckErr(err)
		if !ok {
			fmt.Println("No password entry found with the ID")
		}
	}

	if cmd.Flags().Changed("group") {
		ok, err := hakjpassStorage.DeletePasswordsByGroup(group)
		cobra.CheckErr(err)
		if !ok {
			fmt.Println("Password group not found")
		}
	}
}
