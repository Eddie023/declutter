package declutter

import (
	"io/ioutil"
	"log"
	"os"
)

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
