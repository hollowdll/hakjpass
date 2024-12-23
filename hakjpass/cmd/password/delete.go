package password

import (
	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdPasswordDelete = &cobra.Command{
		Use:   "delete",
		Short: "Delete passwords",
		Long: `Delete passwords. It is possible to delete passwords by specifying the password entry id
or to delete all the passwords of a password group.
`,
		Run: func(cmd *cobra.Command, args []string) {
			deletePasswords(cmd)
		},
	}
)

func init() {
	cmdPasswordDelete.Flags().StringVar(&id, "id", "", "Password entry id")
	cmdPasswordDelete.Flags().StringVarP(&group, "group", "g", "", "Password group")
}

func deletePasswords(cmd *cobra.Command) {
	hakjpassStorage, err := hakjpass.NewHakjpassStorage()
	cobra.CheckErr(err)

	if cmd.Flags().Changed("id") {
		err = hakjpassStorage.DeletePasswordById(id)
		cobra.CheckErr(err)
	}
}
