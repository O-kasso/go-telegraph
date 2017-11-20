package main

import (
	"flag"
	"log"
	"os"

	"github.com/o-kasso/go-telegraph/lib"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("Please specify an IP address")
	}

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}
}
