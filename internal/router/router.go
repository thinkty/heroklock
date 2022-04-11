package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/thinkty/heroklock/internal/timeblock"
)

func Start(port uint64, path string) {

	// Initialize timeblocks
	timeblocks := new(timeblock.Timeblocks)
	timeblocks.Blocks = make([]timeblock.Timeblock, 0)

	addr := fmt.Sprintf("localhost:%d", port)
	log.Printf("Starting server at %s with files served from %s", addr, path)

	http.Handle("/", http.FileServer(http.Dir(path)))
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) { getList(w, r, timeblocks) })
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) { addToList(w, r, timeblocks) })
	http.HandleFunc("/add-random", func(w http.ResponseWriter, r *http.Request) { addRandom(w, r, timeblocks) })
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) { removeFromList(w, r, timeblocks) })

	// Returns only on server error
	log.Panic(http.ListenAndServe(addr, nil))
}

func getList(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s", r.Method, r.Host, r.RequestURI)
	log.Printf("Requested timeblocks\n%s", timeblocks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblocks)
}

// Parse the url, duration, and start time from the request and append to the list.
// Respond with the newly added time block.
func addToList(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s", r.Method, r.Host, r.RequestURI)

	url, duration, startTime, err := verifyUrl(r.URL.Query())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTimeblock := timeblocks.Add(url, duration, startTime)

	// Respond with the newly added timeblock
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTimeblock)
}

// Check that all the properties exist in the parameter and parse it
func verifyUrl(values url.Values) (string, uint16, time.Time, error) {
	if values.Has("url") && values.Has("duration") && values.Has("startTime") {
		duration, err := strconv.ParseUint(values.Get("duration"), 10, 16)

		if err != nil {
			return "", 0, time.Now(), errors.New("Failed to parse duration from request")
		}

		startTime, err := time.Parse(time.RFC3339, values.Get("startTime"))

		return values.Get("url"), uint16(duration), startTime, nil

	} else {
		return "", 0, time.Now(), errors.New("Request is missing either url, duration or start time")
	}
}

func addRandom(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s", r.Method, r.Host, r.RequestURI)

	const length = 8
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	rand.Seed(time.Now().UnixNano())
	randomUrl := make([]rune, length)
	for i := range randomUrl {
		randomUrl[i] = letters[rand.Intn(len(letters))]
	}

	newTimeblock := timeblocks.Add(fmt.Sprintf("https://%s.com", string(randomUrl)), length, time.Now())

	log.Printf("Added new timeblock\n\n%s\n", newTimeblock)

	// Respond with the newly added timeblock
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTimeblock)
}

func removeFromList(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s", r.Method, r.Host, r.RequestURI)

	// TODO:
}
