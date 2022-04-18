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

	"github.com/thinkty/heroklock/internal/pinger"
	"github.com/thinkty/heroklock/internal/timeblock"
)

func Start(port uint64, path string) {

	// Initialize timeblocks
	timeblocks := new(timeblock.Timeblocks)
	timeblocks.Length = 0
	timeblocks.Iteration = 0
	timeblocks.Blocks = make([]timeblock.Timeblock, 0)

	interval := time.Duration(10) // in seconds
	go pinger.Start(timeblocks, interval)

	http.Handle("/", http.FileServer(http.Dir(path)))
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) { getList(w, r, timeblocks) })
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) { addToList(w, r, timeblocks) })
	http.HandleFunc("/add-random", func(w http.ResponseWriter, r *http.Request) { addRandom(w, r, timeblocks) })
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) { removeFromList(w, r, timeblocks) })

	// Returns only on server error
	log.Printf("Starting server at port %d, files served from %s", port, path)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getList(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s %s%s", r.RemoteAddr, r.Proto, r.Method, r.Host, r.RequestURI)
	log.Printf("Requested timeblocks\n%s", timeblocks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeblocks)
}

// Parse the url, duration, and start time from the request and append to the list.
// Respond with the newly added time block.
func addToList(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s %s%s", r.RemoteAddr, r.Proto, r.Method, r.Host, r.RequestURI)

	url, duration, err := verifyUrl(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url = "https://" + url + ".herokuapp.com"

	newTimeblock, err := timeblocks.AddNow(url, duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with the newly added timeblock
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTimeblock)
}

// Check that all the properties exist in the parameter and parse it.
// URL (url) is the name of the heroku server it should ping.
// Duration (duration) is the amount of time in 30 minutes it should ping for.
func verifyUrl(values url.Values) (string, uint16, error) {
	if values.Has("url") && values.Has("duration") {
		duration, err := strconv.ParseUint(values.Get("duration"), 10, 16)

		if err != nil {
			return "", 0, errors.New("Failed to parse duration from request")
		}

		return values.Get("url"), uint16(duration), nil

	} else {
		return "", 0, errors.New("Request is missing either url or duration")
	}
}

func addRandom(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s %s%s", r.RemoteAddr, r.Proto, r.Method, r.Host, r.RequestURI)

	const length = 8
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	rand.Seed(time.Now().UnixNano())
	randomUrl := make([]rune, length)
	for i := range randomUrl {
		randomUrl[i] = letters[rand.Intn(len(letters))]
	}

	newTimeblock, err := timeblocks.AddNow(fmt.Sprintf("https://%s.herokuapp.com", string(randomUrl)), length)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Added new timeblock\n\n%s\n", newTimeblock)

	// Respond with the newly added timeblock
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTimeblock)
}

func removeFromList(w http.ResponseWriter, r *http.Request, timeblocks *timeblock.Timeblocks) {
	log.Printf("%s %s %s %s%s", r.RemoteAddr, r.Proto, r.Method, r.Host, r.RequestURI)

	if !r.URL.Query().Has("url") {
		http.Error(w, "Request is missing url", http.StatusBadRequest)
		return
	}

	url := r.URL.Query().Get("url")

	success := timeblocks.Remove(url)
	if !success {
		http.Error(w, "Could not find any valid URL with the given URL", http.StatusBadRequest)
		return
	}

	// Send the removed URL back as response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
}
