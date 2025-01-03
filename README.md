# hakjpass
Secure CLI based password manager with useful features.

You can save, show, edit an delete password entries.
Password entries can be listed with multiple different ways, whether you want to list all of them, only a single one,
or list all the passwords of a password group. When listing password, you can specify whether to show or hide the password.
You can save other data along the password to the password entry such as username, password group and description.

Passwords are managed in a password storage file. This file is encrypted using AES-256 with a symmetric encryption key.
The key is also encrypted and protected with a password using PBKDF2. The encryption key and password are needed to access the password storage file.

hakjpass also has a command to generate random secure passwords with different lengths.

Note: This tool is designed for technical people who know how to use CLI based apps. It is also the responsibility of the user to safely backup the password storage file and encryption keys. The encryption key password should also be kept safe!

# Install

The tool can be installed with Go. Minimum Go version required is 1.22.

Instructions for installing Go [here](https://go.dev/doc/install)

Run this command to install the binary:
```sh
go install github.com/hollowdll/hakjpass/hakjpass@latest
```

Or install a specific version:
```sh
go install github.com/hollowdll/hakjpass/hakjpass@v0.1.0
```

This installs the binary to your Go bin directory. If you installed Go properly, the Go bin directory should be in your PATH environment variable. You can now use the tool from any directroy in your terminal by using:

```sh
hakjpass
```

If it doesn't work, find instructions how to add Go bin directory to the PATH environment variable.

# Build from source

Building from source also requires having Go installed.

Move to the project root where the go.mod file is and run
```sh
go build -o ./bin/ ./hakjpass/
```
This builds the binary to the ./bin directory creating it if it doesn't exist

Alternatively use the build script which does the same thing
```sh
# give permission to execute
chmod u+x build.sh
# run the script
./build.sh
```

# Getting started

Check the help page for general information
```sh
hakjpass help
```
Every command has its own help page and examples. You can pass --help or -h flag to the commands to see their help page.

When you use the tool for the first time, it creates an encryption key and asks the user to enter the password for it. After this it creates the initial password storage file and encrypts it with the key. It asks the user to enter the password when using the commands.

Now you can start using the tool!

Use the below command to see the expected locations of the files
```sh
hakjpass paths
```
When importing the file backups, you need to place them to the locations shown in the command output. The command also tells you where you can find the files to take backups of them.

A password entry can be saved with
```sh
hakjpass password new
```
It asks to enter the password and optional username, password group and description.

You can list it with
```sh
hakjpass password ls
```

By default it hides the password but you can show it with
```sh
hakjpass password ls --show
```

These are just simple examples to get started. You can do a lot more with the tool than just these. Use the help pages of commands for detailed information.
