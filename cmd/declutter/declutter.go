package declutter

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/eddie023/declutter/internal"
)

func MoveFiles(cmd string) int {
	dir := ReadDir(cmd)

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

	return len(filteredDir)
}

func ReadDir(path string) []os.FileInfo {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if len(dir) == 0 {
		log.Print("Empty Directory")

		return []os.FileInfo{}
	}

	return dir
}
