package key

import "github.com/spf13/cobra"

var (
	CmdKey = &cobra.Command{
		Use:   "key",
		Short: "Manage encryption keys",
		Long:  "Manage encryption keys",
	}
)

func init() {
	CmdKey.AddCommand(cmdKeyNew)
}
