package main

import (
	"fmt"
	"os"

	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
)

// Exit status
const (
	Good = iota
	InvalidOption
	MissingOption
	InvalidArgument
	MissingArgument
	ExtraArgument
	Exception
)

// Flags
var flags struct {
	nocolor bool
	sample  bool
	table   bool
	verbose bool
	version bool
}

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = os.Args[0] + " [OPTIONS] [color1]... [colorN]"
	cli.BugEmail = "hilighter.bugs@whitta.dev"

	cli.ExitStatus(
		"Normally the exit status is 0. In the event of an error the",
		"exit status will be one of the below:\n\n",
		fmt.Sprintf("  %d: Invalid option\n", InvalidOption),
		fmt.Sprintf("  %d: Missing option\n", MissingOption),
		fmt.Sprintf("  %d: Invalid argument\n", InvalidArgument),
		fmt.Sprintf("  %d: Missing argument\n", MissingArgument),
		fmt.Sprintf("  %d: Extra argument\n", ExtraArgument),
		fmt.Sprintf("  %d: Exception", Exception),
	)
	cli.Info(
		"Hilights the text from stdin using the methods passed on",
		"the CLI.",
	)

	cli.Title = "Hilighter"

	// Parse cli flags
	cli.Flag(
		&flags.nocolor,
		"no-color",
		false,
		"Disable colorized output.",
	)
	cli.Flag(
		&flags.sample,
		"s",
		"sample",
		false,
		"Show sample foreground/background colors.",
	)
	cli.Flag(
		&flags.table,
		"t",
		"table",
		false,
		"Show the color table.",
	)
	cli.Flag(
		&flags.verbose,
		"v",
		"verbose",
		false,
		"Show stacktrace, if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	hl.Disable(flags.nocolor)

	// Short circuit if version was requested
	if flags.version {
		fmt.Printf("hilighter version %s\n", hl.Version)
		os.Exit(Good)
	}

	// Validate cli flags
	switch {
	case !flags.sample && !flags.table && (cli.NArg() == 0):
		cli.Usage(MissingArgument)
	case (flags.sample || flags.table) && (cli.NArg() != 0):
		cli.Usage(InvalidOption)
	case flags.sample && flags.table:
		cli.Usage(InvalidOption)
	}
}
