package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/skratchdot/open-golang/open"
	"nwjs-autoupdater/updater"
	"nwjs-autoupdater/wait"
)

func main() {
	var bundle, instDir, appName string
	var processId int

	flag.StringVar(&bundle, "bundle", "", "Path to the update package")
	flag.StringVar(&instDir, "inst-dir", "", "Path to the application install dir")
	flag.StringVar(&appName, "app-name", "my_app", "Application executable name")
	flag.IntVar(&processId, "wait", -1, "PID that must exit before updater starts. Optional.")
	flag.Parse()

	cwd, _ := os.Getwd()
	logfile, err := os.Create(filepath.Join(cwd, "updater.log"))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	logger := log.New(logfile, "", log.LstdFlags)
	logger.Print("Start")
	logger.Print("bundle: ", bundle)
	logger.Print("instDir: ", instDir)
	logger.Print("appName: ", appName)

	if processId != -1 {
		logger.Print("Waiting process to exit: ", processId)
		wait.WaitProcess(processId, logger)
	}

	err, appExec := updater.Update(bundle, instDir, appName)
	if err != nil {
		logger.Print("Updater error: ", 	err)
	}
	logger.Print("Finish")

	open.Start(appExec)
}
