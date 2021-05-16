package main

import (
	"log"
	"os"

	"github.com/eddie023/declutter/cmd/declutter"
	"github.com/eddie023/declutter/internal"
)

const CURRENT_DIR = "."

func main() {
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		// If no cmd is passed, then use current path as dir.
		cmd = CURRENT_DIR
	}

	dirPath := internal.GetDirPath(cmd)
	// TODO: Don't return ff, but return summary
	filteredFiles := declutter.MoveFiles(dirPath)

	log.Println("There are", filteredFiles, "items in this dir")
}
