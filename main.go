package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eddie023/declutter/cmd/declutter"
	"github.com/eddie023/declutter/internal"
)

const CURRENT_DIR = "."

func main() {
	var c internal.Conf
	fmt.Println("the config is", c.GetConf())

	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		// If no cmd is passed, then use current path as dir.
		cmd = CURRENT_DIR
	}

	filteredFiles := declutter.MoveFiles(cmd)

	log.Println("There are", filteredFiles, "items in this dir")
}
