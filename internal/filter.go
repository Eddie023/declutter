package internal

import (
	"log"
	"os"
	"path/filepath"
)

func FilterByName(ss []os.FileInfo, test func(string) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s)
		}
	}
	return ret
}

func FilterByFileInfo(ss []os.FileInfo, test func(os.FileInfo) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

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
