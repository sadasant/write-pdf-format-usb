package browser

import (
	"fmt"
	"os/exec"
	"runtime"
)

var Verbose bool

func Open(address string) {
	logf("Opening chrome on mode app at %s", address)

	// The GET param no-cache is just a placeholder but the browser
	// trully stops caching this with that
	address = fmt.Sprintf("http://%s?no-cache=true", address)

	binary := "google-chrome"

	var params []string

	if runtime.GOOS == "windows" {
		binary = "cmd"
		params = append(params, "/C", "start", "chrome")
	}

	params = append(params, []string{
		"-app=" + address,
		"--window-size=600,500",
		"--window-position=300,50",
		"--disable-cache",
	}...)

	_, error := exec.Command(binary, params...).CombinedOutput()
	if error != nil {
		println("browser error:", error.Error())
	}
}
