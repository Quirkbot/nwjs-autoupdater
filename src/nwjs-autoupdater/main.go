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
	var pid int

	flag.StringVar(&bundle, "bundle", "", "Path to the update package")
	flag.StringVar(&instDir, "inst-dir", "", "Path to the application install dir")
	flag.StringVar(&appName, "app-name", "my_app", "Application executable name")
	flag.IntVar(&pid, "waitpid", -1, "Pid of running application that must exit before updater starts")
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
	logger.Print("waitpid: ", pid)


	if pid != -1 {
		logger.Print("Waiting process to exit")
		wait.Waitpid(pid, logger)
	}

	var appExec string;
	err, appExec = updater.Update(bundle, instDir, appName)
	logger.Print("Finish")
	if err != nil {
		logger.Fatal(err)
	}

	open.Start(appExec)
}
