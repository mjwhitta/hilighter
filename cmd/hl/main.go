package main

import (
	"bufio"
	"os"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
)

func err(msg string) {
	hl.PrintlnRed("[!] " + msg)
}

func errx(status int, msg string) {
	err(msg)
	os.Exit(status)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			errx(Exception, r.(error).Error())
		}
	}()

	validate()

	if flags.sample {
		for _, line := range hl.Sample() {
			hl.Println(line)
		}
	} else if flags.table {
		for _, line := range hl.Table() {
			hl.Println(line)
		}
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
			errx(Exception, scanner.Err().Error())
		}
	}
}
