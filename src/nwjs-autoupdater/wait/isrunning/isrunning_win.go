// +build windows

package isrunning

import (
	"os/exec"
	"strconv"
	"fmt"
	"syscall"
)

func IsRunning(pid int) (bool, string) {
	// Convert interger pid to string
	processId := strconv.Itoa(pid)
	// Create a query for tasklist
	query := fmt.Sprintf("PID eq %s", processId)
	// Create command
	cmd := exec.Command("tasklist.exe", "/nh", "/fi", query)
	// Hide window
	// https://stackoverflow.com/a/48365926
	// https://github.com/golang/go/blob/master/src/syscall/exec_windows.go#L222
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	// Get command output
	cmdOutput, err := cmd.Output()
	// If the command returned with error, assume the process is still running
	if err != nil {
		return true, err.Error()
	}
	// Read 4 first bytes as string
	result := string(cmdOutput[:4])
	// If tasklist returns with an "INFO" message it's because no process is
	// running with that pid
	if result == "INFO" {
		return false, string(cmdOutput)
	}
	return true, string(cmdOutput)
}
