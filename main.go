package main

import (
	"github.com/sadasant/write-pdf-format-usb/browser"
	"github.com/sadasant/write-pdf-format-usb/server"
	"os"
)

func main() {
	// Verbosity
	Verbose := os.Getenv("VERBOSE")
	if Verbose != "" {
		server.Verbose = true
		browser.Verbose = true
	}
	// open chrome tab
	go browser.Open(server.Address())
	// init server
	server.Start()
}

func initServer() {
}
