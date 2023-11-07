package file

import (
	"testing"
)

func TestIsValidPath(t *testing.T) {

	testCases := []struct {
		want bool
		path string
	}{
		{want: false, path: "/invalid/path"},
		{want: true, path: "/Users/eddie/dev/go"}}

	for _, test := range testCases {

		got := IsValidPath(test.path)

		if got != test.want {
			t.Errorf("Got %v, wanted %v", got, test.want)
		}
	}
}
