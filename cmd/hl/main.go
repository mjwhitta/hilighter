package main

import (
	"bufio"
	"os"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
)

// Exit status
const (
	Good            int = 0
	InvalidOption   int = 1
	MissingArgument int = 2
	Exception       int = 3
	Stdin           int = 4
)

func main() {
	hl.Disable(flags.nocolor)

	defer func() {
		if r := recover(); r != nil {
			errx(Exception, r.(error).Error())
		}
	}()

	validate()

	if flags.sample {
		hl.Sample()
	} else if flags.table {
		hl.Table()
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
			errx(Stdin, scanner.Err().Error())
		}
	}
}
