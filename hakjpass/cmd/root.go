package cmd

import (
	"github.com/hollowdll/hakjpass/hakjpass/cmd/group"
	"github.com/hollowdll/hakjpass/hakjpass/cmd/key"
	"github.com/hollowdll/hakjpass/hakjpass/cmd/password"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hakjpass",
	Short: "CLI based password manager",
	Long: `hakjpass is a CLI based password manager. You can save, show, edit an delete password entries.
Password entries can be listed with multiple different ways, whether you want to list all of them, only a single one,
or list all the passwords of a password group. When listing password, you can specify whether to show or hide the password.
You can save other data along the password to the password entry such as username, password group and description.

Passwords are managed in a password storage file. This file is encrypted using AES-256 with a symmetric encryption key.
The key is also encrypted and protected with a password. The encryption key and password are needed to access the password storage file.
The key can be rotated so the password storage file will use a new key. Use 'hakjpass key rotate --help' for details.

It is recommended to back up the encryption key and the password storage file using best practices. The password storage file should be
backed up regularly so newly saved passwords won't be lost if something happens to the main file. Keep the files safe!!!

Use 'hakjpass paths' to see the file paths of the different files the program uses. The program expects the files to be found
in the locations listed by this command.

It is also possible to generate random secure passwords with different lengths. Use 'hakjpass password gen --help' for details.

Source code for this is available at https://github.com/hollowdll/hakjpass
`,
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
