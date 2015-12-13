package server

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ControllerDrives(w http.ResponseWriter, r *http.Request) {
	out, _ := exec.Command("wmic", "logicaldisk", "where", "drivetype=2", "get", "deviceid").CombinedOutput()
	rege := regexp.MustCompile("[A-Z]:")
	drives := rege.FindAllString(string(out), -1)
	fmt.Println(drives)
	// Format: "A,B,C,D,E"
	fmt.Fprintf(w, strings.Join(drives, ","))
}

func ControllerKeepAlive(w http.ResponseWriter, r *http.Request) {
	Alive += 1
	fmt.Fprintf(w, "ALIVE")
}

func ControllerExit(w http.ResponseWriter, r *http.Request) {
	println("EXIT!")
	os.Exit(1)
}
