package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
)

// Helpers begin

func err(msg string) { fmt.Println(hl.Red("[!] %s", msg)) }
func errx(status int, msg string) {
	err(msg)
	os.Exit(status)
}
func good(msg string)    { fmt.Println(hl.Green("[+] %s", msg)) }
func info(msg string)    { fmt.Println(hl.White("[*] %s", msg)) }
func subinfo(msg string) { fmt.Println(hl.Cyan("[=] %s", msg)) }
func warn(msg string)    { fmt.Println(hl.Yellow("[-] %s", msg)) }

// Helpers end

var nocolor bool
var sample bool
var table bool
var version bool

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = fmt.Sprintf(
		"%s [OPTIONS] [color1]... [colorN]",
		os.Args[0],
	)
	cli.BugEmail = "hilighter.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of invalid",
			"or missing arguments, the exit status will be non-zero.",
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
		fmt.Printf("Version: %s\n", hl.Version)
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
			fmt.Println(line)
		}

		if scanner.Err() != nil {
			errx(1, scanner.Err().Error())
		}
	}
}
