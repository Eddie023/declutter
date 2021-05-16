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

func MoveFiles(path string) int {
	var oldPath string
	var newPath string

	files := readFiles(path)
	// Function to filter strings with "." at the beginning. i.e hidden files
	ss := func(fName string) bool { return !strings.HasPrefix(string(fName), ".") }

	// Filter hidden files.
	filteredFiles := internal.FilterByName(files, ss)
	outputFolders := createOutputFolders(path, filteredFiles)

	for _, file := range filteredFiles {
		fName := file.Name()
		fType := fName[strings.LastIndex(fName, ".")+1:]

		// Leave files with no extension as it is.
		if len(fType) == len(fName) {
			fmt.Println(fName, " has no extension")
			continue
		}

		oldPath = filepath.Join(path, fName)

		// Move files to folders specified in config file
		if shouldMoveFile(fType, outputFolders) {
			outputFolderFilePath, ok := mapkey(outputFolders, fType)
			if !ok {
				panic("Output type doesn't exist")
			}

			newPath = filepath.Join(path, outputFolderFilePath, fName)

			err := moveFile(oldPath, newPath)
			if err != nil {
				log.Fatal("Error when moving files", err)
			}
		}
	}

	return len(filteredFiles)
}

func readFiles(path string) []os.FileInfo {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if len(dir) == 0 {
		log.Print("Empty Directory\n")

		return []os.FileInfo{}
	}

	// Filter nested directories.
	ss := func(f os.FileInfo) bool { return !f.IsDir() }

	files := internal.FilterByFileInfo(dir, ss)

	return files
}

// Check if output folders are present, if not then create the folders.
// TODO: Right now all the folders are created even it it's not required.
func createOutputFolders(p string, files []os.FileInfo) outputFoldersMap {
	var c internal.Conf
	config := c.GetConf()
	outputFolders := config.Output

	fmt.Println("the outputFolders is", outputFolders)

	// Required output folders
	reqOutFolders := []string{}

	for _, f := range files {
		fName := f.Name()
		fType := fName[strings.LastIndex(fName, ".")+1:]

		rq, ok := getRequiredFolderName(fType, outputFolders)
		fmt.Println("the rq is", rq, ok)
		if ok {
			reqOutFolders = append(reqOutFolders, rq)
		}
	}

	fmt.Println("the requiredOutFOlders is", reqOutFolders)

	for _, folder := range reqOutFolders {
		pathToFolder := filepath.Join(p, folder)
		if _, err := os.Stat(pathToFolder); os.IsNotExist(err) {
			fmt.Printf("Folder name: %s doesn't exit\n", folder)

			// FIX ME: os.Mkdir is case insensitive. However, we should know the actual case of key dir.
			// TODO: Folder should be created in a correct path
			err := os.Mkdir(filepath.Join(p, folder), 0755)
			if err != nil {
				log.Fatal("Error when creating new folder\n", err)
			}
		} else {
			// If folder exist
			fmt.Printf("Folder name: %s already exists\n", folder)
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
			}
			continue
		}
	}

	return "", false
}

func moveFile(oldpath string, newpath string) error {
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
