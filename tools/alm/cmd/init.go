/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	helpers "github.com/kavishbaghel/dxkit/internal/helpers"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and setup the necessary configurations for your shell environment.",
	Long: `Initialize and setup the necessary configurations for your shell environment.
	You can provide the shell type as an argument (bash or zsh). If no argument is provided, the tool will attempt to detect your shell automatically.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("alm: Initializing setup for your shell environment...\n")

		var shell string
		if len(args) > 0 {
			shell = args[0]
			fmt.Println("alm: Using provided shell type -", shell)
		} else {
			shell, err := helpers.DetectShell()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("alm: Detected shell type -", shell)
		}
		fmt.Println("alm: Setting up configurations for your shell environment...\n")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
