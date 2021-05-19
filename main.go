package main

import (
	"os"

	"github.com/eddie023/declutter/cmd/declutter"
	"github.com/eddie023/declutter/internal"
	log "github.com/sirupsen/logrus"
)

const CURRENT_DIR = "."

func main() {
	log.SetLevel(log.DebugLevel)
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		// If no cmd is passed, then use current path as dir.
		log.Debug("No cmd arg provided. Using current dir as arg")
		cmd = CURRENT_DIR
	}

	dirPath := internal.GetDirPath(cmd)
	declutter.MoveFiles(dirPath)
}
