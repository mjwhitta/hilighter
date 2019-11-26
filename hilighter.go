package hilighter

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Boolean to disable all color codes
var Disable = false

// Cached hex to xterm-256 8-bit mappings
var cachedCodes = map[string]string{}

// Boolean to track version
const Version = "1.5.1"

// Various regular expressions
var allCodes = regexp.MustCompile(`\x1b\[([0-9;]*m|K)`)
var bgCodes = regexp.MustCompile(`\x1b\[(4|10)[0-9;]+m`)
var doubleno = regexp.MustCompile(`no_no_`)
var fgCodes = regexp.MustCompile(`\x1b\[[39][0-9;]+m`)
var hexCode = regexp.MustCompile(`(?i)(on_)?([0-9a-f]{6})`)
var iterate = regexp.MustCompile(
	`(\x1b\[([0-9;]*m|K))*[^\x1b](\x1b\[([0-9;]*m|K))*`,
)
var newline = regexp.MustCompile(`\n`)
var notwhitespace = regexp.MustCompile(`\S+`)
var onlyCodes = regexp.MustCompile(`^(\x1b\[([0-9;]+m|K))+$`)
var parseHex = regexp.MustCompile(
	`(?i)^#?([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})?$`,
)
var wrap = regexp.MustCompile(`wrap(_(\d+))?`)

// Convert hex to xterm-256 8-bit value
// https://stackoverflow.com/questions/11765623/convert-hex-hex_to_256bitto-closest-x11-color-number
func hex_to_x256(hex string) string {
	var hasKey bool
	if _, hasKey = cachedCodes[hex]; hasKey {
		return cachedCodes[hex]
	}

	// For simplicity, assume RGB space is perceptually uniform.
	// There are 5 places where one of two outputs needs to be
	// chosen when the input is the exact middle:
	// - The r/g/b channels and the gray value:
	//     - the higher value output is chosen
	// - If the gray and color values have same distance from the
	//   input
	//     - color is chosen

	// Calculate the nearest 0-based color index at 16..231
	var r, g, b = uint64(0), uint64(0), uint64(0)
	var matches = parseHex.FindAllStringSubmatch(hex, -1)
	for _, match := range matches {
		r, _ = strconv.ParseUint(match[1], 16, 8)
		g, _ = strconv.ParseUint(match[2], 16, 8)
		b, _ = strconv.ParseUint(match[3], 16, 8)
	}

	// 0..5 each
	var ir = (r - 35) / 40
	if r < 48 {
		ir = 0
	} else if r < 115 {
		ir = 1
	}
	var ig = (g - 35) / 40
	if g < 48 {
		ig = 0
	} else if g < 115 {
		ig = 1
	}
	var ib = (b - 35) / 40
	if b < 48 {
		ib = 0
	} else if b < 115 {
		ib = 1
	}

	// 0..215 lazy evaluation
	var cidx = (36 * ir) + (6 * ig) + ib + 16

	// Calculate the nearest 0-based gray index at 232..255
	var average = (r + g + b) / 3

	// 0..23
	var gidx = (average - 3) / 10
	if average > 238 {
		gidx = 23
	}

	// Calculate the represented colors back from the index
	var i2cv = []uint64{0, 0x5f, 0x87, 0xaf, 0xd7, 0xff}

	// r/g/b 0..255 each
	var cr = i2cv[ir]
	var cg = i2cv[ig]
	var cb = i2cv[ib]

	// same value for r/g/b 0..255
	var gv = (10 * gidx) + 8

	// Return the one which is nearer to the original rgb values
	var clr_err = math.Pow(float64(cr-r), 2) +
		math.Pow(float64(cg-g), 2) + math.Pow(float64(cb-b), 2)
	var gray_err = math.Pow(float64(gv-r), 2) +
		math.Pow(float64(gv-g), 2) + math.Pow(float64(gv-b), 2)

	if clr_err <= gray_err {
		cachedCodes[hex] = fmt.Sprintf("color_%03d", cidx)
	} else {
		cachedCodes[hex] = fmt.Sprintf("color_%03d", gidx+232)
	}

	return cachedCodes[hex]
}

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
			// Check if hex color code
			var matches = hexCode.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				var clr = hex_to_x256(match[2])
				if strings.HasPrefix(code, "on_") {
					clr = "on_" + clr
				}
				return colorize(clr, str, args...)
			}

			// Check if wrap
			matches = wrap.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				// Determine wrap width, default to 80
				var width = 80
				if len(match) == 3 && len(match[2]) > 0 {
					width, _ = strconv.Atoi(match[2])
				}
				return Wrap(width, str, args...)
			}

			// Otherwise panic
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

func PrintHex(hex string, str string, args ...interface{}) {
	fmt.Print(Hex(hex, str, args...))
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

func PrintlnHex(hex string, str string, args ...interface{}) {
	fmt.Println(Hex(hex, str, args...))
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

func PrintlnOnHex(hex string, str string, args ...interface{}) {
	fmt.Println(OnHex(hex, str, args...))
}

func PrintlnOnRainbow(str string, args ...interface{}) {
	fmt.Println(OnRainbow(str, args...))
}

func PrintlnRainbow(str string, args ...interface{}) {
	fmt.Println(Rainbow(str, args...))
}

func PrintlnWrap(width int, str string, args ...interface{}) {
	fmt.Println(Wrap(width, str, args...))
}

func PrintOnHex(hex string, str string, args ...interface{}) {
	fmt.Print(OnHex(hex, str, args...))
}

func PrintOnRainbow(str string, args ...interface{}) {
	fmt.Print(OnRainbow(str, args...))
}

func PrintRainbow(str string, args ...interface{}) {
	fmt.Print(Rainbow(str, args...))
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
	str = fmt.Sprintf(str, args...)

	var line = ""
	var lines []string
	var words = notwhitespace.FindAllString(str, -1)

	// Loop thru words
	for _, word := range words {
		if len(Plain(line))+len(Plain(word)) > width {
			// Wrap if line would be longer than width
			lines = append(lines, line)
			line = word
		} else if len(line) == 0 {
			// Can't wrap less than a single word
			line = word
		} else {
			// Append word to line
			line += " " + word
		}
	}

	// Ensure last line is not forgotten
	if len(line) != 0 {
		lines = append(lines, line)
	}

	// Join lines and return
	return strings.Join(lines, "\n")
}
