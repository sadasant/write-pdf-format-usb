package server

import (
	"fmt"
	"net/http"
	"os"
)

func ControllerDrives(w http.ResponseWriter, r *http.Request) {
	println("DRIVERS!")
	fmt.Fprintf(w, "DRIVERS!")
}

func ControllerKeepAlive(w http.ResponseWriter, r *http.Request) {
	Alive += 1
	fmt.Fprintf(w, "ALIVE")
}

func ControllerExit(w http.ResponseWriter, r *http.Request) {
	println("EXIT!")
	os.Exit(1)
}
