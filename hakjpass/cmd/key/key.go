package key

import "github.com/spf13/cobra"

var (
	CmdKey = &cobra.Command{
		Use:   "key",
		Short: "Manage encryption keys",
		Long:  "Manage encryption keys",
	}
	newKeyPath string
)

func init() {
	CmdKey.AddCommand(cmdKeyNew)
	CmdKey.AddCommand(cmdKeyRotate)
}
