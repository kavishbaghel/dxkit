package helpers

import (
	"fmt"
	"os"
)

func GetShell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", fmt.Errorf("shell not found")
	}
	return shell, nil
}
