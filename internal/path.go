package internal

import (
	"log"
	"os"
	"path/filepath"
)

func GetDirPath(cmd string) (path string) {
	if cmd == "." {
		path = getAbsPath()
	} else {
		path = cmd
	}
	return path
}

func getAbsPath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Dir(ex)

	return path
}
