// Command declutter is a utility that organizes provided
// directory by moving files to relevant folder.
package declutter

import (
	"log"
	"os"

	"github.com/eddie023/declutter/pkg/config"
	filesvc "github.com/eddie023/declutter/pkg/file"
	"github.com/gabriel-vasile/mimetype"
)

// Get the list of files in provided path
// check the file mtype and move them to correct folder based on memetype.
func Tidy(path string, flags config.Flags) {
	c := &config.Config{
		Path:  path,
		Flags: flags,
	}

	if err := run(c); err != nil {
		log.Fatalf("Failed with err : %v", err)
	}
}

func run(c *config.Config) error {
	// TODO: Don't return os.Exit(-1)
	// return correct error
	if ok := filesvc.IsValidPath(c.Path); !ok {
		os.Exit(-1)
	}

	// filter hidden files, sub-directories.
	files := filesvc.ReadFiles(c.Path)

	for _, file := range files {
		mtype, err := mimetype.DetectFile(c.Path + "/" + file.Name())
		if err != nil {
			log.Printf("Skipping file : %v . Cant figure out the mime type with error: %v", file.Name(), err)

			continue
		}

		// moveFile must be concurrent.
		// can add a logic to show summary of things here.
		filesvc.MoveFile(c.Path, file.Name(), mtype.String())
	}

	return nil
}
