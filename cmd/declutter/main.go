package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eddie023/declutter/pkg/declutter"
	log "github.com/sirupsen/logrus"
)

const CURRENT_DIR = "."

const USAGE = `Usage: declutter [options...] <filepath>

Options: 
	-c Path to override config file. 
	-v Show debug logs.
	-r Show what would the output look like without moving files.
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE)

	}

	log.SetLevel(log.DebugLevel)

	flag.Parse()
	// If no filepath is provided.
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	path := flag.Args()[0]

	declutter.Tidy(path)
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}

	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
