package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	dir := declutter.ReadDir(cmd)

	// Function to filter strings with "." at the beginning. i.e hidden files
	ss := func(fName string) bool { return !strings.HasPrefix(string(fName), ".") }

	// Filter hidden files.
	filteredDir := internal.Filter(dir, ss)

	for _, file := range filteredDir {
		fName := file.Name()
		fType := fName[strings.LastIndex(fName, ".")+1:]

		// Leave files with no extension as it is.
		if len(fType) == len(fName) {
			fmt.Println(fName, " has no extension")
			continue
		}

		// Move files to folders specified in config file
		// fmt.Println(fName, "and", fType)
	}

	fmt.Println(os.Stat("/home/rattlehead/Desktop"))

	log.Println("There are", len(filteredDir), "items in this dir")
}

// func getFileContentType(file os.FileInfo) (string, error) {

// 	buffer := make([]byte, 512)

// 	_, err := file.

// }
