package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eddie023/declutter/pkg/declutter"
	log "github.com/sirupsen/logrus"
)

const USAGE = `Usage: declutter [options...] <filepath>

Options: 
	-v Show verbose logs. (WIP)
	-r Show what would the output look like without moving files. (WIP)
`

const CURRENT_DIR = "."

var (
	// main operation mode
	isDebugMode = flag.Bool("v", false, "show verbose logs")
	// isReadOnlyMode = flag.Bool("r", false, "show output without moving files")
)

// TODO: Each file can be moved concurrently.

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

	declutter.Tidy(path, getFlags())
}

func getFlags() declutter.Flags {
	flags := make(map[string]bool)

	flags["isDebugMode"] = *isDebugMode
	// flags["isReadOnlyMode"] = *isReadOnlyMode

	return flags
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
