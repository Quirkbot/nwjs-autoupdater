// +build windows

package wait

import (
	"os"
	"log"
	"time"
)

func Waitpid(pid int, logger *log.Logger) {
	exited := false
	for !exited {
		process, err := os.FindProcess(pid)
		if err != nil {
			logger.Fatal(err)
		}
		if process == nil {
			exited = true
		}
		if !exited {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
