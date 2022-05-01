package pinger

import (
	"fmt"
	"log"
	"net/http"
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

		// Ping all the urls
		for _, url := range validURLs {
			go ping(url)
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

	log += fmt.Sprintf(", Iteration: %d, Length: %d", timeblocks.Iteration, timeblocks.Length)
	return log
}

// This does not actually ping the server as heroku servers don't have the
// ability to handle ping by default. So, it just sends an HTTP GET request to
// the given address
func ping(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	fmt.Printf("\t%s\t%s\n", url, resp.Status)
}
