// Command declutter is a utility that organizes provided
// directory by moving files to relevant folder.
package declutter

import (
	"os"

	"github.com/eddie023/declutter/internal"
	log "github.com/sirupsen/logrus"
)

// Get the list of files in provided path
// check the file type and move them to correct folder based on config.
func Tidy(path string) {
	if err := run(path); err != nil {
		log.Error("Failed with err", err)
	}
}

func run(p string) error {
	if ok := internal.IsValidPath(p); !ok {
		os.Exit(-1)
	}

	internal.MoveFiles(p)

	return nil
}
