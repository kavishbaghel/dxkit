package helpers

import (
	"fmt"
	"os"
	"strings"
)

func DetectShell() (string, error) {

	/* Check if shellType is provided as an argument. If it is provided, use it.
	Otherwise, try to detect the shell from the environment. */

	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", fmt.Errorf("Unable to detect shell. Please provide the shell type as an argument (bash or zsh).")
	}
	if strings.Contains(shell, "bash") {
		return "bash", nil
	} else if strings.Contains(shell, "zsh") {
		return "zsh", nil
	} else {
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}
