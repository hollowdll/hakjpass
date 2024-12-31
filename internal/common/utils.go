package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func ContainsCharFromCharset(input string, charset string) bool {
	for _, char := range input {
		if strings.ContainsRune(charset, char) {
			return true
		}
	}
	return false
}

// PromptInput prompts the user the prompt in the terminal.
// It reads the user input and returns it.
func PromptInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// PromptPassword prompts the user to enter a password.
// The input is not echoed to the terminal.
func PromptPassword() (string, error) {
	fmt.Print("Password (input hidden): ")
	fd := int(os.Stdin.Fd())
	bytes, err := term.ReadPassword(fd)
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(bytes), nil
}

func PromptPasswordWithPrecedingText(text string) (string, error) {
	fmt.Println(text)
	return PromptPassword()
}

// StringInSlice returns true if target is in slice.
func StringInSlice(target string, slice []string) bool {
	for _, elem := range slice {
		if elem == target {
			return true
		}
	}
	return false
}
