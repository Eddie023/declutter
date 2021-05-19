package internal

import (
	"os"
)

// Filter slice of files based on provided string.
func FilterByName(ss []os.FileInfo, test func(string) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s)
		}
	}
	return ret
}

// Filter slice of files based on file properties.
// For example: This can be used to filter files that are directory.
func FilterByFileInfo(ss []os.FileInfo, test func(os.FileInfo) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}
