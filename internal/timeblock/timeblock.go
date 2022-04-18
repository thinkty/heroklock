package timeblock

import (
	"errors"
	"fmt"
	"time"
)

type Timeblock struct {
	Url       string    `json:"URL"`
	Duration  uint16    `json:"Duration"` // in 30 minutes (0 ~ 1460)
	StartTime time.Time `json:"StartTime"`
	Valid     bool      `json:"Valid"`
}

func (timeblock Timeblock) String() string {
	var result string = ""
	result += fmt.Sprintf("\turl:\t%s\n", timeblock.Url)
	result += fmt.Sprintf("\tduration:\t%.2f hrs\n", float32(timeblock.Duration)/2)
	result += fmt.Sprintf("\tstartTime:\t%s\n", timeblock.StartTime)
	result += fmt.Sprintf("\tValid:\t%v\n", timeblock.Valid)

	return result
}

type Timeblocks struct {
	Length    uint64
	Blocks    []Timeblock
	Iteration uint
}

// Add a new timeblock to the list if it has a unique url
func (timeblocks *Timeblocks) Add(url string, duration uint16, startTime time.Time) (Timeblock, error) {
	newTimeblock := Timeblock{Url: url, Duration: duration, StartTime: startTime, Valid: true}

	// Check if url is unique
	for _, timeblock := range timeblocks.Blocks {
		if timeblock.Valid && url == timeblock.Url {
			return timeblock, errors.New("The given URL already exists in the list")
		}
	}

	timeblocks.Blocks = append(timeblocks.Blocks, newTimeblock)
	timeblocks.Length++
	return newTimeblock, nil
}

// Add a new timeblock to the list with current time as start time
func (timeblocks *Timeblocks) AddNow(url string, duration uint16) (Timeblock, error) {
	return timeblocks.Add(url, duration, time.Now())
}

// Remove the timeblock with the given url from the list if it exists.
// Return true if removal was successful.
func (timeblocks *Timeblocks) Remove(url string) bool {
	for _, timeblock := range timeblocks.Blocks {
		if timeblock.Valid && timeblock.Url == url {
			timeblock.Valid = false
			timeblocks.Length--
			return true
		}
	}

	return false
}

// Go through the list and return urls that are active
func (timeblocks *Timeblocks) Check() []string {
	activeUrls := make([]string, 0)

	for _, timeblock := range timeblocks.Blocks {
		// Skip invalid timeblocks
		if !timeblock.Valid {
			continue
		}

		endTime := timeblock.StartTime.Add(30 * time.Minute * time.Duration(timeblock.Duration))

		// Compare each timeblock's (duration + start time) to (the current time)
		if time.Now().After(endTime) {
			timeblock.Valid = false
			timeblocks.Length--

		} else {
			// Append active timeblock's url to the list
			activeUrls = append(activeUrls, timeblock.Url)
		}
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
