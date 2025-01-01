package cmd

import (
	"github.com/hollowdll/hakjpass/hakjpass/cmd/group"
	"github.com/hollowdll/hakjpass/hakjpass/cmd/key"
	"github.com/hollowdll/hakjpass/hakjpass/cmd/password"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "hakjpass",
	Short:   "CLI based password manager",
	Long:    "CLI based password manager",
	Version: "0.1.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(password.CmdPassword)
	rootCmd.AddCommand(group.CmdGroup)
	rootCmd.AddCommand(key.CmdKey)
	rootCmd.AddCommand(cmdPaths)

	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
