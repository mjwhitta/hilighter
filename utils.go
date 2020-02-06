package hilighter

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func bgColor(code string, str string) string {
	var colorized string

	// Strip all other bg color codes and don't extend bg color over
	// newlines
	str = newline.ReplaceAllString(
		plainBg(str),
		"\x1b["+Colors["on_default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["on_default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
}

func colorize(clr string, str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	// Call the appropriate function
	if strings.HasPrefix(clr, "on_") {
		return bgColor(clr, str)
	}
	return fgColor(clr, str)
}

func fgColor(code string, str string) string {
	var colorized string

	// Strip all other fg color codes and don't extend fg color over
	// newlines
	str = newline.ReplaceAllString(
		plainFg(str),
		"\x1b["+Colors["default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["default"] + "m"

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(colorized, "$1$4")
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

func modify(mode string, str string) string {
	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
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
	str = r.ReplaceAllString(str, "")

	// Wrap the whole thing with specified color code
	modified = "\x1b[" + Modes[mode] + "m" + str + off

	// Remove color codes, if the line only contains color codes
	return onlyCodes.ReplaceAllString(modified, "$1$4")
}

func plainBg(str string) string {
	return bgCodes.ReplaceAllString(str, "")
}

func plainFg(str string) string {
	return fgCodes.ReplaceAllString(str, "")
}

func rainbowColors() []int {
	// Don't include black, white, light_black, and light_white
	return []int{31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96}
}
