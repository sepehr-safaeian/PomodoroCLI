package cmd

import (
	"PomodoroCLI/internal"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func startPomodoro(cmd *cobra.Command, args []string) {
	duration := 25 * time.Minute
	if len(args) > 0 {
		var err error
		duration, err = time.ParseDuration(args[0])
		if err != nil {
			fmt.Println("Invalid duration format. Using default 25 minutes.")
			duration = 25 * time.Minute
		}
	}

	internal.Timer(duration, "Pomodoro Timer Started!")
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "start [duration]",
		Short: "Start a Pomodoro timer. Default is 25 minutes.",
		Args:  cobra.MaximumNArgs(1),
		Run:   startPomodoro,
	})
}
