package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func startPomodoro(cmd *cobra.Command, args []string) {
	fmt.Println("Pomodoro started! Work for 25 minutes.")
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start a 25-minute Pomodoro timer",
		Run:   startPomodoro,
	})
}
