package main

import (
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
	declutter.MoveFiles(dirPath)
}
