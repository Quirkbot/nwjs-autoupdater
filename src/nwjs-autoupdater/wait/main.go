package wait

import (
	"log"
	"time"
	"nwjs-autoupdater/wait/isrunning"
)

func WaitProcess(processId int, logger *log.Logger) {
	running := true
	tries := 0
	var msg string
	for running && tries < 60 {
		running, msg = isrunning.IsRunning(processId)
		logger.Print("Running: ", running)
		logger.Print("Message: ", msg)
		if running {
			time.Sleep(1000 * time.Millisecond)
			tries += 1
		}
	}
}
