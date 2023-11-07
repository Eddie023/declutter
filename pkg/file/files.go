package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Filter slice of files based on provided string.
func FilterByName(ss []os.DirEntry, test func(string) bool) (ret []os.DirEntry) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s)
		}
	}
	return ret
}

// Filter slice of files based on file properties.
// For example: This can be used to filter files that are directory.
func FilterByFileInfo(ss []os.DirEntry, test func(os.DirEntry) bool) (ret []os.DirEntry) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

// Function to filter strings with "." at the beginning. i.e hidden files
func FilterHiddenFiles(f []os.DirEntry) []os.DirEntry {
	ss := func(fName string) bool { return !strings.HasPrefix(string(fName), ".") }

	return FilterByName(f, ss)
}

func Move(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// Returns the list of files available in provided path.
// Filter hidden files and sub-directories before returning.
func ReadFiles(path string) []os.DirEntry {
	var files []os.DirEntry
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	if len(dir) == 0 {
		log.Println("Empty Directory")

		return []os.DirEntry{}
	}

	// Filter nested directories.
	ss := func(f os.DirEntry) bool { return !f.IsDir() }

	files = FilterByFileInfo(dir, ss)
	files = FilterHiddenFiles(files)

	return files
}

// Move file based on provided path, file type, and file name.
func MoveFile(path string, fileName string, mimeType string) {
	outputFolderName := GetFolderName(mimeType)

	// check if we already have folder with this name in given path.
	// create if doesn't exist.
	if _, err := os.Stat(path + "/" + outputFolderName); os.IsNotExist(err) {
		log.Println("Creating folder: ", outputFolderName)

		// FIX ME: os.Mkdir is case insensitive. However, we should know the actual case of key dir.
		err := os.Mkdir(filepath.Join(path, outputFolderName), 0755)
		if err != nil {
			log.Fatal("Error when creating new folder\n", err)
		}
	}

	previousPath := path + "/" + fileName
	newPath := path + "/" + outputFolderName + "/" + fileName

	Move(previousPath, newPath)
}

func IsValidPath(fp string) bool {
	if _, err := os.Stat(fp); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// get the first section of the memetype string.
// for example video/mp4 -> video
func GetFolderName(mType string) string {
	return mType[:strings.IndexByte(mType, '/')]
}
