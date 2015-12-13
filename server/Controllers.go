package server

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ControllerKeepAlive(w http.ResponseWriter, r *http.Request) {
	Alive += 1
	fmt.Fprintf(w, "ALIVE")
}

func ControllerExit(w http.ResponseWriter, r *http.Request) {
	println("EXIT!")
	os.Exit(1)
}

func ControllerDrives(w http.ResponseWriter, r *http.Request) {
	out, _ := exec.Command("wmic", "logicaldisk", "where", "drivetype=2", "get", "deviceid").CombinedOutput()
	rege := regexp.MustCompile("[A-Z]:")
	drives := rege.FindAllString(string(out), -1)
	fmt.Println(drives)
	// Format: "A,B,C,D,E"
	fmt.Fprintf(w, strings.Join(drives, ","))
}

func ControllerFormat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usb := r.FormValue("usb")

	println("Formatting", usb)
	out, err := exec.Command("format", "/q", "/X", usb).CombinedOutput()
	fmt.Println(out)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		fmt.Fprintf(w, "ERR")
		return
	}

	fmt.Println("Sending OK")
	fmt.Fprintf(w, "OK")
}

func ControllerWrite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usb := r.FormValue("usb")
	data := r.FormValue("data")

	println("usb: ", usb)
	println("data: ", data)

	data = strings.Split(data, ",")[1]
	println("Writing", data, "into", usb)

	file_path := fmt.Sprintf("%s\\medpass.png", usb)

	decoded, err := base64.StdEncoding.DecodeString(data)
	if	err != nil {
		fmt.Println("ERROR:", err.Error())
		fmt.Fprintf(w, "ERR")
		return
	}

	// 0444 means read only, for everyone
	err = ioutil.WriteFile(file_path, decoded, 0444)
	if	err != nil {
		fmt.Println("ERROR:", err.Error())
		fmt.Fprintf(w, "ERR")
		return
	}

	fmt.Println("Sending OK")
	fmt.Fprintf(w, "OK")
}
