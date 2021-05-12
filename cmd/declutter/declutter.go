package declutter

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/eddie023/declutter/internal"
)

type outputFoldersMap map[string][]string

func MoveFiles(cmd string) int {
	outputFolders := getOutputFolders()

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
		if shouldMoveFile(fType, outputFolders) {
			outputFolderFilePath, ok := mapkey(outputFolders, fType)
			if !ok {
				panic("Output type doesn't exist")
			}

			oldpath, _ := filepath.Abs(fName)
			newpath, err := filepath.Abs(outputFolderFilePath)
			if err != nil {
				log.Fatal("Can't create new path", err)
			}
			newpath = newpath + "/" + fName

			err = MoveFile(oldpath, newpath)
			if err != nil {
				log.Fatal("Error when moving files", err)
			}
		}
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
		log.Print("Empty Directory\n")

		return []os.FileInfo{}
	}

	return dir
}

func getOutputFolders() outputFoldersMap {
	var c internal.Conf
	config := c.GetConf()
	outputFolder := config.Output

	for folder := range outputFolder {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			fmt.Printf("Folder name: %s doesn't exit\n", folder)

			err := os.Mkdir(folder, 0755)
			if err != nil {
				log.Fatal("Error when creating new folder\n", err)
			}
		} else {
			// If folder exist
			fmt.Printf("Folder name: %s already exists\n", folder)
		}
	}

	return outputFolder
}

func shouldMoveFile(t string, o outputFoldersMap) bool {
	for _, v := range o {
		for _, fileType := range v {
			if fileType == t {
				fmt.Println("File type matched", fileType, t)

				return true
			}
		}
	}
	return false
}

func MoveFile(oldpath string, newpath string) error {
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
