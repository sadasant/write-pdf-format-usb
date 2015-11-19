package browser

import (
	"os/exec"
)

var Verbose bool

func Open(address string) {
	logf("Opening chrome on mode app at %s", address)
	exec.Command("google-chrome", "--app=http://"+address, "--app-shell-host-windows-bounds", "--window-position=300,50").CombinedOutput()
}
