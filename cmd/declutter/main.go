package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/eddie023/declutter/internal/build"
	"github.com/eddie023/declutter/pkg/declutter"
	"github.com/eddie023/declutter/pkg/logger"
	"github.com/eddie023/declutter/pkg/tree"
	"github.com/gabriel-vasile/mimetype"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const USAGE = `Usage: declutter [options...] <filepath>

Options: 
	-v Show verbose logs. (WIP)
	-r Show what would the output look like without moving files. (WIP)
`

func main() {
	showDebugLogs := false
	for _, arg := range os.Args {
		if strings.Contains(arg, "verbose") {
			showDebugLogs = true
		}
	}

	var log *zap.SugaredLogger
	if showDebugLogs {
		log = logger.New(logger.WithLogLevel(zap.DebugLevel))
	} else {
		log = logger.New()
	}

	err := run(log, os.Args)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger, args []string) error {
	app := cli.App{
		Name: "declutter",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "show all debug logs",
			},
			&cli.BoolFlag{
				Name:  "readonly",
				Usage: "show the output of the action without moving files",
			},
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "provide the directory location of the folder you want to tidy up.",
			},
			&cli.BoolFlag{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "will show how to output of will look without moving the files",
			},
		},
		Description: "automatically move files to correct folder based on filetype",
		UsageText:   "declutter [options...] <filepath>",
		Version:     build.Version,
		Action: func(c *cli.Context) error {
			path := c.String("path")

			if c.String("path") == "" {
				prompt := promptui.Prompt{
					IsConfirm:   true,
					HideEntered: true,
					Label:       "The location to tidy up is not provided. Do you wanna use the current directory",
				}

				_, err := prompt.Run()
				if err != nil {
					fmt.Println("Exiting")
					return nil
				}

				path = "."
			}

			if ok := declutter.IsValidPath(path); !ok {
				log.Errorf("the provided path '%s' does not exist", path)
				return nil
			}

			// filter hidden files, sub-directories.
			files := declutter.ReadFiles(path)

			if len(files) == 0 {
				log.Warnln("No files to move. Use a different location.")

				return nil
			}

			if c.Bool("show") {
				tree := tree.New()

				for _, file := range files {
					mtype, err := mimetype.DetectFile(path + "/" + file.Name())
					if err != nil {
						log.Warnf("Skipping file : %v . Cant figure out the mime type with error: %v", file.Name(), err)

						continue
					}

					tree.Add(file.Name(), declutter.GetFolderName(mtype.String()))
				}

				tree.Display()

				return nil
			}

			var wg sync.WaitGroup
			for _, file := range files {
				mtype, err := mimetype.DetectFile(path + "/" + file.Name())
				if err != nil {
					log.Warnf("Skipping file : %v . Cant figure out the mime type with error: %v", file.Name(), err)

					continue
				}

				wg.Add(1)

				go func(filename, mimeType string) {
					defer wg.Done()

					err = declutter.MoveFile(path, filename, mtype.String())
					if err != nil {
						log.Debugf("failed to move file %s, skipping", file.Name())

					}
				}(file.Name(), mtype.String())

			}

			wg.Wait()
			log.Infof("Successfully moved")

			return nil
		},
	}

	return app.Run(args)
}
