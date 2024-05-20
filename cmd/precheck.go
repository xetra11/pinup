/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// precheckCmd represents the precheck command
var precheckCmd = &cobra.Command{
	Use:   "precheck",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking pinup prerequisites")
		checkForDocker()
	},
}

func checkForDocker() {
	path, err := exec.LookPath("docker")
	fmt.Print("Check for Docker: ")
	if err != nil {
		fmt.Print(color.RedString("failed") + "\n")
		return
	}
	fmt.Print(color.GreenString("passed") + "\n")
	fmt.Print(" (path: " + path + ")\n")
}

func init() {
	rootCmd.AddCommand(precheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// precheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// precheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
