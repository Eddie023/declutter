// Command declutter is a utility that organizes provided
// directory by moving files to relevant folder.
package declutter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Get the list of files in provided path
// check the file type and move them to correct folder based on config.
func Tidy(path string) {
	if err := run(); err != nil {
		log.Error("Failed with err", err)
	}
}

func run() error {

	// =========================
	// Configuration

	cfg := struct {
		Foo string
	}{
		Foo: "apple",
	}

	fmt.Println("the cfg", cfg)

	return nil
}
