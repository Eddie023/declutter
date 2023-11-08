package tree

import (
	"fmt"

	"github.com/xlab/treeprint"
)

type TreeDisplay struct {
	treeprint          treeprint.Tree
	visitedFoldernames []string
}

func New() *TreeDisplay {
	t := TreeDisplay{
		treeprint:          treeprint.New(),
		visitedFoldernames: []string{},
	}

	return &t
}

func (t *TreeDisplay) Add(fName string, fileType string) {
	for _, f := range t.visitedFoldernames {
		if f == fileType {
			b := t.treeprint.FindByValue(fileType)

			b.AddNode(fName)

			return
		}
	}

	t.treeprint.AddBranch(fileType).AddNode(fName)
	t.visitedFoldernames = append(t.visitedFoldernames, fileType)
}

func (t *TreeDisplay) Display() {
	fmt.Print("Here is how your output will look like")
	fmt.Printf("%s", t.treeprint.String())
}
