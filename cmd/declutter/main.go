package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eddie023/declutter/pkg/config"
	"github.com/eddie023/declutter/pkg/declutter"
)

const USAGE = `Usage: declutter [options...] <filepath>

Options: 
	-v Show verbose logs. (WIP)
	-r Show what would the output look like without moving files. (WIP)
`

// TODO: config file should be created if not found
// there should be no required for config file
// try to add the declutter binary in cron
// CLI should have good informative debug as we all info logs

var (
	// main operation mode
	isDebugMode    = flag.Bool("v", false, "show verbose logs")
	isReadOnlyMode = flag.Bool("r", false, "show output without moving files")
)

// TODO: Each file can be moved concurrently.

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, USAGE)
	}

	flag.Parse()
	// If no filepath is provided.
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	path := flag.Args()[0]

	declutter.Tidy(path, getFlags())
}

func getFlags() config.Flags {
	flags := make(map[string]bool)

	flags["isDebugMode"] = *isDebugMode
	flags["isReadOnlyMode"] = *isReadOnlyMode

	return flags
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}

	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
