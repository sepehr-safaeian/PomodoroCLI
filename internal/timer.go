package internal

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

const (
	WorkDuration  = 25 * time.Minute
	BreakDuration = 5 * time.Minute
)

func Timer(duration time.Duration, message string) {
	fmt.Println(message)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	startTime := time.Now()
	for remaining := duration - time.Since(startTime); remaining > 0; remaining = duration - time.Since(startTime) {
		s.Suffix = fmt.Sprintf(" Time Remaining: %02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
		time.Sleep(1 * time.Second)
	}

	s.Stop()
	fmt.Println("\nTime's up!")
}

func PomodoroCycle() {
	for {
		Timer(WorkDuration, "Pomodoro Timer Started!")
		Timer(BreakDuration, "Break Time Started!")
	}
}
