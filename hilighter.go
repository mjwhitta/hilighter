package hilighter

import (
	"errors"
	"fmt"
	"regexp"
)

// Boolean to disable all color codes
var Disable = false

// Various regular expressions
var allCodes = regexp.MustCompile(`\x1b\[([0-9;]*m|K)`)
var bgCodes = regexp.MustCompile(`\x1b\[(4|10)[0-9;]+m`)
var doubleno = regexp.MustCompile(`no_no_`)
var fgCodes = regexp.MustCompile(`\x1b\[[39][0-9;]+m`)
var iterate = regexp.MustCompile(
	`(\x1b\[([0-9;]*m|K))*[^\x1b](\x1b\[([0-9;]*m|K))*`,
)
var newline = regexp.MustCompile(`\n`)
var onlyCodes = regexp.MustCompile(`^(\x1b\[([0-9;]+m|K))+$`)

// Boolean to track version
const Version = "1.5.1"

func Hilight(code string, str string, args ...interface{}) string {
	// Call the appropriate function
	var hasKey bool
	if _, hasKey = Colors[code]; hasKey {
		return colorize(code, str, args...)
	} else if _, hasKey = Modes[code]; hasKey {
		return modify(code, str, args...)
	} else {
		switch code {
		case "on_rainbow":
			return OnRainbow(str, args...)
		case "plain":
			return Plain(str, args...)
		case "rainbow":
			return Rainbow(str, args...)
		default:
			panic(errors.New("Invalid color or mode: " + code))
		}
	}
}

func Hilights(
	codes []string,
	str string,
	args ...interface{},
) string {
	str = fmt.Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	return str
}

func Plain(str string, args ...interface{}) string {
	// Strip all color codes
	return allCodes.ReplaceAllString(fmt.Sprintf(str, args...), "")
}

func Print(args ...interface{}) {
	fmt.Print(args...)
}

func Printf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

func PrintHilight(code string, str string, args ...interface{}) {
	fmt.Print(Hilight(code, str, args...))
}

func PrintHilights(codes []string, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	// Print the result
	fmt.Print(str)
}

func Println(args ...interface{}) {
	fmt.Println(args...)
}

func PrintlnHilight(code string, str string, args ...interface{}) {
	fmt.Println(Hilight(code, str, args...))
}

func PrintlnHilights(codes []string, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	// Apply all specified color codes
	for i := range codes {
		str = Hilight(codes[i], str)
	}

	// Print the result
	fmt.Println(str)
}

func PrintlnWrap(width int, str string, args ...interface{}) {
	fmt.Println(Wrap(width, str, args...))
}

func PrintWrap(width int, str string, args ...interface{}) {
	fmt.Print(Wrap(width, str, args...))
}

func Sample() {
	// Show all bg/fg combos of the first 16 colors
	var fg, bg string
	for f := 0; f < 16; f++ {
		for b := 0; b < 16; b++ {
			fg = fmt.Sprintf("color_%03d", f)
			bg = fmt.Sprintf("on_color_%03d", b)
			fmt.Print(colorize(fg, colorize(bg, " mw ")))
		}
		fmt.Print("\n")
	}
}

func Sprint(args ...interface{}) string {
	return fmt.Sprint(args...)
}

func Sprintf(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}

func Sprintln(args ...interface{}) string {
	return fmt.Sprintln(args...)
}

func Table() {
	// Show a pretty table of all 8-bit colors
	var bg string
	for i := 0; i < 16; i++ {
		bg = fmt.Sprintf("on_color_%03d", i)
		fmt.Print(
			colorize(
				bg,
				" %s %s ",
				Black("%03d", i),
				White("%03d", i),
			),
		)
		if (i+1)%8 == 0 {
			fmt.Print("\n")
		}
	}
	for i := 16; i < 256; i++ {
		bg = fmt.Sprintf("on_color_%03d", i)
		fmt.Print(
			colorize(
				bg,
				" %s %s ",
				Black("%03d", i),
				White("%03d", i),
			),
		)
		if (i-15)%6 == 0 {
			fmt.Print("\n")
		}
	}
}

func Wrap(width int, str string, args ...interface{}) string {
	// TODO wrap
	return fmt.Sprintf(str, args...)
}
