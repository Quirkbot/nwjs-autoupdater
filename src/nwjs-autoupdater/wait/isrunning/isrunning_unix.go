// +build linux darwin

package isrunning

import (
	"os"
	"syscall"
)

func IsRunning(pid int) (bool, string) {
	// Try to find the process
	process, err := os.FindProcess(pid)
	if process == nil || err != nil {
		return false, err.Error()
	}
	// Send a `0` signal to process and check the response
	err = process.Signal(syscall.Signal(0))
	if err == nil {
		// No errors means process is running
		return true, "Process still running"
	} else {
		if err.Error() == "os: process already finished" {
			return false, err.Error()
		}
		// Any other error probably means process is still running
		return true, err.Error()
	}
}
