package internal

import (
	"os"
)

func FilterByName(ss []os.FileInfo, test func(string) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s)
		}
	}
	return ret
}

func FilterByFileInfo(ss []os.FileInfo, test func(os.FileInfo) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}
