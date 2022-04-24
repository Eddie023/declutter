package internal

import (
	"os"
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

func FilterHiddenFiles(f []os.DirEntry) []os.DirEntry {
	// Function to filter strings with "." at the beginning. i.e hidden files
	ss := func(fName string) bool { return !strings.HasPrefix(string(fName), ".") }

	return FilterByName(f, ss)
}
