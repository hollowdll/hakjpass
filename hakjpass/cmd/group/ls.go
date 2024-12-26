package group

import (
	"fmt"
	"strings"

	"github.com/hollowdll/hakjpass"
	"github.com/spf13/cobra"
)

var (
	cmdGroupLs = &cobra.Command{
		Use:   "ls",
		Short: "List password groups",
		Long:  "List password groups. This command lists the groups that the saved passwords belong to.",
		Run: func(cmd *cobra.Command, args []string) {
			listGroups()
		},
	}
)

func init() {
	cmdGroupLs.Flags().BoolVarP(&numberOnly, "number-only", "N", false, "Show only the number of results")
}

func listGroups() {
	hakjpassStorage, err := hakjpass.NewHakjpassStorage()
	cobra.CheckErr(err)
	passwordEntries, err := hakjpassStorage.GetPasswords()
	cobra.CheckErr(err)

	groups := hakjpass.FindPasswordGroups(passwordEntries)
	if numberOnly {
		fmt.Printf("Number of results: %d\n", len(groups))
	} else {
		fmt.Println(strings.Join(groups, "\n"))
	}
}
