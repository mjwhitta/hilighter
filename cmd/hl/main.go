package main

import (
	"bufio"
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
)

// Helpers begin

func err(msg string) { hl.PrintlnRed("[!] %s", msg) }
func errx(status int, msg string) {
	err(msg)
	os.Exit(status)
}
func good(msg string)    { hl.PrintlnGreen("[+] %s", msg) }
func info(msg string)    { hl.PrintlnWhite("[*] %s", msg) }
func subinfo(msg string) { hl.PrintlnCyan("[=] %s", msg) }
func warn(msg string)    { hl.PrintlnYellow("[-] %s", msg) }

// Helpers end

var nocolor bool
var sample bool
var table bool
var version bool

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
			"Normally the exit status is 0. In the event of invalid",
			"or missing arguments, the exit status will be 1. If an",
			"error reading stdin occurs, the exit status will be 2.",
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
	cli.Flag(&nocolor, "no-color", false, "Disable colorized output.")
	cli.Flag(
		&sample,
		"s",
		"sample",
		false,
		"Show sample foreground/background colors.",
	)
	cli.Flag(&table, "t", "table", false, "Show the color table.")
	cli.Flag(&version, "V", "version", false, "Show version.")
	cli.Parse()

	// Validate cli flags
	if (!sample && !table && !version && (cli.NArg() == 0)) ||
		((sample || table || version) && (cli.NArg() != 0)) {
		cli.Usage(1)
	}
}

func main() {
	hl.Disable = nocolor

	defer func() {
		if r := recover(); r != nil {
			err(r.(error).Error())
		}
	}()

	if sample {
		hl.Sample()
	} else if table {
		hl.Table()
	} else if version {
		hl.Printf("hilighter version %s\n", hl.Version)
	} else {
		var line string
		var scanner = bufio.NewScanner(os.Stdin)

		// Read line by line
		for scanner.Scan() {
			line = scanner.Text()

			// Apply all specified color codes
			for i := range cli.Args() {
				line = hl.Hilight(cli.Arg(i), line)
			}

			// Print the result
			hl.Println(line)
		}

		if scanner.Err() != nil {
			errx(2, scanner.Err().Error())
		}
	}
}
