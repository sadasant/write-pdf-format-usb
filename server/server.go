package server

import (
	"github.com/rakyll/statik/fs"
	_ "github.com/sadasant/write-pdf-format-usb/server/statik"
	"log"
	"net/http"
	"os"
	"time"
)

const PORT = "8889"

var Alive int = 1
var Verbose bool
var StatikFS http.FileSystem

func init() {
	StatikFS, _ = fs.New()
}

func Start() {
	http.Handle("/", http.FileServer(StatikFS))
	http.HandleFunc("/keep-alive", ControllerKeepAlive)
	http.HandleFunc("/exit", ControllerExit)
	http.HandleFunc("/drives", ControllerDrives)
	http.HandleFunc("/format", ControllerFormat)
	http.HandleFunc("/write", ControllerWrite)

	// Killing all if we're not alive
	go checkAlive()

	logf("Starting web server at %s", Address())
	err := http.ListenAndServe(":"+PORT, mainHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

// Logs the requests, creates the sessions and continues with the next handler.
func mainHandler(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log
		logf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		// Resume
		handler.ServeHTTP(w, r)
	}
}

func Address() string {
	return "localhost:" + PORT
}

func checkAlive() {
	time.Sleep(time.Second * 3)
	Alive -= 1
	if Alive == 0 {
		logf("The window was closed")
		os.Exit(1)
		return
	}
	checkAlive()
}
