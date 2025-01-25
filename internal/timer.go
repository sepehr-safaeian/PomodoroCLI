package internal

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func StartTimer() {
	duration := 25 * time.Minute
	fmt.Println("Pomodoro Timer Started!")

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	startTime := time.Now()
	for remaining := duration; remaining > 0; remaining -= time.Second {
		elapsed := time.Since(startTime)
		remaining = duration - elapsed

		s.Suffix = fmt.Sprintf(" Time Remaining: %02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
		time.Sleep(1 * time.Second)
	}

	s.Stop()
	fmt.Println("\nTime's up! Take a break.")
}
