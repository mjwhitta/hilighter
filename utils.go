package hilighter

import (
	"image/color"
	"math"
	"regexp"
	"strings"
)

func adjustedRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	var a uint32
	var b uint32
	var g uint32
	var r uint32

	r, g, b, a = c.RGBA()

	if (a != 0) && (a != math.MaxUint32) {
		r = uint32(float64(r*math.MaxUint32) / float64(a))
		g = uint32(float64(g*math.MaxUint32) / float64(a))
		b = uint32(float64(b*math.MaxUint32) / float64(a))
	}

	r >>= 8
	g >>= 8
	b >>= 8
	a >>= 8

	//nolint:gosec // No overflow issues b/c I shifted already
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func bgColor(code string, str string) string {
	var colorized string

	// Strip all other bg color codes and don't extend bg color over
	// newlines
	str = reNewline.ReplaceAllString(
		plainBg(str),
		"\x1b["+Colors["on_default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["on_default"] + "m"

	// Remove color codes, if the line only contains color codes
	return reOnlyCodes.ReplaceAllString(colorized, "$1$4")
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
	str = reNewline.ReplaceAllString(
		plainFg(str),
		"\x1b["+Colors["default"]+"m\n\x1b["+Colors[code]+"m",
	)

	// Wrap whole thing with specified color code
	colorized = "\x1b[" + Colors[code] + "m" + str +
		"\x1b[" + Colors["default"] + "m"

	// Remove color codes, if the line only contains color codes
	return reOnlyCodes.ReplaceAllString(colorized, "$1$4")
}

func modify(mode string, str string) string {
	var modified string
	var off string
	var opposite string
	var r *regexp.Regexp
	var rm string

	if IsDisabled() {
		// Return the string w/o any color codes
		return Plain(str)
	}

	// Reverse mode
	opposite = "no_" + mode
	if strings.HasPrefix(opposite, "no_no_") {
		opposite = reDoubleNo.ReplaceAllString(opposite, "")
	}

	// Store specified mode code for removal
	rm = Modes[mode]

	// Determine the off color code, if it exists
	if _, ok := Modes[opposite]; ok {
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
	return reOnlyCodes.ReplaceAllString(modified, "$1$4")
}

func plainBg(str string) string {
	return reBgCodes.ReplaceAllString(str, "")
}

func plainFg(str string) string {
	return reFgCodes.ReplaceAllString(str, "")
}

func rainbowColors() []int {
	// Don't include black, white, light_black, and light_white
	return []int{31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96}
}
