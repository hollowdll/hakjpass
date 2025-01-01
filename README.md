# hakjpass
Secure CLI based password manager with useful features. Always free and open source!

You can save, show, edit an delete password entries.
Password entries can be listed with multiple different ways, whether you want to list all of them, only a single one,
or list all the passwords of a password group. When listing password, you can specify whether to show or hide the password.
You can save other data along the password to the password entry such as username, password group and description.

Passwords are managed in a password storage file. This file is encrypted using AES-256 with a symmetric encryption key.
The key is also encrypted and protected with a password. The encryption key and password are needed to access the password storage file.

Note: This tool is designed for technical people who know how to use CLI based apps. It is also the responsibility of the user to safely backup the password storage file and encryption keys. The encryption key password should also be kept safe!

# Install

WIP

# Build from source

WIP

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
