package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Pomodoro CLI for time management",
	Long: `A Simple and Effective CLI Tool for managing your tasks
using the Pomodoro technique.`,
}

// Execute is the entry point for the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
