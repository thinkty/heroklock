package router

import (
	"fmt"
	"log"
	"net/http"
)

func Start(port uint64, path string) {
	addr := fmt.Sprintf("localhost:%d", port)
	log.Printf("Starting server at %s", addr)
	http.Handle("/", http.FileServer(http.Dir(path)))

	log.Panic(http.ListenAndServe(addr, nil))
}
