package password

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var CmdPassword = &cobra.Command{
	Use:   "password",
	Short: "Manage passwords",
	Long:  "Manage passwords",
}

func init() {
	CmdPassword.AddCommand(cmdGen)
}

func promptPassword() string {
	fmt.Print("Password (input hidden): ")
	fd := int(os.Stdin.Fd())
	bytes, err := term.ReadPassword(fd)
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Println()
	return string(bytes)
}
