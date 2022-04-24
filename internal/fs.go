package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

type outputFoldersMap map[string][]string

type File os.DirEntry

func MoveFile(p string, f File, fType string) {
	// var oldPath string
	// var newPath string

	// outputFolders := createOutputFolders(path, filteredFiles)
}

// Read files present in the provided path.
// Filter the list to exclude dirs present in the given path.

// Check if output folders are present, if not then create the folders.
// TODO: Refactor this function.
func createOutputFolders(p string, files []os.DirEntry) outputFoldersMap {
	reqOutFolderNames := []string{}
	var c Conf
	config := c.GetConf()
	outputFolders := config.Output

	for _, f := range files {
		fName := f.Name()
		fType := fName[strings.LastIndex(fName, ".")+1:]

		rq, ok := getRequiredFolderName(fType, outputFolders)
		if ok {
			if len(reqOutFolderNames) < 1 {
				reqOutFolderNames = append(reqOutFolderNames, rq)
			} else {
				for _, elm := range reqOutFolderNames {
					if elm == rq {
						continue
					} else {
						reqOutFolderNames = append(reqOutFolderNames, rq)
					}
				}
			}
		}
	}

	for _, folder := range reqOutFolderNames {
		pathToFolder := filepath.Join(p, folder)

		// Check if required folder exists or not
		// Create if doesn't exist.
		if _, err := os.Stat(pathToFolder); os.IsNotExist(err) {
			log.Info("Creating folder: ", folder)

			// FIX ME: os.Mkdir is case insensitive. However, we should know the actual case of key dir.
			err := os.Mkdir(filepath.Join(p, folder), 0755)
			if err != nil {
				log.Fatal("Error when creating new folder\n", err)
			}
		}
	}

	return outputFolders
}

func shouldMoveFile(t string, o outputFoldersMap) bool {
	for _, v := range o {
		for _, fileType := range v {
			if fileType == t {
				return true
			}
		}
	}
	return false
}

func getRequiredFolderName(t string, o outputFoldersMap) (rq string, ok bool) {
	for key, v := range o {
		for _, fileType := range v {
			if fileType == t {
				rq = key
				ok = true

				return rq, ok
			}
			continue
		}
	}

	return "", false
}

func moveFilepath(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func mapkey(m map[string][]string, value string) (key string, ok bool) {
	for k, v := range m {
		if contains(v, value) {
			key = k
			ok = true
			return
		}
	}
	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func showFName(f []os.FileInfo) (fName []string) {
	for _, file := range f {
		fName = append(fName, file.Name())
	}

	return
}

func IsValidPath(fp string) bool {
	if _, err := os.Stat(fp); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
