package cmd

import (
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
	rootCmd.AddCommand(cmdGen)

	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
