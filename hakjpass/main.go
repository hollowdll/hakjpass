package main

import (
	"os"

	"github.com/hollowdll/hakjpass/hakjpass/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
