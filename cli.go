package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"image"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		output string

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&output, "output", "", "Output File Name")
	flags.StringVar(&output, "o", "", "Output File Name(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if flags.NArg() != 2 {
		log.Printf("invalid number of positional arguments")
		return ExitCodeError
	}

	img1 := read(flags.Arg(0))
	img2 := read(flags.Arg(1))
	img := stitch([]image.Image{img1, img2})

	write(output, img)

	_ = output

	return ExitCodeOK
}
