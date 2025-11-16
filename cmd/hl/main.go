package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
)

func err(msg string) {
	fmt.Println(hl.Red("[!]") + " " + msg)
}

func errx(status int, msg string) {
	err(msg)
	os.Exit(status)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				if flags.verbose {
					panic(r)
				}

				errx(Exception, r.Error())
			case string:
				if flags.verbose {
					panic(r)
				}

				errx(Exception, r)
			}
		}
	}()

	var line string
	var scanner *bufio.Scanner

	validate()

	switch {
	case flags.sample:
		for _, line := range hl.Sample() {
			fmt.Println(line)
		}
	case flags.table:
		for _, line := range hl.Table() {
			fmt.Println(line)
		}
	default:
		scanner = bufio.NewScanner(os.Stdin)

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
			errx(Exception, scanner.Err().Error())
		}
	}
}
