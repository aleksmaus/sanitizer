package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [show|gen|san] [filename]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t Show detected PII:                          %s show [filename]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t Generate PII replacement configuration:     %s gen [filename]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t Sanitize file with optional configuration:  %s san [filename] [configfile]\n", os.Args[0])
	os.Exit(1)
}

const (
	cmdShow     = "show" // Shows the possible PII in the file
	cmdGen      = "gen"  // Generate the possible replacement JSON
	cmdSanitize = "san"  // Sanitize file

)

var supportedCmds = []string{cmdShow, cmdGen, cmdSanitize}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed %v:\n", err)
		os.Exit(1)
	}
}

func main() {
	cmd, fn, cfgfn, err := parseArgs()
	if err != nil {
		usage()
		exitOnError(err)
	}

	switch cmd {
	case cmdShow:
		err = handleShow(fn)
	case cmdGen:
		err = handleGenerate(fn)
	case cmdSanitize:
		err = handleSanitize(fn, cfgfn)
	}
	exitOnError(err)
}

func parseArgs() (cmd, fn, cfgfn string, err error) {
	if len(os.Args) >= 3 {
		cmd = os.Args[1]
		fn = os.Args[2]
		if len(os.Args) > 3 {
			cfgfn = os.Args[3]
		}
	} else {
		return "", "", "", errors.New("invalid arguments")
	}

	// Check if supported cmd
	if !slices.Contains(supportedCmds, cmd) {
		return "", "", "", fmt.Errorf("unsupported command: %v", cmd)
	}

	return cmd, fn, cfgfn, nil
}
