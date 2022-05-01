package internal

import (
	"fmt"
	"os"
	"strings"
)

type outputFoldersMap map[string][]string

type File os.DirEntry

func Move(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
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
