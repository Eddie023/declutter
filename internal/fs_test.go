package internal_test

import (
	"testing"

	"github.com/eddie023/declutter/internal"
)

func TestIsValidPath(t *testing.T) {

	testCases := []struct {
		want bool
		path string
	}{
		{want: false, path: "/invalid/path"},
		{want: true, path: "/Users/eddie/dev/go"}}

	for _, test := range testCases {

		got := internal.IsValidPath(test.path)

		if got != test.want {
			t.Errorf("Got %v, wanted %v", got, test.want)
		}
	}
}
