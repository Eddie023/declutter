package internal

import "os"

func Filter(ss []os.FileInfo, test func(string) bool) (ret []os.FileInfo) {
	for _, s := range ss {
		if test(s.Name()) {
			ret = append(ret, s)
		}
	}
	return ret
}
