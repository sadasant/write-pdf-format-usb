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
	if runtime.GOOS == "windows" {
		binary = "chrome.exe"
	}

	params := []string{
		binary,
		"-app=" + address,
		"--window-size=300,400",
		"--window-position=300,50",
		"--disable-cache",
	}

	exec.Command("google-chrome", params...).CombinedOutput()
}
