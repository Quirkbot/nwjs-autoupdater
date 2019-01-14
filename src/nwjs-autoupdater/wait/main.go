package wait

import (
	"log"
	"time"
	"nwjs-autoupdater/wait/isrunning"
)

func WaitProcess(processId int, logger *log.Logger) {
	running := true
	tries := 0
	for running && tries < 60 {
		running = isrunning.IsRunning(processId, logger)
		logger.Printf("process %s still running: %d\n", processId, running)
		if running {
			time.Sleep(1000 * time.Millisecond)
			tries += 1
		}
	}
}
