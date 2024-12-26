package group

import (
	"github.com/spf13/cobra"
)

var (
	CmdGroup = &cobra.Command{
		Use:   "group",
		Short: "Manage password groups",
		Long:  "Manage password groups",
	}
	numberOnly = false
)

func init() {
	CmdGroup.AddCommand(cmdGroupLs)
}
