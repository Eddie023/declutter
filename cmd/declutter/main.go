package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/eddie023/declutter/internal/build"
	"github.com/eddie023/declutter/pkg/dir"
	"github.com/gabriel-vasile/mimetype"
	"github.com/urfave/cli/v2"
)

var ErrPathUndefined = errors.New("path must be provided")

// TODO:
// 1. `verbose` flag should show debug logs
// 2. `declutter path` must be provided
// 3. `known error must be shown properly
// 4. Add linter

func main() {
	app := cli.App{
		Name:        "declutter",
		Description: "Automatically move files to correct folder based on filetype",
		UsageText:   "declutter [options...] directory",
		Version:     build.Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"d"},
				Usage:   "show all debug logs",
			},
		},
		Action: func(c *cli.Context) error {
			return run(c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	path := c.Args().First()

	if path == "" {
		return ErrPathUndefined
	}

	if ok := dir.IsValidPath(path); !ok {
		// log.Errorf("the provided path '%s' does not exist", path)
		return nil
	}

	// filter hidden files, sub-directories.
	files := dir.ReadFiles(path)
	if len(files) == 0 {
		// log.Warnln("No files to move. Use a different location.")
		return nil
	}

	var wg sync.WaitGroup
	for _, file := range files {
		mtype, err := mimetype.DetectFile(path + "/" + file.Name())
		if err != nil {
			fmt.Printf("Skipping file: '%s' . Error detecting mimetype: %s", file.Name(), err.Error())
			continue
		}

		wg.Add(1)
		go func(filename, mimeType string) {
			defer wg.Done()

			err = dir.MoveFile(path, filename, mtype.String())
			if err != nil {
				fmt.Printf("failed to move file %s, skipping", filename)

			}
		}(file.Name(), mtype.String())
	}

	wg.Wait()
	fmt.Println("Successfully moved")
	return nil
}
