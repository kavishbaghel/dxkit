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
	Short: "Initialize and setup the necessary configurations for the tool.",
	Long:  `Initialize and setup the necessary configurations for the tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		shell, err := helpers.GetShell()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Detected shell:", shell)
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
