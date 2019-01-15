package updater

import (
	"path/filepath"
	"os"
	"nwjs-autoupdater/unzip"
)

func Update(bundle, instDir, appName string) (error, string) {
	appExecName := appName + ".exe"
	appExec := filepath.Join(instDir, appExecName)

	err := unzip.Unzip(bundle, instDir)
	os.Remove(bundle)
	if err != nil {
		return err, appExec
	}

	return nil, appExec
}
