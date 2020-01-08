package hilighter

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func bgColor(code string, str string, args ...interface{}) string {
	var colorized string

	// Strip all other bg color codes and don't extend bg color over
	// newlines
	str = newline.ReplaceAllString(
		plainBg(str, args...),
		"\x1b["+Colors["on_default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["on_default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
}

func colorize(clr string, str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	// Call the appropriate function
	if strings.HasPrefix(clr, "on_") {
		return bgColor(clr, str, args...)
	}
	return fgColor(clr, str, args...)
}

func fgColor(code string, str string, args ...interface{}) string {
	var colorized string

	// Strip all other fg color codes and don't extend fg color over
	// newlines
	str = newline.ReplaceAllString(
		plainFg(str, args...),
		"\x1b["+Colors["default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
}

// Hex will take a hex color code and add the appropriate ANSI color
// codes to the provided string.
func Hex(hex string, str string, args ...interface{}) string {
	return colorize(hexToXterm256(hex), str, args...)
}

// Convert hex to xterm-256 8-bit value
// https://stackoverflow.com/questions/11765623/convert-hex-hex_to_256bitto-closest-x11-color-number
func hexToXterm256(hex string) string {
	var hasKey bool
	if _, hasKey = cachedCodes[hex]; hasKey {
		return cachedCodes[hex]
	}

	var average uint64
	var b uint64
	var cb uint64
	var cg uint64
	var cidx uint64
	var clrErr float64
	var cr uint64
	var g uint64
	var gidx uint64
	var grayErr float64
	var gv uint64
	var i2cv []uint64
	var ib uint64
	var ig uint64
	var ir uint64
	var matches [][]string
	var r uint64

	// For simplicity, assume RGB space is perceptually uniform.
	// There are 5 places where one of two outputs needs to be
	// chosen when the input is the exact middle:
	// - The r/g/b channels and the gray value:
	//     - the higher value output is chosen
	// - If the gray and color values have same distance from the
	//   input
	//     - color is chosen

	// Calculate the nearest 0-based color index at 16..231
	r, g, b = 0, 0, 0
	matches = parseHex.FindAllStringSubmatch(hex, -1)
	for _, match := range matches {
		r, _ = strconv.ParseUint(match[1], 16, 8)
		g, _ = strconv.ParseUint(match[2], 16, 8)
		b, _ = strconv.ParseUint(match[3], 16, 8)
	}

	// 0..5 each
	ir = (r - 35) / 40
	if r < 48 {
		ir = 0
	} else if r < 115 {
		ir = 1
	}
	ig = (g - 35) / 40
	if g < 48 {
		ig = 0
	} else if g < 115 {
		ig = 1
	}
	ib = (b - 35) / 40
	if b < 48 {
		ib = 0
	} else if b < 115 {
		ib = 1
	}

	// 0..215 lazy evaluation
	cidx = (36 * ir) + (6 * ig) + ib + 16

	// Calculate the nearest 0-based gray index at 232..255
	average = (r + g + b) / 3

	// 0..23
	gidx = (average - 3) / 10
	if average > 238 {
		gidx = 23
	}

	// Calculate the represented colors back from the index
	i2cv = []uint64{0, 0x5f, 0x87, 0xaf, 0xd7, 0xff}

	// r/g/b 0..255 each
	cr = i2cv[ir]
	cg = i2cv[ig]
	cb = i2cv[ib]

	// same value for r/g/b 0..255
	gv = (10 * gidx) + 8

	// Return the one which is nearer to the original rgb values
	clrErr = math.Pow(float64(cr-r), 2) +
		math.Pow(float64(cg-g), 2) + math.Pow(float64(cb-b), 2)
	grayErr = math.Pow(float64(gv-r), 2) +
		math.Pow(float64(gv-g), 2) + math.Pow(float64(gv-b), 2)

	if clrErr <= grayErr {
		cachedCodes[hex] = Sprintf("color_%03d", cidx)
	} else {
		cachedCodes[hex] = Sprintf("color_%03d", gidx+232)
	}

	return cachedCodes[hex]
}

// Hilight will add the appropriate ANSI code to the specified
// string.
func Hilight(code string, str string, args ...interface{}) string {
	var clr string
	var hasKey bool
	var matches [][]string
	var width int

	// Call the appropriate function
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
			matches = hexCode.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				clr = hexToXterm256(match[2])
				if strings.HasPrefix(code, "on_") {
					clr = "on_" + clr
				}
				return colorize(clr, str, args...)
			}

			// Check if wrap
			matches = wrap.FindAllStringSubmatch(code, -1)
			for _, match := range matches {
				// Determine wrap width, default to 80
				width = 80
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

// Hilights will add the appropriate ANSI codes to the specified
// string.
func Hilights(
	codes []string,
	str string,
	args ...interface{},
) string {
	str = Sprintf(str, args...)

	// Apply all specified color codes
	for _, code := range codes {
		str = Hilight(code, str)
	}

	return str
}

func modify(mode string, str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	var hasKey bool
	var modified string
	var off string
	var opposite string
	var r *regexp.Regexp
	var rm string

	// Reverse mode
	opposite = "no_" + mode
	if strings.HasPrefix(opposite, "no_no_") {
		opposite = doubleno.ReplaceAllString(opposite, "")
	}

	// Store specified mode code for removal
	rm = Modes[mode]

	// Determine the off color code, if it exists
	off = ""
	if _, hasKey = Modes[opposite]; hasKey {
		// Store opposite code for removal
		rm += "|" + Modes[opposite]

		// Save opposite mode code sequence if it starts with "no_"
		if strings.HasPrefix(opposite, "no_") {
			off = "\x1b[" + Modes[opposite] + "m"
		}
	}

	// Remove other occurrences of specified mode and opposite
	r = regexp.MustCompile(`\x1b\[(` + rm + ")m")
	str = r.ReplaceAllString(Sprintf(str, args...), "")

	// Wrap the whole thing with specified color code
	modified = "\x1b[" + Modes[mode] + "m" + str + off

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(modified, "$1$4")
}

// OnHex will take a hex color code and add the appropriate ANSI color
// codes to the provided string.
func OnHex(hex string, str string, args ...interface{}) string {
	return colorize("on_"+hexToXterm256(hex), str, args...)
}

// OnRainbow will add multiple ANSI color codes to a string for a
// rainbow effect.
func OnRainbow(str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	var chars []string
	var code string
	var colors []int
	var end string
	var line []string
	var lines []string
	var out []string

	// Strip all other bg color codes and split on newline
	lines = strings.Split(plainBg(str, args...), "\n")

	// Loop thru lines and apply bg color codes
	colors = rainbowColors()
	end = "\x1b[" + Colors["on_default"] + "m"
	for i := range lines {
		chars = iterate.FindAllString(lines[i], -1)
		line = []string{}

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			code = strconv.Itoa(colors[idx%len(colors)] + 10)
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again, ensure on_default code at end
		out = append(out, strings.Join(line, "")+end)
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

// Plain will strip all ANSI color codes from a string.
func Plain(str string, args ...interface{}) string {
	// Strip all color codes
	return allCodes.ReplaceAllString(Sprintf(str, args...), "")
}

func plainBg(str string, args ...interface{}) string {
	return bgCodes.ReplaceAllString(Sprintf(str, args...), "")
}

func plainFg(str string, args ...interface{}) string {
	return fgCodes.ReplaceAllString(Sprintf(str, args...), "")
}

// Rainbow will add multiple ANSI color codes to a string for a
// rainbow effect.
func Rainbow(str string, args ...interface{}) string {
	if Disable {
		// Return the string w/o any color codes
		return Plain(str, args...)
	}

	var chars []string
	var code string
	var colors []int
	var line []string
	var lines []string
	var out []string

	// Strip all other fg color codes and split on newline
	lines = strings.Split(plainFg(str, args...), "\n")

	// Loop thru lines and apply fg color codes
	colors = rainbowColors()
	for i := range lines {
		chars = iterate.FindAllString(lines[i], -1)
		line = []string{}

		// Loop thru non-color-code bytes and apply on_rainbow
		for idx, char := range chars {
			code = strconv.Itoa(colors[idx%len(colors)])
			line = append(line, "\x1b["+code+"m"+char)
		}

		// Put line back together again
		out = append(out, strings.Join(line, ""))
	}

	// Put lines back together, and remove color codes if the line
	// only contains color codes
	return onlyCodes.ReplaceAllString(strings.Join(out, "\n"), "$1$4")
}

func rainbowColors() []int {
	// Don't include black, white, light_black, and light_white
	return []int{31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96}
}

// Sample will show all bg/fg combos of the first 16 8-bit colors.
func Sample() {
	var bg string
	var fg string

	for f := 0; f < 16; f++ {
		for b := 0; b < 16; b++ {
			fg = Sprintf("color_%03d", f)
			bg = Sprintf("on_color_%03d", b)
			Print(colorize(fg, colorize(bg, " mw ")))
		}
		Print("\n")
	}
}

// Table will display a pretty table of all 8-bit colors.
func Table() {
	var bg string

	for i := 0; i < 16; i++ {
		bg = Sprintf("on_color_%03d", i)
		Print(
			colorize(
				bg,
				" %s %s ",
				Black("%03d", i),
				White("%03d", i),
			),
		)
		if (i+1)%8 == 0 {
			Print("\n")
		}
	}

	for i := 16; i < 256; i++ {
		bg = Sprintf("on_color_%03d", i)
		Print(
			colorize(
				bg,
				" %s %s ",
				Black("%03d", i),
				White("%03d", i),
			),
		)
		if (i-15)%6 == 0 {
			Print("\n")
		}
	}
}

// Wrap will wrap a string to the specified width.
func Wrap(width int, str string, args ...interface{}) string {
	var lc int
	var line = ""
	var lines []string
	var wc int
	var words = strings.Fields(str)

	str = Sprintf(str, args...)

	// Loop thru words
	for _, word := range words {
		lc = len([]rune(Plain(line)))
		wc = len([]rune(Plain(word)))

		if lc == 0 {
			// Can't wrap less than a single word
			line = word
		} else if lc+wc+1 > width {
			// Wrap if line would be longer than width
			lines = append(lines, line)
			line = word
		} else {
			// Append word to line
			line += " " + word
		}
	}

	// Ensure last line is not forgotten
	if len([]rune(line)) != 0 {
		lines = append(lines, line)
	}

	// If original string ended with newline, put it back
	if strings.HasSuffix(str, "\n") {
		lines = append(lines, "")
	}

	// Join lines and return
	return strings.Join(lines, "\n")
}
