package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/eddie023/declutter/internal/build"
	"github.com/eddie023/declutter/pkg/dir"
	"github.com/gabriel-vasile/mimetype"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

var ErrPathUndefined = errors.New("path must be provided")

// TODO:
// 1. `verbose` flag should show debug logs
// 3. `known error must be shown properly

func main() {
	app := cli.App{
		Name:        "declutter",
		Description: "Automatically move files to correct folder based on filetype",
		UsageText:   "declutter [options...] directory",
		Version:     build.Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "show all debug logs",
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
	p := c.Args().First()

	if p == "" {
		return ErrPathUndefined
	}

	foo, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println("os.Getwd()", foo)

	if ok := dir.IsValidPath(p); !ok {
		// log.Errorf("the provided path '%s' does not exist", path)
		return nil
	}

	// filter hidden files, sub-directories.
	files := dir.ReadFiles(p)
	if len(files) == 0 {
		// log.Warnln("No files to move. Use a different location.")
		return nil
	}

	// ctx, err := context.WithTimeout(context.Background(), 100*time.Millisecond)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	eg := errgroup.Group{}
	for _, file := range files {
		mtype, err := mimetype.DetectFile(p + "/" + file.Name())
		if err != nil {
			fmt.Printf("Skipping file: '%s' . Error detecting mimetype: %s", file.Name(), err.Error())
			continue
		}

		go func(filename, mimeType string) {
			eg.Go(func() error {
				err = dir.MoveFile(p, filename, mtype.String())
				if err != nil {
					fmt.Printf("failed to move file %s, skipping", filename)

				}
				return nil
			})

		}(file.Name(), mtype.String())
	}

	err = eg.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Successfully moved")
	return nil
}
