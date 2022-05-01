// Command declutter is a utility that organizes provided
// directory by moving files to relevant folder.
package declutter

import (
	"os"
	"path/filepath"

	"github.com/eddie023/declutter/internal"
	"github.com/gabriel-vasile/mimetype"
	log "github.com/sirupsen/logrus"
)

type Flags = map[string]bool

type Config struct {
	path  string
	flags Flags
}

// Get the list of files in provided path
// check the file mtype and move them to correct folder based on memetype.
func Tidy(path string, flags Flags) {
	c := &Config{
		path:  path,
		flags: flags,
	}
	if err := run(c); err != nil {
		log.Error("Failed with err", err)
	}
}

func run(c *Config) error {
	if ok := internal.IsValidPath(c.path); !ok {
		os.Exit(-1)
	}

	// filter hidden files, sub-directories.
	files := readFiles(c.path)

	for _, file := range files {
		mtype, err := mimetype.DetectFile(c.path + "/" + file.Name())
		if err != nil {
			log.Warn("Skipping file : %v . Cant figure out the mime type with error: %v", file.Name(), err)
		}

		moveFile(c.path, file.Name(), mtype.String())
	}

	return nil
}

// Returns the list of files available in provided path.
// Filter hidden files and sub-directories before returning.
func readFiles(path string) []os.DirEntry {
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

	files = internal.FilterByFileInfo(dir, ss)
	files = internal.FilterHiddenFiles(files)

	return files
}

// Move file based on provided path, file type, and file name.
func moveFile(path string, fileName string, mimeType string) {

	outputFolderName := internal.GetFolderName(mimeType)

	// check if we already have folder with this name in given path.
	// create if doesn't exist.
	if _, err := os.Stat(path + "/" + outputFolderName); os.IsNotExist(err) {
		log.Info("Creating folder: ", outputFolderName)

		// FIX ME: os.Mkdir is case insensitive. However, we should know the actual case of key dir.
		err := os.Mkdir(filepath.Join(path, outputFolderName), 0755)
		if err != nil {
			log.Fatal("Error when creating new folder\n", err)
		}
	}

	previousPath := path + "/" + fileName
	newPath := path + "/" + outputFolderName + "/" + fileName

	internal.Move(previousPath, newPath)
}
