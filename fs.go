package hakjpass

import (
	"os"
)

func writeToFile(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
