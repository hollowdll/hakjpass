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
)

func init() {
	CmdGroup.AddCommand(cmdGroupLs)
}
