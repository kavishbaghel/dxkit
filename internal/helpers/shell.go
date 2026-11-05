package helpers

import (
	"fmt"
	"os"
	"strings"
)

func GetShell(shellType string) (string, error) {

	/* Check if shellType is provided as an argument. If it is provided, use it.
	Otherwise, try to detect the shell from the environment. */
	if shellType == "" {
		shell := os.Getenv("SHELL")
		if shell == "" {
			return "", fmt.Errorf("shell not found")
		}
		if strings.Contains(shell, "bash") {
			shellType = "bash"
		} else if strings.Contains(shell, "zsh") {
			shellType = "zsh"
		} else {
			return "", fmt.Errorf("unsupported shell: %s", shell)
		}
	}

	return shellType, nil

}
