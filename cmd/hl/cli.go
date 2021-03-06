package main

import (
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
)

// Flags
type cliFlags struct {
	nocolor bool
	sample  bool
	table   bool
	version bool
}

var flags cliFlags

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf(
		"%s [OPTIONS] [color1]... [colorN]",
		os.Args[0],
	)
	cli.BugEmail = "hilighter.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of an error",
			"the exit status will be one of the below:\n\n",
			"  1: Invalid option\n",
			"  2: Missing argument\n",
			"  3: Exception\n",
			"  4: Error reading stdin",
		},
		" ",
	)
	cli.Info = strings.Join(
		[]string{
			"Hilights the text from stdin using the methods passed",
			"on the CLI.",
		},
		" ",
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
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	// Short circuit if version was requested
	if flags.version {
		hl.Printf("hilighter version %s\n", hl.Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if !flags.sample && !flags.table && (cli.NArg() == 0) {
		cli.Usage(MissingArgument)
	} else if (flags.sample || flags.table) && (cli.NArg() != 0) {
		cli.Usage(InvalidOption)
	} else if flags.sample && flags.table {
		cli.Usage(InvalidOption)
	}
}
