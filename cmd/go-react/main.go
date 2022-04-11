package main

import (
	"log"

	"github.com/thinkty/heroklock/internal/router"
)

func main() {
	log.Print("Starting Heroklock...")
	router.Start(9000, "./web/dist")
}
