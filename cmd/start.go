package cmd

import (
	"PomodoroCLI/internal"

	"github.com/spf13/cobra"
)

func startPomodoro(cmd *cobra.Command, args []string) {
	internal.StartTimer()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start a 25-minute Pomodoro timer",
		Run:   startPomodoro,
	})
}
