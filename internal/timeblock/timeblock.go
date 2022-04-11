package timeblock

import (
	"fmt"
	"time"
)

type Timeblock struct {
	Url       string    `json:"URL"`
	Duration  uint16    `json:"Duration"` // in 30 minutes (0 ~ 1460)
	StartTime time.Time `json:"StartTime"`
}

func (timeblock Timeblock) String() string {
	var result string = ""
	result += fmt.Sprintf("\turl:\t%s\n", timeblock.Url)
	result += fmt.Sprintf("\tduration:\t%.2f hrs\n", float32(timeblock.Duration)/2)
	result += fmt.Sprintf("\tstartTime:\t%s\n", timeblock.StartTime)

	return result
}

type Timeblocks struct {
	Length uint64
	Blocks []Timeblock
}

// Add a new timeblock to the list if it has a unique url
func (timeblocks *Timeblocks) Add(url string, duration uint16, startTime time.Time) Timeblock {
	newTimeblock := Timeblock{Url: url, Duration: duration, StartTime: startTime}

	// TODO: check if url is unique

	timeblocks.Blocks = append(timeblocks.Blocks, newTimeblock)
	timeblocks.Length++
	return newTimeblock
}

// Remove the timeblock with the given url from the list if it exists
func (timeblocks *Timeblocks) Remove(url string) {
	// TODO:
}

// Go through the list and return urls that are active
func (timeblocks *Timeblocks) Check() []string {
	activeUrls := make([]string, 0)

	// TODO:
	for _, timeblock := range timeblocks.Blocks {
		activeUrls = append(activeUrls, timeblock.Url)
	}

	return activeUrls
}

func (timeblocks *Timeblocks) String() string {
	if timeblocks.Length == 0 {
		return "<Empty>"
	}

	var result string = "[\n"

	for _, timeblock := range timeblocks.Blocks {
		result += fmt.Sprint(timeblock)
		result += "\n"
	}
	result += "]"

	return result
}
