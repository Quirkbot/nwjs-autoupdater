// +build linux darwin

package isrunning

import (
	"os"
	"syscall"
	"log"
)

func IsRunning(pid int, logger *log.Logger) bool {
	// Try to find the process
	process, err := os.FindProcess(pid)
	logger.Print("find process: ", pid)
	logger.Print("process: ", process)
	logger.Print("err: ", err)
	if process == nil || err != nil {
		return false
	}
	// Send a `0` signal to process and check the response
	err = process.Signal(syscall.Signal(0))
	if err == nil {
		// No errors means process is running
		return true
	} else {
		if err.Error() == "os: process already finished" {
			return false
		}
		// Any other error probably means process is still running
		return true
	}
}
