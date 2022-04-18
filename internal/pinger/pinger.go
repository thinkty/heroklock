package pinger

import (
	"fmt"
	"log"
	"time"

	"github.com/thinkty/heroklock/internal/timeblock"
)

// Start the pinger with the given timeblocks and interval (seconds)
func Start(timeblocks *timeblock.Timeblocks, interval time.Duration) {

	log.Printf("Starting pinger with an interval of %d min %d sec", interval/60, interval%60)

	// Interval loop from https://pkg.go.dev/time#Tick
	c := time.Tick(time.Second * interval)
	for range c {
		timeblocks.Iteration++
		log.Printf(status(interval, timeblocks))

		validURLs := timeblocks.Check()

		// TODO: send http request to all the urls
		for _, url := range validURLs {
			log.Print(url)
		}
	}
}

// Function to log the current duration and timeblock status
func status(interval time.Duration, timeblocks *timeblock.Timeblocks) string {
	log := fmt.Sprintf("Interval: ")

	if interval < 60 {
		log += fmt.Sprintf("%d sec", interval)
	} else {
		log += fmt.Sprintf("%d min %d sec", interval/60, interval%60)
	}

	log += fmt.Sprintf(", Iteration: %d, Length: %d\n%s\n", timeblocks.Iteration, timeblocks.Length, timeblocks)
	return log
}
