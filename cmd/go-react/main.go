package main

import (
	"log"

	"github.com/thinkty/go-react/internal/router"
)

func main() {
	log.Print("Starting go-react...")
	router.Start(9000, "./web/dist")
}
