package wait

import (
	"os"
	"log"
	"time"
	"syscall"
)

func Waitpid(pid int, logger *log.Logger) {
	exited := false
	for !exited {
		process, err := os.FindProcess(pid)
		if err != nil {
			/*
			Unix systems never return an error even if process doesn't
			exist. So if it gets an error, it's because it's on Windows
			and haven't found the process. In this case, consider exited.
			*/
			logger.Fatal(err)
			exited = true
		} else {
			/*
			Trying to send a `0` signal to the process won't send any
			signal but will still perform error checking.
			*/
			err := process.Signal(syscall.Signal(0))
			logger.Print("Sent signal to process: ", err)
			/*
			If there are no errors the process is still running.
			Otherwise it might be already finished, never started or
			the belongs to other owner.
			*/
			if err != nil {
				if err.Error() == "os: process already finished" {
					exited = true
				}
				if err.Error() == "operation not permitted" {
					exited = false
				}
			} else {
				exited = false
			}
		}
		if !exited {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
