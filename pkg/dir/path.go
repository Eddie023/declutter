package dir

import (
	"os"
	"path/filepath"
)

// Check if cmd arg is provided.
// If provided then use that as a path; else use absolute path.
func GetDirPath(cmd string) (path string) {
	if cmd == "." {
		path = getAbsPath()
	} else {
		path = cmd
	}
	return path
}

// Get absolute path of the executable file.
// NOTE: This will not work for `go run main.go` cmd as in that case the executable file is created in tmp dir.
func getAbsPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(ex)

	return path
}
